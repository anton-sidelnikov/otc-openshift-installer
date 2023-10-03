package machines

import (
	"fmt"
	baremetalhost "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/ignition/machine"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/machines/machineconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/machines/openstack"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/rhcos"
	rhcosutils "github.com/anton-sidelnikov/otc-openshift-installer/pkg/rhcos"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	openstacktypes "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
	machinev1 "github.com/openshift/api/machine/v1"
	machinev1alpha1 "github.com/openshift/api/machine/v1alpha1"
	machinev1beta1 "github.com/openshift/api/machine/v1beta1"
	mcfgv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
)

// Master generates the machines for the `master` machine pool.
type Master struct {
	UserDataFile           *asset.File
	MachineConfigFiles     []*asset.File
	MachineFiles           []*asset.File
	ControlPlaneMachineSet *asset.File

	// SecretFiles is used by the baremetal platform to register the
	// credential information for communicating with management
	// controllers on hosts.
	SecretFiles []*asset.File

	// NetworkConfigSecretFiles is used by the baremetal platform to
	// store the networking configuration per host
	NetworkConfigSecretFiles []*asset.File

	// HostFiles is the list of baremetal hosts provided in the
	// installer configuration.
	HostFiles []*asset.File
}

const (
	directory = "openshift"

	// secretFileName is the format string for constructing the Secret
	// filenames for baremetal clusters.
	secretFileName = "99_openshift-cluster-api_host-bmc-secrets-%s.yaml"

	// networkConfigSecretFileName is the format string for constructing
	// the networking configuration Secret filenames for baremetal
	// clusters.
	networkConfigSecretFileName = "99_openshift-cluster-api_host-network-config-secrets-%s.yaml"

	// hostFileName is the format string for constucting the Host
	// filenames for baremetal clusters.
	hostFileName = "99_openshift-cluster-api_hosts-%s.yaml"

	// masterMachineFileName is the format string for constucting the
	// master Machine filenames.
	masterMachineFileName = "99_openshift-cluster-api_master-machines-%s.yaml"

	// masterUserDataFileName is the filename used for the master
	// user-data secret.
	masterUserDataFileName = "99_openshift-cluster-api_master-user-data-secret.yaml"

	// masterUserDataFileName is the filename used for the control plane machine sets.
	controlPlaneMachineSetFileName = "99_openshift-machine-api_master-control-plane-machine-set.yaml"
)

var (
	secretFileNamePattern              = fmt.Sprintf(secretFileName, "*")
	networkConfigSecretFileNamePattern = fmt.Sprintf(networkConfigSecretFileName, "*")
	hostFileNamePattern                = fmt.Sprintf(hostFileName, "*")
	masterMachineFileNamePattern       = fmt.Sprintf(masterMachineFileName, "*")

	_ asset.WritableAsset = (*Master)(nil)
)

// Name returns a human friendly name for the Master Asset.
func (m *Master) Name() string {
	return "Master Machines"
}

// Dependencies returns all of the dependencies directly needed by the
// Master asset
func (m *Master) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.ClusterID{},
		// PlatformCredsCheck just checks the creds (and asks, if needed)
		// We do not actually use it in this asset directly, hence
		// it is put in the dependencies but not fetched in Generate
		&installconfig.PlatformCredsCheck{},
		&installconfig.InstallConfig{},
		new(rhcos.Image),
		&machine.Master{},
	}
}

