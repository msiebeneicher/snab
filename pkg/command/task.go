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

// maxSubCommandLevel defines the max tries to init sub-commands
const maxSubCommandLevel = 3

// map of all commands
var taskCommands = map[string]*cobra.Command{}

// InitTaskCommands init dynamic commands with flags and args
func InitTaskCommands(t snabfile.Tasks, workingDir string, root *cobra.Command) {
	// init root-commands
	for use, task := range t {
		if task.Parent == "" {
			logger.WithField("task", task).Debugf("init task `%s`", use)

			cmd := newTaskCommand(use, task, workingDir)
			if len(task.Flags) > 0 {
				initFlagsForTask(task, cmd)
			}

			taskCommands[use] = cmd
			root.AddCommand(cmd)
		}
	}

	initTaskSubCommands(t, workingDir, 0)
}

// initTaskSubCommands inits sub-commands defined by task.Parent
// this function works with a simple re-try functionality instead of a parent/child tree
func initTaskSubCommands(t snabfile.Tasks, workingDir string, try int) {
	missParentCommand := false
	logger.Debugf("initTaskSubCommands try `%d`", try)

	// loop through all sub-commands with a parent definition
	for use, task := range t {
		if task.Parent != "" {
			logger.WithField("task", task).Debugf("init subtask `%s`", use)
			_, cmdExists := taskCommands[use]
			if cmdExists {
				// the sub-command already initialized
				continue
			}

			_, parentExists := taskCommands[task.Parent]
			if !parentExists {
				if try == maxSubCommandLevel {
					logger.Warnf("the parent command `%s` of `%s` do not exists. please check your snabfile.", task.Parent, use)
				}
				missParentCommand = true
				continue
			}

			cmd := newTaskCommand(use, task, workingDir)
			if len(task.Flags) > 0 {
				initFlagsForTask(task, cmd)
			}

			// not recommended to define tasks with commands as parent
			if taskCommands[task.Parent].Run != nil {
				logger.Warnf(
					"the parent command `%s` of `%s` has already commands defined. please check your snabfile.",
					task.Parent,
					use,
				)
			}

			taskCommands[use] = cmd
			taskCommands[task.Parent].AddCommand(cmd)
		}
	}

	// break the recursion after maxSubCommandLevel tries
	if missParentCommand {
		if try == maxSubCommandLevel {
			logger.Warnf(
				"max level for subcommands of %d is reached and some commands can't be added to the configured parent. please check your snabfile.",
				try,
			)
			return
		}

		initTaskSubCommands(t, workingDir, try+1)
	}
}

// newTaskCommand returns new &cobra.Command
func newTaskCommand(use string, task snabfile.Task, workingDir string) *cobra.Command {
	c := &cobra.Command{
		Use:     use,
		Short:   task.Description.Short,
		Long:    task.Description.Long,
		Example: task.Description.Example,
	}

	if len(task.Commands) > 0 {
		c.Run = func(cmd *cobra.Command, args []string) {
			execCobraCommand(task, cmd, workingDir, args)
		}
	}

	return c
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

// getExecDirectory validate and return exec directory path
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
