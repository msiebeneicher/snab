package snabfile

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"snab/pkg/common"
	"snab/pkg/logger"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// default SnaB file name
const snabfileName = ".snab.yml"

type Config struct {
	SchemaVersion int         `yaml:"schema_version"`
	Name          string      `yaml:"name"`
	Version       string      `yaml:"version"`
	Description   Description `yaml:"description"`
	Tasks         Tasks
	Snabfile      string
}

type Tasks map[string]Task

type Task struct {
	Description Description `yaml:"description"`
	Commands    []string    `yaml:"cmds"`
	Flags       []Flag      `yaml:"flags"`
	Dir         string      `yaml:"dir"`
}

type Flag struct {
	Name      string `yaml:"name"`
	Shorthand string `yaml:"shorthand"`
	Usage     string `yaml:"usage"`
	Value     string `yaml:"value"`
}

type Description struct {
	Short   string `yaml:"short"`
	Long    string `yaml:"long"`
	Example string `yaml:"example"`
}

func NewSnabConfigByYaml() (Config, error) {
	c := Config{
		Snabfile: getSnabfilePath(),
	}
	yamlFile, err := os.ReadFile(c.Snabfile)

	if err != nil {
		logger.Errorf("Error reading YAML file: %s\n", err)
		return c, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	return c, err
}

func getSnabfilePathInput() string {

	fs := pflag.NewFlagSet("snab", pflag.ContinueOnError)

	fs.String("snabfile", snabfileName, "Path to your snabfile") //nolint:golint,unused
	fs.SetOutput(io.Discard)                                     //disable output for flags
	fs.Parse(os.Args[1:])                                        //nolint:golint,errcheck

	viper.BindEnv("snabfile") //nolint:golint,errcheck
	viper.BindPFlags(fs)      //nolint:golint,errcheck

	return viper.GetString("snabfile")
}

func getSnabfilePath() string {
	p := getSnabfilePathInput()
	if p == "" {
		logger.Fatalln("Please set a path to your snabfile")
	}

	snabfilePath, err := filepath.Abs(p)
	if err != nil {
		logger.WithField("err", err).Fatalln("snabfile not found: path is not valid")
	}

	if !strings.Contains(snabfilePath, snabfileName) {
		snabfilePath = fmt.Sprintf("%s/%s", snabfilePath, snabfileName)
	}

	isFile, err := common.IsFile(snabfilePath)
	if !isFile || err != nil {
		logger.WithField("err", err).Fatalf("snabfile `%s` not found", snabfilePath)
	}

	return snabfilePath
}
