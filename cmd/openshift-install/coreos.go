package main

import (
	"github.com/spf13/cobra"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/coreoscli"
)

func newCoreOSCmd() *cobra.Command {
	return coreoscli.NewCmd()
}
