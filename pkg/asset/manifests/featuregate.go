package manifests

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/installconfig"
	configv1 "github.com/openshift/api/config/v1"
)

var fgFileName = filepath.Join(openshiftManifestDir, "99_feature-gate.yaml")

// FeatureGate generates the feature gate manifest.
type FeatureGate struct {
	FileList []*asset.File
	Config   configv1.FeatureGate
}

var _ asset.WritableAsset = (*Proxy)(nil)

// Name returns a human-friendly name for the asset.
func (*FeatureGate) Name() string {
	return "Feature Gate Config"
}

// Dependencies returns all of the dependencies directly needed to generate
// the asset.
func (*FeatureGate) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.InstallConfig{},
	}
}

// Generate generates the FeatureGate CRD.
func (f *FeatureGate) Generate(dependencies asset.Parents) error {
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(installConfig)

	f.Config = configv1.FeatureGate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: configv1.SchemeGroupVersion.String(),
			Kind:       "FeatureGate",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "cluster",
		},
		Spec: configv1.FeatureGateSpec{
			FeatureGateSelection: configv1.FeatureGateSelection{
				FeatureSet: installConfig.Config.FeatureSet,
			},
		},
	}

	if len(installConfig.Config.FeatureGates) > 0 {
		if installConfig.Config.FeatureSet != configv1.CustomNoUpgrade {
			return errors.Errorf("custom features can only be used with the CustomNoUpgrade feature set")
		}

		customFeatures, err := generateCustomFeatures(installConfig.Config.FeatureGates)
		if err != nil {
			return errors.Wrapf(err, "failed to generate custom features")
		}
		f.Config.Spec.CustomNoUpgrade = customFeatures
	}

	configData, err := yaml.Marshal(f.Config)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s manifests from InstallConfig", f.Name())
	}

	f.FileList = []*asset.File{
		{
			Filename: fgFileName,
			Data:     configData,
		},
	}

	return nil
}

// Files returns the files generated by the asset.
func (f *FeatureGate) Files() []*asset.File {
	return f.FileList
}

// Load loads the already-rendered files back from disk.
func (f *FeatureGate) Load(ff asset.FileFetcher) (bool, error) {
	return false, nil
}

// generateCustomFeatures generates the custom feature gates from the install config.
func generateCustomFeatures(features []string) (*configv1.CustomFeatureGates, error) {
	customFeatures := &configv1.CustomFeatureGates{}

	for _, feature := range features {
		featureName, enabled, err := parseCustomFeatureGate(feature)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to parse custom feature %s", feature)
		}

		if enabled {
			customFeatures.Enabled = append(customFeatures.Enabled, featureName)
		} else {
			customFeatures.Disabled = append(customFeatures.Disabled, featureName)
		}
	}

	return customFeatures, nil
}

// parseCustomFeatureGates parses the custom feature gate string into the feature name and whether it is enabled.
// The expected format is <FeatureName>=<Enabled>.
func parseCustomFeatureGate(rawFeature string) (configv1.FeatureGateName, bool, error) {
	var featureName string
	var enabled bool

	featureParts := strings.Split(rawFeature, "=")
	if len(featureParts) != 2 {
		return "", false, errors.Errorf("feature not in expected format %s", rawFeature)
	}

	featureName = featureParts[0]

	var err error
	enabled, err = strconv.ParseBool(featureParts[1])
	if err != nil {
		return "", false, errors.Wrapf(err, "feature not in expected format %s, could not parse boolean value", rawFeature)
	}

	return configv1.FeatureGateName(featureName), enabled, nil
}