// Generate generates the Master asset.
func (m *Master) Generate(dependencies asset.Parents) error {
	clusterID := &installconfig.ClusterID{}
	installConfig := &installconfig.InstallConfig{}
	rhcosImage := new(rhcos.Image)
	mign := &machine.Master{}
	dependencies.Get(clusterID, installConfig, rhcosImage, mign)

	masterUserDataSecretName := "master-user-data"

	ic := installConfig.Config

	pool := *ic.ControlPlane
	var err error
	machines := []machinev1beta1.Machine{}
	var controlPlaneMachineSet *machinev1.ControlPlaneMachineSet
	switch ic.Platform.Name() {
	case openstacktypes.Name:
		mpool := defaultOpenStackMachinePoolPlatform()
		mpool.Set(ic.Platform.OpenStack.DefaultMachinePlatform)
		mpool.Set(pool.Platform.OpenStack)
		pool.Platform.OpenStack = &mpool

		imageName, _ := rhcosutils.GenerateOpenStackImageName(string(*rhcosImage), clusterID.InfraID)

		machines, controlPlaneMachineSet, err = openstack.Machines(clusterID.InfraID, ic, &pool, imageName, "master", masterUserDataSecretName)
		if err != nil {
			return errors.Wrap(err, "failed to create master machine objects")
		}
		openstack.ConfigMasters(machines, clusterID.InfraID)
	default:
		return fmt.Errorf("invalid Platform")
	}

	data, err := userDataSecret(masterUserDataSecretName, mign.File.Data)
	if err != nil {
		return errors.Wrap(err, "failed to create user-data secret for master machines")
	}

	m.UserDataFile = &asset.File{
		Filename: filepath.Join(directory, masterUserDataFileName),
		Data:     data,
	}

	machineConfigs := []*mcfgv1.MachineConfig{}
	if pool.Hyperthreading == types.HyperthreadingDisabled {
		ignHT, err := machineconfig.ForHyperthreadingDisabled("master")
		if err != nil {
			return errors.Wrap(err, "failed to create ignition for hyperthreading disabled for master machines")
		}
		machineConfigs = append(machineConfigs, ignHT)
	}
	if ic.SSHKey != "" {
		ignSSH, err := machineconfig.ForAuthorizedKeys(ic.SSHKey, "master")
		if err != nil {
			return errors.Wrap(err, "failed to create ignition for authorized SSH keys for master machines")
		}
		machineConfigs = append(machineConfigs, ignSSH)
	}
	if ic.FIPS {
		ignFIPS, err := machineconfig.ForFIPSEnabled("master")
		if err != nil {
			return errors.Wrap(err, "failed to create ignition for FIPS enabled for master machines")
		}
		machineConfigs = append(machineConfigs, ignFIPS)
	}
	// The maximum number of networks supported on ServiceNetwork is two, one IPv4 and one IPv6 network.
	// The cluster-network-operator handles the validation of this field.
	// Reference: https://github.com/openshift/cluster-network-operator/blob/fc3e0e25b4cfa43e14122bdcdd6d7f2585017d75/pkg/network/cluster_config.go#L45-L52
	if ic.Networking != nil && len(ic.Networking.ServiceNetwork) == 2 &&
		(ic.Platform.Name() == openstacktypes.Name) {
		// Only configure kernel args for dual-stack clusters.
		ignIPv6, err := machineconfig.ForDualStackAddresses("master")
		if err != nil {
			return errors.Wrap(err, "failed to create ignition to configure IPv6 for master machines")
		}
		machineConfigs = append(machineConfigs, ignIPv6)
	}

	m.MachineConfigFiles, err = machineconfig.Manifests(machineConfigs, "master", directory)
	if err != nil {
		return errors.Wrap(err, "failed to create MachineConfig manifests for master machines")
	}

	m.MachineFiles = make([]*asset.File, len(machines))
	if controlPlaneMachineSet != nil && *pool.Replicas > 1 {
		data, err := yaml.Marshal(controlPlaneMachineSet)
		if err != nil {
			return errors.Wrapf(err, "marshal control plane machine set")
		}
		m.ControlPlaneMachineSet = &asset.File{
			Filename: filepath.Join(directory, controlPlaneMachineSetFileName),
			Data:     data,
		}
	}
	padFormat := fmt.Sprintf("%%0%dd", len(fmt.Sprintf("%d", len(machines))))
	for i, machine := range machines {
		data, err := yaml.Marshal(machine)
		if err != nil {
			return errors.Wrapf(err, "marshal master %d", i)
		}

		padded := fmt.Sprintf(padFormat, i)
		m.MachineFiles[i] = &asset.File{
			Filename: filepath.Join(directory, fmt.Sprintf(masterMachineFileName, padded)),
			Data:     data,
		}
	}
	return nil
}

// Files returns the files generated by the asset.
func (m *Master) Files() []*asset.File {
	files := make([]*asset.File, 0, 1+len(m.MachineConfigFiles)+len(m.MachineFiles))
	if m.UserDataFile != nil {
		files = append(files, m.UserDataFile)
	}
	files = append(files, m.MachineConfigFiles...)
	// Hosts refer to secrets, so place the secrets before the hosts
	// to avoid unnecessary reconciliation errors.
	files = append(files, m.SecretFiles...)
	files = append(files, m.NetworkConfigSecretFiles...)
	// Machines are linked to hosts via the machineRef, so we create
	// the hosts first to ensure if the operator starts trying to
	// reconcile a machine it can pick up the related host.
	files = append(files, m.HostFiles...)
	files = append(files, m.MachineFiles...)
	if m.ControlPlaneMachineSet != nil {
		files = append(files, m.ControlPlaneMachineSet)
	}
	return files
}

