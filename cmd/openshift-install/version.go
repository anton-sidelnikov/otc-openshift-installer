package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/asset/releaseimage"
	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/version"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Long:  "",
		Args:  cobra.ExactArgs(0),
		RunE:  runVersionCmd,
	}
}

func runVersionCmd(cmd *cobra.Command, args []string) error {
	versionString, err := version.Version()
	if err != nil {
		return err
	}

	fmt.Printf("%s %s\n", os.Args[0], versionString)
	if version.Commit != "" {
		fmt.Printf("built from commit %s\n", version.Commit)
	}
	if image, err := releaseimage.Default(); err == nil {
		fmt.Printf("release image %s\n", image)
	}
	fmt.Printf("release architecture %s\n", version.DefaultArch())
	return nil
}
