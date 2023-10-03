package main

import (
	"github.com/spf13/cobra"

	"github.com/anton-sidelnikov/otc-openshift-installer/pkg/explain"
)

func newExplainCmd() *cobra.Command {
	return explain.NewCmd()
}
