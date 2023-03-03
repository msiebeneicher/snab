package cmd

import (
	"fmt"
	"os"

	"snab/pkg/command"
	"snab/pkg/logger"
	"snab/pkg/snabfile"

	"github.com/spf13/cobra"
)

var verbose bool
var trace bool

var RootCmd *cobra.Command

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	c, _ := snabfile.NewTasksByYaml("./examples/.snab.yml") //TODO: handle snapfile path (by env var?)

	RootCmd = &cobra.Command{
		Use:     c.Name,
		Short:   c.Description.Short,
		Long:    c.Description.Long,
		Example: c.Description.Example,
	}

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "set verbose output to debug level")
	RootCmd.PersistentFlags().BoolVarP(&trace, "trace", "", false, "set verbose output to trace level")
	cobra.OnInitialize(handleVerbosity)

	command.InitTaskCommands(c.Tasks, RootCmd)
}

func handleVerbosity() {
	if trace {
		logger.SetLevelTrace()
	} else if verbose {
		logger.SetLevelDebug()
	}
}
