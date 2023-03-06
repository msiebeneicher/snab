package command

import (
	"bytes"
	"snab/pkg/snabfile"
	"text/template"

	"github.com/spf13/cobra"
)

type taskFlags map[string]map[string]*string

// storage of command flags value pointer tFlags["cmd"]["flag"]=*string
var tFlags taskFlags

// init tFlags for global usage
func init() {
	tFlags = taskFlags{}
}

// initFlagsForTask init string flags for task command
func initFlagsForTask(t snabfile.Task, cmd *cobra.Command) {
	tFlags[cmd.Use] = map[string]*string{}
	for _, flag := range t.Flags {
		tFlags[cmd.Use][flag.Name] = cmd.Flags().StringP(flag.Name, flag.Shorthand, flag.Value, flag.Usage)
	}
}

// parseFlags will parse setted flags by go-template
func parseFlags(use string, cmd string) (string, error) {
	tpl, err := template.New("cmd").Parse(cmd)
	if err != nil {
		return "", err
	}

	var parsed bytes.Buffer
	err = tpl.Execute(&parsed, tFlags[use])
	if err != nil {
		return "", err
	}

	return parsed.String(), nil
}
