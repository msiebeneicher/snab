package command

import (
	"snab/pkg/snabfile"

	"github.com/spf13/cobra"
)

type taskStringFlags map[string]map[string]*string
type taskBoolFlags map[string]map[string]*bool

// storage of string task flags value pointer tStringFlags["cmd"]["flag"]=*string
var tStringFlags taskStringFlags

// storage of bool task flags value pointer tBoolFlags["cmd"]["flag"]=*bool
var tBoolFlags taskBoolFlags

// init tFlags for global usage
func init() {
	tStringFlags = taskStringFlags{}
	tBoolFlags = taskBoolFlags{}
}

// initFlagsForTask init string flags for task command
func initFlagsForTask(t snabfile.Task, cmd *cobra.Command) {
	tBoolFlags[cmd.Use] = map[string]*bool{}
	tStringFlags[cmd.Use] = map[string]*string{}

	for _, flag := range t.Flags {
		switch fType := flag.Type; fType {
		case "bool":
			tBoolFlags[cmd.Use][flag.Name] = cmd.Flags().BoolP(flag.Name, flag.Shorthand, flag.GetValueAsBoolean(), flag.Usage)
		case "string":
			tStringFlags[cmd.Use][flag.Name] = cmd.Flags().StringP(flag.Name, flag.Shorthand, flag.Value, flag.Usage)
		default:
			tStringFlags[cmd.Use][flag.Name] = cmd.Flags().StringP(flag.Name, flag.Shorthand, flag.Value, flag.Usage)
		}
	}
}
