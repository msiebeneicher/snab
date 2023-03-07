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
	Type      string `yaml:"type"`
}

type Description struct {
	Short   string `yaml:"short"`
	Long    string `yaml:"long"`
	Example string `yaml:"example"`
}

// GetWorkingDir returns the path of the used snabfile as working directory
func (c *Config) GetWorkingDir() string {
	return filepath.Dir(c.Snabfile)
}

// GetValueAsBoolean returns the Flag.Value as bool
func (f *Flag) GetValueAsBoolean() bool {
	if f.Value == "true" {
		return true
	}

	if f.Value == "false" {
		return false
	}

	logger.Warnf("unable to convert value `%s` as clear boolean. converted to `false` now.", f.Value)
	return false
}

// NewSnabConfigByYaml returns snabfile.Config
func NewSnabConfigByYaml() (Config, error) {
	snabfile, err := getSnabfilePath()
	c := Config{
		Snabfile: snabfile,
	}

	if snabfile == "" {
		logger.Debug(err)
		logger.Warnln("no snabfile set")
		return c, nil
	}

	yamlFile, err := os.ReadFile(c.Snabfile)
	if err != nil {
		logger.Errorf("error reading YAML file: %s\n", err)
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

func getSnabfilePath() (string, error) {
	p := getSnabfilePathInput()
	if p == "" {
		return "", fmt.Errorf("please set a path to your snabfile")
	}

	snabfilePath, err := filepath.Abs(p)
	if err != nil {
		return "", fmt.Errorf("snabfile not found: path is not valid: %s", err)
	}

	if !strings.Contains(snabfilePath, snabfileName) {
		snabfilePath = fmt.Sprintf("%s/%s", snabfilePath, snabfileName)
	}

	isFile, err := common.IsFile(snabfilePath)
	if !isFile || err != nil {
		return "", fmt.Errorf("snabfile `%s` not found", snabfilePath)
	}

	return snabfilePath, nil
}
