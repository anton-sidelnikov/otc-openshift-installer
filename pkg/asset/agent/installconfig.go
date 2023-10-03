package agent

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types/validation"
)

const (
	installConfigFilename = "install-config.yaml"
)

// OptionalInstallConfig is an InstallConfig where the default is empty, rather
// than generated from running the survey.
type OptionalInstallConfig struct {
	installconfig.AssetBase
	Supplied bool
}

var _ asset.WritableAsset = (*OptionalInstallConfig)(nil)

// Dependencies returns all of the dependencies directly needed by an
// InstallConfig asset.
func (a *OptionalInstallConfig) Dependencies() []asset.Asset {
	// Return no dependencies for the Agent install config, because it is
	// optional. We don't need to run the survey if it doesn't exist, since the
	// user may have supplied cluster-manifests that fully define the cluster.
	return []asset.Asset{}
}

// Generate generates the install-config.yaml file.
func (a *OptionalInstallConfig) Generate(parents asset.Parents) error {
	// Just generate an empty install config, since we have no dependencies.
	return nil
}

// Load returns the installconfig from disk.
func (a *OptionalInstallConfig) Load(f asset.FileFetcher) (bool, error) {
	found, err := a.LoadFromFile(f)
	if found && err == nil {
		a.Supplied = true
		if err := a.validateInstallConfig(a.Config).ToAggregate(); err != nil {
			return false, errors.Wrapf(err, "invalid install-config configuration")
		}
		if err := a.RecordFile(); err != nil {
			return false, err
		}
	}
	return found, err
}

func (a *OptionalInstallConfig) validateInstallConfig(installConfig *types.InstallConfig) field.ErrorList {
	var allErrs field.ErrorList
	if err := validation.ValidateInstallConfig(a.Config, true); err != nil {
		allErrs = append(allErrs, err...)
	}

	if err := a.validateSupportedArchs(installConfig); err != nil {
		allErrs = append(allErrs, err...)
	}

	warnUnusedConfig(installConfig)

	numMasters, numWorkers := GetReplicaCount(installConfig)
	logrus.Infof(fmt.Sprintf("Configuration has %d master replicas and %d worker replicas", numMasters, numWorkers))

	if err := a.validateSNOConfiguration(installConfig); err != nil {
		allErrs = append(allErrs, err...)
	}

	return allErrs
}

func (a *OptionalInstallConfig) validateSupportedArchs(installConfig *types.InstallConfig) field.ErrorList {
	var allErrs field.ErrorList

	fieldPath := field.NewPath("ControlPlane", "Architecture")

	switch string(installConfig.ControlPlane.Architecture) {
	case types.ArchitectureAMD64:
	case types.ArchitectureARM64:
	case types.ArchitecturePPC64LE:
	default:
		allErrs = append(allErrs, field.NotSupported(fieldPath, installConfig.ControlPlane.Architecture, []string{types.ArchitectureAMD64, types.ArchitectureARM64, types.ArchitecturePPC64LE}))
	}

	for i, compute := range installConfig.Compute {
		fieldPath := field.NewPath(fmt.Sprintf("Compute[%d]", i), "Architecture")

		switch string(compute.Architecture) {
		case types.ArchitectureAMD64:
		case types.ArchitectureARM64:
		case types.ArchitecturePPC64LE:
		default:
			allErrs = append(allErrs, field.NotSupported(fieldPath, compute.Architecture, []string{types.ArchitectureAMD64, types.ArchitectureARM64, types.ArchitecturePPC64LE}))
		}
	}

	return allErrs
}

