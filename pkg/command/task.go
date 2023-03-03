package command

import (
	"snab/pkg/common"
	"snab/pkg/logger"
	"snab/pkg/snabfile"

	"github.com/spf13/cobra"
)

func InitTaskCommands(t snabfile.Tasks, root *cobra.Command) {
	for use, task := range t {
		logger.Info(use, ":Task:", task.Description.Short)

		root.AddCommand(
			newTaskCommand(use, task),
		)
	}
}

func newTaskCommand(use string, task snabfile.Task) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Short:   task.Description.Short,
		Long:    task.Description.Long,
		Example: task.Description.Example,
		Run: func(cmd *cobra.Command, args []string) {
			for _, c := range task.Commands {
				execDir, err := getExecDirectory(task.Dir)
				if err != nil {
					logger.Fatal(err)
				}

				logger.Debugf("exec %s::%s ..", use, c)
				common.Exec(c, []string{}, execDir, "")
			}
		},
	}
}

// TODO: add proper base dir handling
func getExecDirectory(d string) (string, error) {
	isValid, err := common.IsDirectory(d)
	if !isValid {
		return "", err
	}

	return d, nil
}
