package content

import (
	"io"
	"path"

	"github.com/anton-sidelnikov/otc-openshift-installer/data"
)

const (
	// TemplateDir is the target directory for all template assets' files
	TemplateDir      = "templates"
	bootkubeDataDir  = "manifests/bootkube/"
	openshiftDataDir = "manifests/openshift/"
)

// GetBootkubeTemplate returns the contents of the file in bootkube data dir
func GetBootkubeTemplate(uri string) ([]byte, error) {
	return getFileContents(path.Join(bootkubeDataDir, uri))
}

// GetOpenshiftTemplate returns the contents of the file in openshift data dir
func GetOpenshiftTemplate(uri string) ([]byte, error) {
	return getFileContents(path.Join(openshiftDataDir, uri))
}

// getFileContents the content of the given URI, assuming that it's a file
func getFileContents(uri string) ([]byte, error) {
	file, err := data.Assets.Open(uri)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
