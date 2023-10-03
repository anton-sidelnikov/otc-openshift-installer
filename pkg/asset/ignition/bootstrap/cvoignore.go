package bootstrap

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/manifests"
	configv1 "github.com/openshift/api/config/v1"
)

var (
	_ asset.WritableAsset = (*CVOIgnore)(nil)
)

const (
	cvoOverridesFilename      = "manifests/cvo-overrides.yaml"
	originalOverridesFilename = "original_cvo_overrides.patch"
)

// CVOIgnore adds bootstrap files needed to inform CVO to ignore resources for which the installer is providing manifests.
type CVOIgnore struct {
	FileList []*asset.File
}

// Name returns a human friendly name for the operator
func (a *CVOIgnore) Name() string {
	return "CVO Ignore"
}

// Dependencies returns all of the dependencies directly needed by the CVOIgnore asset
func (a *CVOIgnore) Dependencies() []asset.Asset {
	return []asset.Asset{
		&manifests.Manifests{},
		&manifests.Openshift{},
	}
}

// Generate generates the respective operator config.yml files
func (a *CVOIgnore) Generate(dependencies asset.Parents) error {
	operators := &manifests.Manifests{}
	openshiftManifests := &manifests.Openshift{}
	dependencies.Get(operators, openshiftManifests)

	var clusterVersion *unstructured.Unstructured
	var ignoredResources []interface{}
	var files []*asset.File
	files = append(files, operators.FileList...)
	files = append(files, openshiftManifests.FileList...)

	seen := make(map[string]string, len(files))
	for _, file := range files {
		u := &unstructured.Unstructured{}
		if err := yaml.Unmarshal(file.Data, u); err != nil {
			return errors.Wrapf(err, "could not unmarshal %q", file.Filename)
		}
		group := u.GetObjectKind().GroupVersionKind().Group
		kind := u.GetKind()
		namespace := u.GetNamespace()
		name := u.GetName()

		key := fmt.Sprintf("%s |! %s |! %s |! %s", group, kind, namespace, name)
		if previousFile, ok := seen[key]; ok {
			return fmt.Errorf("multiple manifests for group %s kind %s namespace %s name %s: %s, %s", group, kind, namespace, name, previousFile, file.Filename)
		}
		seen[key] = file.Filename

		if file.Filename == cvoOverridesFilename {
			clusterVersion = u
			continue
		}
		ignoredResources = append(ignoredResources,
			configv1.ComponentOverride{
				Kind:      kind,
				Group:     group,
				Namespace: namespace,
				Name:      name,
				Unmanaged: true,
			})
	}

	specAsInterface, ok := clusterVersion.Object["spec"]
	if !ok {
		specAsInterface = map[string]interface{}{}
		clusterVersion.Object["spec"] = specAsInterface
	}
	spec, ok := specAsInterface.(map[string]interface{})
	if !ok {
		return errors.Errorf("unexpected type (%T) for .spec in clusterversion", specAsInterface)
	}
	originalOverridesAsInterface := spec["overrides"]
	originalOverrides, ok := originalOverridesAsInterface.([]interface{})
	if !ok && originalOverridesAsInterface != nil {
		return errors.Errorf("unexpected type (%T) for .spec.overrides in clusterversion", originalOverridesAsInterface)
	}
	originalOverridesPatch := map[string]interface{}{
		"spec": map[string]interface{}{
			"overrides": originalOverrides,
		},
	}
	spec["overrides"] = append(ignoredResources, originalOverrides...)

	cvData, err := yaml.Marshal(clusterVersion)
	if err != nil {
		return errors.Wrap(err, "error marshalling clusterversion")
	}
	a.FileList = append(a.FileList, &asset.File{
		Filename: cvoOverridesFilename,
		Data:     cvData,
	})

	origOverrideData, err := json.Marshal(originalOverridesPatch)
	if err != nil {
		return errors.Wrap(err, "error marshalling original overrides")
	}
	a.FileList = append(a.FileList, &asset.File{
		Filename: originalOverridesFilename,
		Data:     origOverrideData,
	})

	return nil
}

// Files returns the files generated by the asset.
func (a *CVOIgnore) Files() []*asset.File {
	return a.FileList
}

// Load does nothing as the file should not be loaded from disk.
func (a *CVOIgnore) Load(f asset.FileFetcher) (bool, error) {
	return false, nil
}
