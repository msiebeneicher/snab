package cmd

import (
	"snab/cmd/snab"

	"github.com/spf13/cobra"
)

func init() {
	docsGenCmd := snab.InitDocsGenCmd(RootCmd)
	snab.DocsCmd.AddCommand(docsGenCmd)

	OriginCmd.AddCommand(snab.InstallCmd, snab.UninstallCmd, snab.DocsCmd)
	RootCmd.AddCommand(OriginCmd)
}

var OriginCmd = &cobra.Command{
	Use:     "origin",
	Aliases: []string{"snab"},
	Short:   "SnaB subcommands",
	Long:    `SnaB subcommands`,
}
