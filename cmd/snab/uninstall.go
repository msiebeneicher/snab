package snab

import (
	"fmt"
	"os"
	"snab/pkg/common"
	"snab/pkg/logger"
	"snab/pkg/snabfile"

	"github.com/spf13/cobra"
)

var UninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall snab basket from /usr/local/bin",
	Long:  `Uninstall snab basket from /usr/local/bin`,
	Run: func(cmd *cobra.Command, args []string) {
		c, _ := snabfile.NewSnabConfigByYaml()
		scriptPath := fmt.Sprintf("/usr/local/bin/%s", c.Name)

		fileExists, _ := common.IsFile(scriptPath)
		if !fileExists {
			logger.Infof("snab exec file `%s` don't exists\n", scriptPath)
		} else {
			err := os.Remove(scriptPath)
			if err != nil {
				logger.Error(err)
				logger.Fatalf("can't remove snab exec file `%s`\n", scriptPath)
			}

			logger.Infof("snab exec file `%s` successfully removed\n", scriptPath)
		}
	},
}
