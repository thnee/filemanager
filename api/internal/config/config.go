package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type FileArea struct {
	Id   string `yaml:"id"`
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type FileAreas []FileArea

type User struct {
	Username      string   `yaml:"username"`
	PasswordClear string   `yaml:"password-clear"`
	FileAreas     []string `yaml:"file-areas"`
}

type Users []User

type ConfigSchema struct {
	FileAreas FileAreas `yaml:"file-areas"`
	Users     Users     `yaml:"users"`
}

var Config ConfigSchema

func ReadConfig(fileName string) error {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}

	return nil
}
