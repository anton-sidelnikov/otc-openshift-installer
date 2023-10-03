package bootkube

import (
	"os"
	"path/filepath"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/templates/content"
)

const (
	kubeCloudConfigFileName = "kube-cloud-config.yaml"
)

var _ asset.WritableAsset = (*KubeCloudConfig)(nil)

// KubeCloudConfig is the constant to represent contents of kube_cloudconfig.yaml file
type KubeCloudConfig struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *KubeCloudConfig) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *KubeCloudConfig) Name() string {
	return "KubeCloudConfig"
}

// Generate generates the actual files by this asset
func (t *KubeCloudConfig) Generate(parents asset.Parents) error {
	fileName := kubeCloudConfigFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *KubeCloudConfig) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *KubeCloudConfig) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, kubeCloudConfigFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
