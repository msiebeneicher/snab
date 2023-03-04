package snab

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall snab basket from /usr/local/bin",
	Long:  `Uninstall snab basket from /usr/local/bin`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("TODO: uninstall script under /usr/local/bin/foo")
	},
}
