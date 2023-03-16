package cmd

import (
	"github.com/spf13/cobra"
)

var OriginCmd = &cobra.Command{
	Use:     "origin",
	Aliases: []string{"snab"},
	Short:   "SnaB subcommands",
	Long:    `SnaB subcommands`,
}
