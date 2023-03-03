package snabfile

import (
	"os"
	"snab/pkg/logger"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version     string      `yaml:"version"`
	Name        string      `yaml:"name"`
	Description Description `yaml:"description"`
	Tasks       Tasks
}

type Tasks map[string]Task

type Task struct {
	Description Description `yaml:"description"`
	Commands    []string    `yaml:"cmds"`
	Dir         string      `yaml:"dir"`
}

type Description struct {
	Short   string `yaml:"short"`
	Long    string `yaml:"long"`
	Example string `yaml:"example"`
}

func NewTasksByYaml(p string) (Config, error) {
	c := Config{}
	yamlFile, err := os.ReadFile(p)

	if err != nil {
		logger.Errorf("Error reading YAML file: %s\n", err)
		return c, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	return c, err
}