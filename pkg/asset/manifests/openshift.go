package manifests

import (
	"encoding/base64"
	"github.com/gophercloud/utils/openstack/clientconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/sets"
	"path/filepath"
	"sigs.k8s.io/yaml"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/machines"
	osmachine "github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/machines/openstack"
	openstackmanifests "github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/manifests/openstack"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/openshiftinstall"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/password"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/rhcos"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/templates/content/openshift"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	openstacktypes "github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/openstack"
)

const (
	openshiftManifestDir = "openshift"
)

var (
	_ asset.WritableAsset = (*Openshift)(nil)
)

// Openshift generates the dependent resource manifests for openShift (as against bootkube)
type Openshift struct {
	FileList []*asset.File
}

// Name returns a human friendly name for the operator
func (o *Openshift) Name() string {
	return "Openshift Manifests"
}

// Dependencies returns all of the dependencies directly needed by the
// Openshift asset
func (o *Openshift) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.InstallConfig{},
		&installconfig.ClusterID{},
		&password.KubeadminPassword{},
		&openshiftinstall.Config{},
		&FeatureGate{},

		&openshift.CloudCredsSecret{},
		&openshift.KubeadminPasswordSecret{},
		&openshift.RoleCloudCredsSecretReader{},
		&openshift.BaremetalConfig{},
		new(rhcos.Image),
		&openshift.AzureCloudProviderSecret{},
	}
}

// Generate generates the respective operator config.yml files
func (o *Openshift) Generate(dependencies asset.Parents) error {
	installConfig := &installconfig.InstallConfig{}
	clusterID := &installconfig.ClusterID{}
	kubeadminPassword := &password.KubeadminPassword{}
	openshiftInstall := &openshiftinstall.Config{}
	featureGate := &FeatureGate{}
	dependencies.Get(installConfig, kubeadminPassword, clusterID, openshiftInstall, featureGate)
	var cloudCreds cloudCredsSecretData
	platform := installConfig.Config.Platform.Name()
	switch platform {
	case openstacktypes.Name:
		opts := new(clientconfig.ClientOpts)
		opts.Cloud = installConfig.Config.Platform.OpenStack.Cloud
		cloud, err := clientconfig.GetCloudFromYAML(opts)
		if err != nil {
			return err
		}

		// We need to replace the local cacert path with one that is used in OpenShift
		if cloud.CACertFile != "" {
			cloud.CACertFile = "/etc/kubernetes/static-pod-resources/configmaps/cloud-config/ca-bundle.pem"
		}

		// Application credentials are easily rotated in the event of a leak and should be preferred. Encourage their use.
		authTypes := sets.New(clientconfig.AuthPassword, clientconfig.AuthV2Password, clientconfig.AuthV3Password)
		if cloud.AuthInfo != nil && authTypes.Has(cloud.AuthType) {
			logrus.Warnf(
				"clouds.yaml file is using %q type auth. Consider using the %q auth type instead to rotate credentials more easily.",
				cloud.AuthType,
				clientconfig.AuthV3ApplicationCredential,
			)
		}

		clouds := make(map[string]map[string]*clientconfig.Cloud)
		clouds["clouds"] = map[string]*clientconfig.Cloud{
			osmachine.CloudName: cloud,
		}

		marshalled, err := yaml.Marshal(clouds)
		if err != nil {
			return err
		}

		cloudProviderConf, err := openstackmanifests.CloudProviderConfigSecret(cloud)
		if err != nil {
			return err
		}

		credsEncoded := base64.StdEncoding.EncodeToString(marshalled)
		credsINIEncoded := base64.StdEncoding.EncodeToString(cloudProviderConf)
		cloudCreds = cloudCredsSecretData{
			OpenStack: &OpenStackCredsSecretData{
				Base64encodeCloudCreds:    credsEncoded,
				Base64encodeCloudCredsINI: credsINIEncoded,
			},
		}
	}

	templateData := &openshiftTemplateData{
		CloudCreds:                   cloudCreds,
		Base64EncodedKubeadminPwHash: base64.StdEncoding.EncodeToString(kubeadminPassword.PasswordHash),
	}

	cloudCredsSecret := &openshift.CloudCredsSecret{}
	kubeadminPasswordSecret := &openshift.KubeadminPasswordSecret{}
	roleCloudCredsSecretReader := &openshift.RoleCloudCredsSecretReader{}
	baremetalConfig := &openshift.BaremetalConfig{}
	rhcosImage := new(rhcos.Image)

	dependencies.Get(
		cloudCredsSecret,
		kubeadminPasswordSecret,
		roleCloudCredsSecretReader,
		baremetalConfig,
		rhcosImage)

	assetData := map[string][]byte{
		"99_kubeadmin-password-secret.yaml": applyTemplateData(kubeadminPasswordSecret.Files()[0].Data, templateData),
	}

	switch platform {
	case openstacktypes.Name:
		if installConfig.Config.CredentialsMode != types.ManualCredentialsMode {
			assetData["99_cloud-creds-secret.yaml"] = applyTemplateData(cloudCredsSecret.Files()[0].Data, templateData)
		}
		assetData["99_role-cloud-creds-secret-reader.yaml"] = applyTemplateData(roleCloudCredsSecretReader.Files()[0].Data, templateData)
	}
	o.FileList = []*asset.File{}
	for name, data := range assetData {
		if len(data) == 0 {
			continue
		}
		o.FileList = append(o.FileList, &asset.File{
			Filename: filepath.Join(openshiftManifestDir, name),
			Data:     data,
		})
	}

	o.FileList = append(o.FileList, openshiftInstall.Files()...)
	o.FileList = append(o.FileList, featureGate.Files()...)

	asset.SortFiles(o.FileList)

	return nil
}

// Files returns the files generated by the asset.
func (o *Openshift) Files() []*asset.File {
	return o.FileList
}

// Load returns the openshift asset from disk.
func (o *Openshift) Load(f asset.FileFetcher) (bool, error) {
	yamlFileList, err := f.FetchByPattern(filepath.Join(openshiftManifestDir, "*.yaml"))
	if err != nil {
		return false, errors.Wrap(err, "failed to load *.yaml files")
	}
	ymlFileList, err := f.FetchByPattern(filepath.Join(openshiftManifestDir, "*.yml"))
	if err != nil {
		return false, errors.Wrap(err, "failed to load *.yml files")
	}
	jsonFileList, err := f.FetchByPattern(filepath.Join(openshiftManifestDir, "*.json"))
	if err != nil {
		return false, errors.Wrap(err, "failed to load *.json files")
	}
	fileList := append(yamlFileList, ymlFileList...)
	fileList = append(fileList, jsonFileList...)

	for _, file := range fileList {
		if machines.IsMachineManifest(file) {
			continue
		}

		o.FileList = append(o.FileList, file)
	}

	asset.SortFiles(o.FileList)
	return len(o.FileList) > 0, nil
}
