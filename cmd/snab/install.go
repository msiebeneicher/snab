package snab

import (
	"fmt"

	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install snab basket under /usr/local/bin",
	Long:  `Install snab basket under /usr/local/bin`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("TODO: install script under /usr/local/bin/foo")
	},
}
