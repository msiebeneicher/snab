package snab

import (
	"fmt"
	"os"
	"snab/pkg/common"
	"snab/pkg/logger"
	"snab/pkg/snabfile"

	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install snab basket under /usr/local/bin",
	Long:  `Install snab basket under /usr/local/bin`,
	Run: func(cmd *cobra.Command, args []string) {
		c, _ := snabfile.NewSnabConfigByYaml()
		scriptPath := fmt.Sprintf("/usr/local/bin/%s", c.Name)

		fileExists, _ := common.IsFile(scriptPath)
		if fileExists {
			logger.Infof("snab exec file `%s` already exists\n", scriptPath)
		} else {
			script := fmt.Sprintf("#!/usr/bin/env bash\nSNABFILE=%s snab $@", c.Snabfile)
			logger.WithField("script", script).Debugf("creating snab exec file `%s`\n", scriptPath)

			b := []byte(script)
			err := os.WriteFile(scriptPath, b, 0755)
			if err != nil {
				logger.Error(err)
				logger.Fatalf("can't create snab exec file `%s`", scriptPath)
			}

			logger.Infof("snab exec file `%s` successfully created\n", scriptPath)
		}
	},
}
