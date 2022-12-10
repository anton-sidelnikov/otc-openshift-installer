package image

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/openshift/assisted-image-service/pkg/isoeditor"
	"github.com/openshift/installer/pkg/asset"
)

const (
	agentISOFilename = "agent.%s.iso"
)

// AgentImage is an asset that generates the bootable image used to install clusters.
type AgentImage struct {
	imageReader  isoeditor.ImageReader
	cpuArch      string
	rendezvousIP string
}

var _ asset.WritableAsset = (*AgentImage)(nil)

// Dependencies returns the assets on which the Bootstrap asset depends.
func (a *AgentImage) Dependencies() []asset.Asset {
	return []asset.Asset{
		&Ignition{},
		&BaseIso{},
	}
}

// Generate generates the image file for to ISO asset.
func (a *AgentImage) Generate(dependencies asset.Parents) error {
	ignition := &Ignition{}
	dependencies.Get(ignition)

	baseImage := &BaseIso{}
	dependencies.Get(baseImage)

	ignitionByte, err := json.Marshal(ignition.Config)
	if err != nil {
		return err
	}

	ignitionContent := &isoeditor.IgnitionContent{Config: ignitionByte}
	custom, err := isoeditor.NewRHCOSStreamReader(baseImage.File.Filename, ignitionContent, nil)
	if err != nil {
		return err
	}

	a.imageReader = custom
	a.cpuArch = ignition.CPUArch
	a.rendezvousIP = ignition.RendezvousIP

	return nil
}

// PersistToFile writes the iso image in the assets folder
func (a *AgentImage) PersistToFile(directory string) error {
	// If the imageReader is not set then it means that either one of the AgentImage
	// dependencies or the asset itself failed for some reason
	if a.imageReader == nil {
		return errors.New("cannot generate ISO image due to configuration errors")
	}

	defer a.imageReader.Close()
	agentIsoFile := filepath.Join(directory, fmt.Sprintf(agentISOFilename, a.cpuArch))

	// Remove symlink if it exists
	os.Remove(agentIsoFile)

	output, err := os.Create(agentIsoFile)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, a.imageReader)
	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(directory, "rendezvousIP"), []byte(a.rendezvousIP), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Name returns the human-friendly name of the asset.
func (a *AgentImage) Name() string {
	return "Agent Installer ISO"
}

// Load returns the ISO from disk.
func (a *AgentImage) Load(f asset.FileFetcher) (bool, error) {
	// The ISO will not be needed by another asset so load is noop.
	// This is implemented because it is required by WritableAsset
	return false, nil
}

// Files returns the files generated by the asset.
func (a *AgentImage) Files() []*asset.File {
	// Return empty array because File will never be loaded.
	return []*asset.File{}
}
