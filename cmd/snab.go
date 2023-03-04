package cmd

import (
	"snab/cmd/snab"

	"github.com/spf13/cobra"
)

func init() {
	SnabCmd.AddCommand(snab.InstallCmd, snab.UninstallCmd)
	RootCmd.AddCommand(SnabCmd)
}

var SnabCmd = &cobra.Command{
	Use:   "snab",
	Short: "SnaB subcommands",
	Long:  `SnaB subcommands`,
}
