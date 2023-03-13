package command

import (
	"bytes"
	"text/template"
)

// parseTaskCommand will parse setted flags by go-template
func parseTaskCommand(use string, cmd string) (string, error) {
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
