package command

import (
	"bytes"
	"snab/pkg/snabfile"
	"text/template"

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

// parseFlags will parse setted flags by go-template
func parseFlags(use string, cmd string) (string, error) {
	tpl, err := template.New("cmd").Parse(cmd)
	if err != nil {
		return "", err
	}

	var parsed bytes.Buffer

	// merge task*Flags maps for parsing
	tFlags := map[string]any{}
	for k, v := range tStringFlags[use] {
		tFlags[k] = *v
	}
	for k, v := range tBoolFlags[use] {
		tFlags[k] = *v
	}

	err = tpl.Execute(&parsed, tFlags)
	if err != nil {
		return "", err
	}

	return parsed.String(), nil
}
