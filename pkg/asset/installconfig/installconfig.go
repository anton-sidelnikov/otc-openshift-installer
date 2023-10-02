package installconfig

import (
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig/alibabacloud"
	icopenstack "github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig/openstack"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/defaults"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/validation"
)

const (
	installConfigFilename = "install-config.yaml"
)

// InstallConfig generates the install-config.yaml file.
type InstallConfig struct {
	AssetBase
	AlibabaCloud *alibabacloud.Metadata `json:"alibabacloud,omitempty"`
}

var _ asset.WritableAsset = (*InstallConfig)(nil)

// MakeAsset returns an InstallConfig asset containing a given InstallConfig CR.
func MakeAsset(config *types.InstallConfig) *InstallConfig {
	return &InstallConfig{
		AssetBase: AssetBase{
			Config: config,
		},
	}
}

// Dependencies returns all of the dependencies directly needed by an
// InstallConfig asset.
func (a *InstallConfig) Dependencies() []asset.Asset {
	return []asset.Asset{
		&sshPublicKey{},
		&baseDomain{},
		&clusterName{},
		&networking{},
		&pullSecret{},
		&platform{},
	}
}

// Generate generates the install-config.yaml file.
func (a *InstallConfig) Generate(parents asset.Parents) error {
	sshPublicKey := &sshPublicKey{}
	baseDomain := &baseDomain{}
	clusterName := &clusterName{}
	networking := &networking{}
	pullSecret := &pullSecret{}
	platform := &platform{}
	parents.Get(
		sshPublicKey,
		baseDomain,
		clusterName,
		networking,
		pullSecret,
		platform,
	)

	a.Config = &types.InstallConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: types.InstallConfigVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterName.ClusterName,
		},
		SSHKey:     sshPublicKey.Key,
		BaseDomain: baseDomain.BaseDomain,
		PullSecret: pullSecret.PullSecret,
		Networking: &types.Networking{
			MachineNetwork: networking.machineNetwork,
		},
	}

	a.Config.AlibabaCloud = platform.AlibabaCloud
	a.Config.None = platform.None
	a.Config.OpenStack = platform.OpenStack
	defaults.SetInstallConfigDefaults(a.Config)

	return a.finish("")
}

// Load returns the installconfig from disk.
func (a *InstallConfig) Load(f asset.FileFetcher) (found bool, err error) {
	found, err = a.LoadFromFile(f)
	if found && err == nil {
		if err := a.finish(installConfigFilename); err != nil {
			return false, errors.Wrap(err, asset.InstallConfigError)
		}
	}

	return found, err
}

func (a *InstallConfig) finish(filename string) error {
	if a.Config.AlibabaCloud != nil {
		a.AlibabaCloud = alibabacloud.NewMetadata(a.Config.AlibabaCloud.Region, a.Config.AlibabaCloud.VSwitchIDs)
	}

	if err := validation.ValidateInstallConfig(a.Config, false).ToAggregate(); err != nil {
		if filename == "" {
			return errors.Wrap(err, "invalid install config")
		}
		return errors.Wrapf(err, "invalid %q file", filename)
	}

	if err := a.platformValidation(); err != nil {
		return err
	}

	return a.RecordFile()
}

// platformValidation runs validations that require connecting to the
// underlying platform. In some cases, platforms also duplicate validations
// that have already been checked by validation.ValidateInstallConfig().
func (a *InstallConfig) platformValidation() error {
	if a.Config.Platform.AlibabaCloud != nil {
		client, err := a.AlibabaCloud.Client()
		if err != nil {
			return err
		}
		return alibabacloud.Validate(client, a.Config)
	}
	if a.Config.Platform.OpenStack != nil {
		return icopenstack.Validate(a.Config)
	}
	return field.ErrorList{}.ToAggregate()
}
