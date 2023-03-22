package command

import (
	"bytes"
	"strings"
	"text/template"
)

// parseTaskCommand will parse setted flags by go-template
func parseTaskCommand(use string, cmd string, snabfileDir string) (string, error) {
	tpl, err := template.New("cmd").Parse(cmd)
	if err != nil {
		return "", err
	}

	var parsed bytes.Buffer

	// merge task*Flags maps and defaults for parsing
	tVars := map[string]any{}
	tVars["snabfileDir"] = snabfileDir

	for k, v := range tStringFlags[use] {
		tVars[k] = *v
	}
	for k, v := range tBoolFlags[use] {
		tVars[k] = *v
	}

	err = tpl.Execute(&parsed, tVars)
	if err != nil {
		return "", err
	}

	return strings.Trim(parsed.String(), " \\"), nil
}
