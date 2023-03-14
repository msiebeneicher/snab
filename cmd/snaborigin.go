package cmd

import (
	"snab/cmd/origin"

	"github.com/spf13/cobra"
)

func init() {
	docsGenCmd := origin.InitDocsGenCmd(RootCmd)
	origin.DocsCmd.AddCommand(docsGenCmd)

	OriginCmd.AddCommand(origin.InstallCmd, origin.UninstallCmd, origin.DocsCmd)
	RootCmd.AddCommand(OriginCmd)
}

var OriginCmd = &cobra.Command{
	Use:     "origin",
	Aliases: []string{"snab"},
	Short:   "SnaB subcommands",
	Long:    `SnaB subcommands`,
}
