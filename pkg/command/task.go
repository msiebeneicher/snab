package command

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"snab/pkg/common"
	"snab/pkg/logger"
	"snab/pkg/snabfile"
	"strings"

	"github.com/spf13/cobra"
)

// InitTaskCommands init dynamic commands with flags and args
func InitTaskCommands(t snabfile.Tasks, workingDir string, root *cobra.Command) {
	for use, task := range t {
		logger.WithField("task", task).Debugf("init task `%s`", use)

		cmd := newTaskCommand(use, task, workingDir)
		if len(task.Flags) > 0 {
			initFlagsForTask(task, cmd)
		}

		root.AddCommand(cmd)
	}
}

// newTaskCommand returns new &cobra.Command
func newTaskCommand(use string, task snabfile.Task, workingDir string) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Short:   task.Description.Short,
		Long:    task.Description.Long,
		Example: task.Description.Example,
		Run: func(cmd *cobra.Command, args []string) {
			execCobraCommand(task, cmd, workingDir, args)
		},
	}
}

// execCobraCommand will used and executed in cobra.Command.Run
func execCobraCommand(task snabfile.Task, cmd *cobra.Command, workingDir string, args []string) {
	ctx := context.Background()

	for _, c := range task.Commands {
		execDir, err := getExecDirectory(workingDir, task.Dir)
		if err != nil {
			logger.WithField("err", err).Fatalf("error during getting exec directory `%s`", task.Dir)
		}

		// parse flags in command string from snabfile
		execCmd, err := parseTaskCommand(cmd.Use, c)
		if err != nil {
			logger.WithField("err", err).Fatalf("error during getting exec directory `%s`", task.Dir)
		}

		// re-add args to command for execution
		if len(args) > 0 {
			execCmd = fmt.Sprintf("%s %s", execCmd, strings.Join(args, " "))
		}

		logger.WithField("dir", execDir).Debugf("execute `%s` now ..", execCmd)

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
func getExecDirectory(workingDir string, taskDir string) (string, error) {
	if !filepath.IsAbs(taskDir) {
		taskDir = fmt.Sprintf("%s/%s", workingDir, taskDir)
	}

	d, err := filepath.Abs(taskDir)
	if err != nil {
		return "", err
	}

	isValid, err := common.IsDirectory(d)
	if !isValid {
		return "", err
	}

	return d, nil
}
