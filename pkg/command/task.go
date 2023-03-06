package command

import (
	"bytes"
	"context"
	"io"
	"os"
	"snab/pkg/common"
	"snab/pkg/logger"
	"snab/pkg/snabfile"

	"github.com/spf13/cobra"
)

// InitTaskCommands init dynamic commands with flags and args
func InitTaskCommands(t snabfile.Tasks, root *cobra.Command) {
	for use, task := range t {
		logger.WithField("task", task).Debugf("init task `%s`", use)

		cmd := newTaskCommand(use, task)
		if len(task.Flags) > 0 {
			initFlagsForTask(task, cmd)
		}

		root.AddCommand(cmd)
	}
}

// newTaskCommand returns new &cobra.Command
func newTaskCommand(use string, task snabfile.Task) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Short:   task.Description.Short,
		Long:    task.Description.Long,
		Example: task.Description.Example,
		Run: func(cmd *cobra.Command, args []string) {
			execCobraCommand(task, cmd, args)
		},
	}
}

// execCobraCommand will used and executed in cobra.Command.Run
func execCobraCommand(task snabfile.Task, cmd *cobra.Command, args []string) {
	ctx := context.Background()

	for _, c := range task.Commands {
		execDir, err := getExecDirectory(task.Dir)
		if err != nil {
			logger.Fatal("error during getting exec directory", err)
		}

		execCmd, err := parseFlags(cmd.Use, c)
		if err != nil {
			logger.Fatal("error during getting exec directory", err)
		}

		logger.WithField("dir", execDir).Debugf("execute now `%s` ..", execCmd)

		var stdoutBuf, stderrBuf bytes.Buffer
		options := common.RunCommandOptions{
			Command: execCmd,
			Dir:     execDir,
			Stdout:  io.MultiWriter(os.Stdout, &stdoutBuf),
			Stderr:  io.MultiWriter(os.Stderr, &stderrBuf),
		}

		err = common.RunCommand(ctx, &options)
		if err != nil {
			logger.Fatal(err)
		}
	}
}

// getExecDirectory validate and return exec diretory path
func getExecDirectory(d string) (string, error) {
	isValid, err := common.IsDirectory(d)
	if !isValid {
		return "", err
	}

	return d, nil
}