// Load reads the asset files from disk.
func (m *Master) Load(f asset.FileFetcher) (found bool, err error) {
	file, err := f.FetchByName(filepath.Join(directory, masterUserDataFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	m.UserDataFile = file

	m.MachineConfigFiles, err = machineconfig.Load(f, "master", directory)
	if err != nil {
		return true, err
	}

	var fileList []*asset.File

	fileList, err = f.FetchByPattern(filepath.Join(directory, secretFileNamePattern))
	if err != nil {
		return true, err
	}
	m.SecretFiles = fileList

	fileList, err = f.FetchByPattern(filepath.Join(directory, networkConfigSecretFileNamePattern))
	if err != nil {
		return true, err
	}
	m.NetworkConfigSecretFiles = fileList

	fileList, err = f.FetchByPattern(filepath.Join(directory, hostFileNamePattern))
	if err != nil {
		return true, err
	}
	m.HostFiles = fileList

	fileList, err = f.FetchByPattern(filepath.Join(directory, masterMachineFileNamePattern))
	if err != nil {
		return true, err
	}
	m.MachineFiles = fileList

	file, err = f.FetchByName(filepath.Join(directory, controlPlaneMachineSetFileName))
	if err != nil {
		if os.IsNotExist(err) {
			// Choosing to ignore the CPMS file if it does not exist since UPI does not need it.
			logrus.Debugf("CPMS file missing. Ignoring it while loading machine asset.")
			return true, nil
		}
		return true, err
	}
	m.ControlPlaneMachineSet = file

	return true, nil
}

// Machines returns master Machine manifest structures.
func (m *Master) Machines() ([]machinev1beta1.Machine, error) {
	scheme := runtime.NewScheme()
	scheme.AddKnownTypes(machinev1alpha1.GroupVersion,
		&machinev1alpha1.OpenstackProviderSpec{},
	)
	scheme.AddKnownTypes(machinev1.GroupVersion,
		&machinev1.ControlPlaneMachineSet{},
	)

	machinev1beta1.AddToScheme(scheme)
	machinev1.Install(scheme)
	decoder := serializer.NewCodecFactory(scheme).UniversalDecoder(
		machinev1.GroupVersion,
		machinev1alpha1.GroupVersion,
		machinev1beta1.SchemeGroupVersion,
	)

	machines := []machinev1beta1.Machine{}
	for i, file := range m.MachineFiles {
		machine := &machinev1beta1.Machine{}
		err := yaml.Unmarshal(file.Data, &machine)
		if err != nil {
			return machines, errors.Wrapf(err, "unmarshal master %d", i)
		}

		obj, _, err := decoder.Decode(machine.Spec.ProviderSpec.Value.Raw, nil, nil)
		if err != nil {
			return machines, errors.Wrapf(err, "unmarshal master %d", i)
		}

		machine.Spec.ProviderSpec.Value = &runtime.RawExtension{Object: obj}
		machines = append(machines, *machine)
	}

	return machines, nil
}

// IsMachineManifest tests whether a file is a manifest that belongs to the
// Master Machines or Worker Machines asset.
func IsMachineManifest(file *asset.File) bool {
	if filepath.Dir(file.Filename) != directory {
		return false
	}
	filename := filepath.Base(file.Filename)
	if filename == masterUserDataFileName || filename == workerUserDataFileName || filename == controlPlaneMachineSetFileName {
		return true
	}
	if matched, err := machineconfig.IsManifest(filename); err != nil {
		panic(err)
	} else if matched {
		return true
	}
	if matched, err := filepath.Match(masterMachineFileNamePattern, filename); err != nil {
		panic("bad format for master machine file name pattern")
	} else if matched {
		return true
	}
	if matched, err := filepath.Match(workerMachineSetFileNamePattern, filename); err != nil {
		panic("bad format for worker machine file name pattern")
	} else {
		return matched
	}
}

func createSecretAssetFiles(resources []corev1.Secret, fileName string) ([]*asset.File, error) {

	var objects []interface{}
	for _, r := range resources {
		objects = append(objects, r)
	}

	return createAssetFiles(objects, fileName)
}

func createHostAssetFiles(resources []baremetalhost.BareMetalHost, fileName string) ([]*asset.File, error) {

	var objects []interface{}
	for _, r := range resources {
		objects = append(objects, r)
	}

	return createAssetFiles(objects, fileName)
}

func createAssetFiles(objects []interface{}, fileName string) ([]*asset.File, error) {

	assetFiles := make([]*asset.File, len(objects))
	padFormat := fmt.Sprintf("%%0%dd", len(fmt.Sprintf("%d", len(objects))))
	for i, obj := range objects {
		data, err := yaml.Marshal(obj)
		if err != nil {
			return nil, errors.Wrapf(err, "marshal resource %d", i)
		}
		padded := fmt.Sprintf(padFormat, i)
		assetFiles[i] = &asset.File{
			Filename: filepath.Join(directory, fmt.Sprintf(fileName, padded)),
			Data:     data,
		}
	}

	return assetFiles, nil
}
