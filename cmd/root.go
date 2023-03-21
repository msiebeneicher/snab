package cmd

import (
	"fmt"
	"os"

	"snab/cmd/origin"
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
	initRootCmd()
	initSnabCommands()
}

func initRootCmd() {
	c, _ := snabfile.NewSnabConfigByYaml()

	RootCmd = &cobra.Command{
		Use:     c.Name,
		Short:   c.Description.Short,
		Long:    c.Description.Long,
		Example: c.Description.Example,
	}

	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "set verbose output to debug level")
	RootCmd.PersistentFlags().BoolVarP(&trace, "trace", "", false, "set verbose output to trace level")

	// set snabfile path again in flagset to ignore it
	RootCmd.PersistentFlags().String("snabfile", "", "snabfile path which will be ignored here") //nolint:golint,unused
	RootCmd.PersistentFlags().MarkHidden("snabfile")                                             //nolint:golint,errcheck

	cobra.OnInitialize(handleVerbosity)

	command.InitTaskCommands(c.Tasks, c.GetSnabfileDir(), RootCmd)
}

func initSnabCommands() {
	docsGenCmd := origin.InitDocsGenCmd(RootCmd)
	origin.DocsCmd.AddCommand(docsGenCmd)

	OriginCmd.AddCommand(origin.InstallCmd, origin.UninstallCmd, origin.DocsCmd)
	RootCmd.AddCommand(OriginCmd, VersionCmd)
}

func handleVerbosity() {
	if trace {
		logger.SetLevelTrace()
	} else if verbose {
		logger.SetLevelDebug()
	}
}
