package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Project struct {
		Extension string `yaml:"extension"`
	}
	Modules map[string]struct {
		Description string `yaml:"description"`
		Command     string `yaml:"command"`
		Path        string `yaml:"path"`
		Prefix      string `yaml:"prefix"`
		Suffix      string `yaml:"suffix"`
		Template    string `yaml:"template"`
	}
}

func readConf(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return c, nil
}
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func mkdirAll(p string, template string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}

	if fileExists(filepath.Join(p)) {
		log.Fatalln("File exist")
	}

	fileCreated, err := os.Create(p)

	if err != nil {
		panic(err.Error())
	}

	if len(template) > 0 {
		fileTemplate, err := ioutil.ReadFile(filepath.Join(template))
		if err != nil {
			log.Fatalln(err.Error())
		}
		fileCreated.Write(fileTemplate)
	}

	fmt.Println("File created ", p)

	return fileCreated, nil
}

func main() {
	var rootCmd = &cobra.Command{Use: "myCLI"}
	var yaml, err = readConf("conf.yaml")
	if err != nil {
		panic(err)
	}
	rootCmd.ResetCommands()
	for key := range yaml.Modules {
		vars := yaml.Modules[key]
		var command = key

		command = fmt.Sprintf("make:%s", key)

		if len(vars.Command) > 0 {

			command = fmt.Sprintf("make:%s", vars.Command)

		}

		cmd := &cobra.Command{
			Use:   command,
			Short: vars.Description,
			Args:  cobra.MinimumNArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				filename := args[0]
				file := fmt.Sprintf("%s%s%s.%s", vars.Prefix, filename, vars.Suffix, yaml.Project.Extension)
				// err = makeDirectoryAndFileIfNotExists(vars.Path, file, vars.Template)
				mkdirAll(fmt.Sprintf("%s%s", vars.Path, file), vars.Template)
				if err != nil {
					panic(err.Error())
				}

			},
		}
		rootCmd.AddCommand(cmd)
	}
	rootCmd.AddCommand(versionCmd)
	rootCmd.Execute()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of myCLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("myCLI v0.1")
	},
}
