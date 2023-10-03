package image

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/coreos/stream-metadata-go/arch"
	"github.com/sirupsen/logrus"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/types"
	"github.com/openshift/assisted-image-service/pkg/isoeditor"
)

// AgentPXEFiles is an asset that generates the bootable image used to install clusters.
type AgentPXEFiles struct {
	imageReader          isoeditor.ImageReader
	cpuArch              string
	tmpPath              string
	bootArtifactsBaseURL string
	kernelArgs           string
}

type coreOSKargs struct {
	DefaultKernelArgs string `json:"default"`
}

var _ asset.WritableAsset = (*AgentPXEFiles)(nil)

// Dependencies returns the assets on which the AgentPXEFiles asset depends.
func (a *AgentPXEFiles) Dependencies() []asset.Asset {
	return []asset.Asset{
		&AgentArtifacts{},
	}
}

// Generate generates the image files for PXE asset.
func (a *AgentPXEFiles) Generate(dependencies asset.Parents) error {
	agentArtifacts := &AgentArtifacts{}
	dependencies.Get(agentArtifacts)

	a.tmpPath = agentArtifacts.TmpPath

	ignitionContent := &isoeditor.IgnitionContent{Config: agentArtifacts.IgnitionByte}
	initrdImgPath := filepath.Join(a.tmpPath, "images", "pxeboot", "initrd.img")
	custom, err := isoeditor.NewInitRamFSStreamReader(initrdImgPath, ignitionContent)
	if err != nil {
		return err
	}

	a.imageReader = custom
	a.cpuArch = agentArtifacts.CPUArch
	a.bootArtifactsBaseURL = agentArtifacts.BootArtifactsBaseURL

	kernelArgs, err := getKernelArgs(filepath.Join(a.tmpPath, "coreos", "kargs.json"))
	if err != nil {
		return err
	}
	a.kernelArgs = kernelArgs + string(agentArtifacts.Kargs)
	return nil
}

// PersistToFile writes the PXE assets in the assets folder named pxe.
func (a *AgentPXEFiles) PersistToFile(directory string) error {
	// If the imageReader is not set then it means that either one of the AgentPXEFiles
	// dependencies or the asset itself failed for some reason
	if a.imageReader == nil {
		return errors.New("cannot generate PXE assets due to configuration errors")
	}

	defer a.imageReader.Close()
	bootArtifactsFullPath := filepath.Join(directory, bootArtifactsPath)

	err := createDir(bootArtifactsFullPath)
	if err != nil {
		return err
	}

	err = extractRootFS(bootArtifactsFullPath, a.tmpPath, a.cpuArch)
	if err != nil {
		return err
	}

	agentInitrdFile := filepath.Join(bootArtifactsFullPath, fmt.Sprintf("agent.%s-initrd.img", a.cpuArch))
	err = copyfile(agentInitrdFile, a.imageReader)
	if err != nil {
		return err
	}

	agentVmlinuzFile := filepath.Join(bootArtifactsFullPath, fmt.Sprintf("agent.%s-vmlinuz", a.cpuArch))
	kernelReader, err := os.Open(filepath.Join(a.tmpPath, "images", "pxeboot", "vmlinuz"))
	if err != nil {
		return err
	}
	defer kernelReader.Close()

	if a.cpuArch == arch.RpmArch(types.ArchitectureARM64) {
		gzipReader, err := gzip.NewReader(kernelReader)
		if err != nil {
			panic(err)
		}
		defer gzipReader.Close()
		err = copyfile(agentVmlinuzFile, gzipReader)
		if err != nil {
			return err
		}
	} else {
		err = copyfile(agentVmlinuzFile, kernelReader)
		if err != nil {
			return err
		}
	}

	if a.bootArtifactsBaseURL != "" {
		err = a.createiPXEScript(a.cpuArch, a.bootArtifactsBaseURL, bootArtifactsFullPath, a.kernelArgs)
		if err != nil {
			return err
		}
	}

	logrus.Infof("PXE boot artifacts created in: %s", bootArtifactsFullPath)
	logrus.Infof("Kernel parameters for PXE boot: %s", a.kernelArgs)

	return nil
}

// Name returns the human-friendly name of the asset.
func (a *AgentPXEFiles) Name() string {
	return "Agent Installer PXE Files"
}

// Load returns the PXE image from disk.
func (a *AgentPXEFiles) Load(f asset.FileFetcher) (bool, error) {
	// The PXE image will not be needed by another asset so load is noop.
	// This is implemented because it is required by WritableAsset
	return false, nil
}

// Files returns the files generated by the asset.
func (a *AgentPXEFiles) Files() []*asset.File {
	// Return empty array because File will never be loaded.
	return []*asset.File{}
}

func copyfile(filepath string, src io.Reader) error {
	output, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, src)
	if err != nil {
		return err
	}

	return nil
}

func (a *AgentPXEFiles) createiPXEScript(cpuArch, bootArtifactsBaseURL, pxeAssetsFullPath, kernelArgs string) error {
	iPXEScriptTemplate := `#!ipxe
initrd --name initrd %s/%s
kernel %s/%s initrd=initrd coreos.live.rootfs_url=%s/%s %s
boot
`

	iPXEScript := fmt.Sprintf(iPXEScriptTemplate, bootArtifactsBaseURL,
		fmt.Sprintf("agent.%s-initrd.img", a.cpuArch), bootArtifactsBaseURL,
		fmt.Sprintf("agent.%s-vmlinuz", a.cpuArch), bootArtifactsBaseURL,
		fmt.Sprintf("agent.%s-rootfs.img", a.cpuArch), kernelArgs)

	iPXEFile := fmt.Sprintf("agent.%s.ipxe", a.cpuArch)

	err := os.WriteFile(filepath.Join(pxeAssetsFullPath, iPXEFile), []byte(iPXEScript), 0600)
	if err != nil {
		return err
	}
	logrus.Infof("Created iPXE script %s in %s directory", iPXEFile, pxeAssetsFullPath)

	return nil
}

func getKernelArgs(filepath string) (string, error) {
	kargs, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer kargs.Close()

	data, err := io.ReadAll(kargs)
	if err != nil {
		return "", err
	}

	var args coreOSKargs
	err = json.Unmarshal(data, &args)
	if err != nil {
		return "", err
	}

	// Remove the coreos.liveiso arg
	liveISOArgMatch := regexp.MustCompile(`coreos\.liveiso=[^ ]+ ?`)
	kernelArgs := liveISOArgMatch.ReplaceAllString(args.DefaultKernelArgs, "")
	return kernelArgs, nil
}