func (a *OptionalInstallConfig) validateSNOConfiguration(installConfig *types.InstallConfig) field.ErrorList {
	var allErrs field.ErrorList
	var fieldPath *field.Path

	var workers int
	for _, worker := range installConfig.Compute {
		workers = workers + int(*worker.Replicas)
	}

	if installConfig.ControlPlane != nil && *installConfig.ControlPlane.Replicas == 1 {
		if workers == 0 {
			allErrs = append(allErrs, field.Invalid(fieldPath, installConfig.Networking.NetworkType, "Only OVNKubernetes network type is allowed for Single Node OpenShift (SNO) cluster"))
		} else {
			fieldPath = field.NewPath("Compute", "Replicas")
			allErrs = append(allErrs, field.Required(fieldPath, fmt.Sprintf("Total number of Compute.Replicas must be 0 when ControlPlane.Replicas is 1 for platform. Found %v", workers)))
		}
	}
	return allErrs
}

// ClusterName returns the name of the cluster, or a default name if no
// InstallConfig is supplied.
func (a *OptionalInstallConfig) ClusterName() string {
	if a.Config != nil && a.Config.ObjectMeta.Name != "" {
		return a.Config.ObjectMeta.Name
	}
	return "agent-cluster"
}

func warnUnusedConfig(installConfig *types.InstallConfig) {
	// "Proxyonly" is the default set from generic install config code
	if installConfig.AdditionalTrustBundlePolicy != "Proxyonly" {
		fieldPath := field.NewPath("AdditionalTrustBundlePolicy")
		logrus.Warnf(fmt.Sprintf("%s: %s is ignored", fieldPath, installConfig.AdditionalTrustBundlePolicy))
	}

	for i, compute := range installConfig.Compute {
		if compute.Hyperthreading != "Enabled" {
			fieldPath := field.NewPath(fmt.Sprintf("Compute[%d]", i), "Hyperthreading")
			logrus.Warnf(fmt.Sprintf("%s: %s is ignored", fieldPath, compute.Hyperthreading))
		}

		if compute.Platform != (types.MachinePoolPlatform{}) {
			fieldPath := field.NewPath(fmt.Sprintf("Compute[%d]", i), "Platform")
			logrus.Warnf(fmt.Sprintf("%s is ignored", fieldPath))
		}
	}

	if installConfig.ControlPlane.Hyperthreading != "Enabled" {
		fieldPath := field.NewPath("ControlPlane", "Hyperthreading")
		logrus.Warnf(fmt.Sprintf("%s: %s is ignored", fieldPath, installConfig.ControlPlane.Hyperthreading))
	}

	if installConfig.ControlPlane.Platform != (types.MachinePoolPlatform{}) {
		fieldPath := field.NewPath("ControlPlane", "Platform")
		logrus.Warnf(fmt.Sprintf("%s is ignored", fieldPath))
	}

	// "External" is the default set from generic install config code
	if installConfig.Publish != "External" {
		fieldPath := field.NewPath("Publish")
		logrus.Warnf(fmt.Sprintf("%s: %s is ignored", fieldPath, installConfig.Publish))
	}
	if installConfig.CredentialsMode != "" {
		fieldPath := field.NewPath("CredentialsMode")
		logrus.Warnf(fmt.Sprintf("%s: %s is ignored", fieldPath, installConfig.CredentialsMode))
	}
	if installConfig.BootstrapInPlace != nil && installConfig.BootstrapInPlace.InstallationDisk != "" {
		fieldPath := field.NewPath("BootstrapInPlace", "InstallationDisk")
		logrus.Warnf(fmt.Sprintf("%s: %s is ignored", fieldPath, installConfig.BootstrapInPlace.InstallationDisk))
	}
}

// GetReplicaCount gets the configured master and worker replicas.
func GetReplicaCount(installConfig *types.InstallConfig) (numMasters, numWorkers int64) {
	numRequiredMasters := int64(0)
	if installConfig.ControlPlane != nil && installConfig.ControlPlane.Replicas != nil {
		numRequiredMasters += *installConfig.ControlPlane.Replicas
	}

	numRequiredWorkers := int64(0)
	for _, worker := range installConfig.Compute {
		if worker.Replicas != nil {
			numRequiredWorkers += *worker.Replicas
		}
	}

	return numRequiredMasters, numRequiredWorkers
}
