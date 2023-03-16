package cmd

import (
	"fmt"

	"snab/pkg/common"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Print the version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("snab %s\n", common.Version)
	},
}
