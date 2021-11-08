package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	}
}

func NewConfig(configPath string) (*Config, error) {
	confContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	confContent = []byte(os.ExpandEnv(string(confContent)))
	config := &Config{}

	if err := yaml.Unmarshal(confContent, config); err != nil {
		panic(err)
	}
	return config, nil
}
