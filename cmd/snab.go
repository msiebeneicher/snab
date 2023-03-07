package cmd

import (
	"snab/cmd/snab"

	"github.com/spf13/cobra"
)

func init() {
	docsGenCmd := snab.InitDocsGenCmd(RootCmd)
	snab.DocsCmd.AddCommand(docsGenCmd)

	SnabCmd.AddCommand(snab.InstallCmd, snab.UninstallCmd, snab.DocsCmd)
	RootCmd.AddCommand(SnabCmd)
}

var SnabCmd = &cobra.Command{
	Use:   "snab",
	Short: "SnaB subcommands",
	Long:  `SnaB subcommands`,
}
