package cluster

import (
	"encoding/json"
	igntypes "github.com/coreos/ignition/v2/config/v3_2/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"sigs.k8s.io/yaml"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/ignition"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/ignition/bootstrap"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/ignition/machine"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/machines"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/manifests"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/openshiftinstall"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/rhcos"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/tfvars"
	openstacktfvars "github.com/anton-sidelnikov/otc-openshift-installer/pkg/tfvars/openstack"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
	configv1 "github.com/openshift/api/config/v1"
)

const (
	// GCPFirewallPermission is the role/permission to create or skip the creation of
	// firewall rules for GCP during a xpn installation.
	GCPFirewallPermission = "compute.firewalls.create"

	// TfVarsFileName is the filename for Terraform variables.
	TfVarsFileName = "terraform.tfvars.json"

	// TfPlatformVarsFileName is the name for platform-specific
	// Terraform variable files.
	//
	// https://www.terraform.io/docs/configuration/variables.html#variable-files
	TfPlatformVarsFileName = "terraform.platform.auto.tfvars.json"

	tfvarsAssetName = "Terraform Variables"
)

// TerraformVariables depends on InstallConfig, Manifests,
// and Ignition to generate the terrafor.tfvars.
type TerraformVariables struct {
	FileList []*asset.File
}

var _ asset.WritableAsset = (*TerraformVariables)(nil)

// Name returns the human-friendly name of the asset.
func (t *TerraformVariables) Name() string {
	return tfvarsAssetName
}

// Dependencies returns the dependency of the TerraformVariable
func (t *TerraformVariables) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.ClusterID{},
		&installconfig.InstallConfig{},
		new(rhcos.Image),
		new(rhcos.Release),
		new(rhcos.BootstrapImage),
		&bootstrap.Bootstrap{},
		&machine.Master{},
		&machines.Master{},
		&machines.Worker{},
		&installconfig.PlatformProvisionCheck{},
		&manifests.Manifests{},
	}
}

// Generate generates the terraform.tfvars file.
func (t *TerraformVariables) Generate(parents asset.Parents) error {
	clusterID := &installconfig.ClusterID{}
	installConfig := &installconfig.InstallConfig{}
	bootstrapIgnAsset := &bootstrap.Bootstrap{}
	masterIgnAsset := &machine.Master{}
	mastersAsset := &machines.Master{}
	workersAsset := &machines.Worker{}
	manifestsAsset := &manifests.Manifests{}
	rhcosImage := new(rhcos.Image)
	rhcosRelease := new(rhcos.Release)
	rhcosBootstrapImage := new(rhcos.BootstrapImage)
	parents.Get(clusterID, installConfig, bootstrapIgnAsset, masterIgnAsset, mastersAsset, workersAsset, manifestsAsset, rhcosImage, rhcosRelease, rhcosBootstrapImage)

	platform := installConfig.Config.Platform.Name()

	masterIgn := string(masterIgnAsset.Files()[0].Data)
	bootstrapIgn, err := injectInstallInfo(bootstrapIgnAsset.Files()[0].Data)
	if err != nil {
		return errors.Wrap(err, "unable to inject installation info")
	}

	var useIPv4, useIPv6 bool
	for _, network := range installConfig.Config.Networking.ServiceNetwork {
		if network.IP.To4() != nil {
			useIPv4 = true
		} else {
			useIPv6 = true
		}
	}

	machineV4CIDRs, machineV6CIDRs := []string{}, []string{}
	for _, network := range installConfig.Config.Networking.MachineNetwork {
		if network.CIDR.IPNet.IP.To4() != nil {
			machineV4CIDRs = append(machineV4CIDRs, network.CIDR.IPNet.String())
		} else {
			machineV6CIDRs = append(machineV6CIDRs, network.CIDR.IPNet.String())
		}
	}

	masterCount := len(mastersAsset.MachineFiles)
	mastersSchedulable := false
	for _, f := range manifestsAsset.Files() {
		if f.Filename == manifests.SchedulerCfgFilename {
			schedulerConfig := configv1.Scheduler{}
			err = yaml.Unmarshal(f.Data, &schedulerConfig)
			if err != nil {
				return errors.Wrapf(err, "failed to unmarshall %s", manifests.SchedulerCfgFilename)
			}
			mastersSchedulable = schedulerConfig.Spec.MastersSchedulable
			break
		}
	}

	data, err := tfvars.TFVars(
		clusterID.InfraID,
		installConfig.Config.ClusterDomain(),
		installConfig.Config.BaseDomain,
		machineV4CIDRs,
		machineV6CIDRs,
		useIPv4,
		useIPv6,
		bootstrapIgn,
		masterIgn,
		masterCount,
		mastersSchedulable,
	)
	if err != nil {
		return errors.Wrap(err, "failed to get Terraform variables")
	}
	t.FileList = []*asset.File{
		{
			Filename: TfVarsFileName,
			Data:     data,
		},
	}

	if masterCount == 0 {
		return errors.Errorf("master slice cannot be empty")
	}

	switch platform {
	case openstack.Name:
		data, err = openstacktfvars.TFVars(
			installConfig,
			mastersAsset,
			workersAsset,
			string(*rhcosImage),
			clusterID,
			bootstrapIgn,
		)
		if err != nil {
			return errors.Wrapf(err, "failed to get %s Terraform variables", platform)
		}
		t.FileList = append(t.FileList, &asset.File{
			Filename: TfPlatformVarsFileName,
			Data:     data,
		})
	default:
		logrus.Warnf("unrecognized platform %s", platform)
	}

	return nil
}

// Files returns the files generated by the asset.
func (t *TerraformVariables) Files() []*asset.File {
	return t.FileList
}

// Load reads the terraform.tfvars from disk.
func (t *TerraformVariables) Load(f asset.FileFetcher) (found bool, err error) {
	file, err := f.FetchByName(TfVarsFileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}

	switch file, err := f.FetchByName(TfPlatformVarsFileName); {
	case err == nil:
		t.FileList = append(t.FileList, file)
	case !os.IsNotExist(err):
		return false, err
	}

	return true, nil
}

// injectInstallInfo adds information about the installer and its invoker as a
// ConfigMap to the provided bootstrap Ignition config.
func injectInstallInfo(bootstrap []byte) (string, error) {
	config := &igntypes.Config{}
	if err := json.Unmarshal(bootstrap, &config); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal bootstrap Ignition config")
	}

	cm, err := openshiftinstall.CreateInstallConfigMap("openshift-install")
	if err != nil {
		return "", errors.Wrap(err, "failed to generate openshift-install config")
	}

	config.Storage.Files = append(config.Storage.Files, ignition.FileFromString("/opt/openshift/manifests/openshift-install.yaml", "root", 0644, cm))

	ign, err := ignition.Marshal(config)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal bootstrap Ignition config")
	}

	return string(ign), nil
}
