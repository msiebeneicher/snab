package command

import (
	"os"
	"snab/pkg/snabfile"
	"testing"

	"github.com/spf13/cobra"
)

func TestInitTaskCommands(t *testing.T) {
	cmd := cobra.Command{
		Use: "test",
	}

	task1 := snabfile.Task{
		Parent:      "",
		Description: snabfile.Description{},
		Commands:    []string{},
		Flags:       []snabfile.Flag{},
		Dir:         "",
	}

	task2 := snabfile.Task{
		Parent:      "task1",
		Description: snabfile.Description{},
		Commands:    []string{},
		Flags:       []snabfile.Flag{},
		Dir:         "",
	}

	tasks := snabfile.Tasks{}
	tasks["task1"] = task1
	tasks["task2"] = task2

	InitTaskCommands(tasks, "snabfileDir", &cmd)
}

func TestInitTaskCommandsWithInvalidParent(t *testing.T) {
	cmd := cobra.Command{}
	task1 := snabfile.Task{
		Parent:      "",
		Description: snabfile.Description{},
		Commands:    []string{},
		Flags:       []snabfile.Flag{},
		Dir:         "",
	}

	task2 := snabfile.Task{
		Parent:      "foo",
		Description: snabfile.Description{},
		Commands:    []string{},
		Flags:       []snabfile.Flag{},
		Dir:         "",
	}

	tasks := snabfile.Tasks{}
	tasks["task1"] = task1
	tasks["task2"] = task2

	InitTaskCommands(tasks, "snabfileDir", &cmd)
}

func TestNewTaskCommand(t *testing.T) {
	task1 := snabfile.Task{
		Parent: "",
		Description: snabfile.Description{
			Short:   "short",
			Long:    "long",
			Example: "example",
		},
		Commands: []string{},
		Flags:    []snabfile.Flag{},
		Dir:      "",
	}
	c := newTaskCommand("test", task1, "/tmp")

	if c.Use != "test" {
		t.Errorf("Expected %s, got %s", "test", c.Use)
	}
	if c.Short != "short" {
		t.Errorf("Expected %s, got %s", "short", c.Short)
	}
	if c.Long != "long" {
		t.Errorf("Expected %s, got %s", "long", c.Long)
	}
	if c.Example != "example" {
		t.Errorf("Expected %s, got %s", "example", c.Example)
	}
}

func TestGetExecDirectory(t *testing.T) {
	home, _ := os.UserHomeDir()
	tmp := os.TempDir()

	d, err := getExecDirectory(home, tmp)
	if err != nil {
		t.Error("Expected no error")
	}

	fileInfo, err := os.Stat(d)
	if err != nil {
		t.Errorf("Expected no error in os.Stat for %s", d)
	}
	if !fileInfo.IsDir() {
		t.Errorf("Expected valid dir for %s", d)
	}
}
