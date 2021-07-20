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

func makeDirectoryAndFileIfNotExists(path string, file string, template string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {

		err := os.Mkdir(path, os.ModeDir|0755)

		if err != nil {
			panic(err.Error())
		}

	}

	pathFull := filepath.Join(path, filepath.Base(file))
	if fileExists(pathFull) {
		log.Fatalln("File exist")
	}
	fileCreated, err := os.Create(pathFull)

	if len(template) > 0 {
		fileTemplate, err := ioutil.ReadFile(template)
		if err != nil {
			log.Fatalln(err.Error())
		}
		fileCreated.Write(fileTemplate)

	}

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("File created ", pathFull)

	return nil
}

func main() {
	var rootCmd = &cobra.Command{Use: "myCLI"}
	var yaml, err = readConf("conf.yaml")
	if err != nil {
		panic(err)
	}
	rootCmd.ResetCommands()
	for key, _ := range yaml.Modules {
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
				err = makeDirectoryAndFileIfNotExists(vars.Path, file, vars.Template)
				if err != nil {
					panic(err.Error())
				}

			},
		}
		rootCmd.AddCommand(cmd)
	}

	rootCmd.Execute()
}
