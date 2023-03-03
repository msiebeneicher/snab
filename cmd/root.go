package cmd

import (
	"fmt"
	"os"

	"snab/pkg/logger"

	"github.com/spf13/cobra"
)

var verbose bool
var trace bool

var RootCmd = &cobra.Command{
	Use:   "snab",
	Short: "SnaB - bundle shell script to a powerful modern CLI applications",
	Long:  `SnaB - enable you to bundle shell script to a powerful modern CLI applications`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "set verbose output to debug level")
	RootCmd.PersistentFlags().BoolVarP(&trace, "trace", "", false, "set verbose output to trace level")
	cobra.OnInitialize(handleVerbosity)
}

func handleVerbosity() {
	if trace {
		logger.SetLevelTrace()
	} else if verbose {
		logger.SetLevelDebug()
	}
}
