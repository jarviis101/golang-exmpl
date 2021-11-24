package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const Path = "configs/config.yml"

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

type Config struct {
	Server
	Database
}

func GetConfig() Config {
	content, err := ioutil.ReadFile(Path)
	if err != nil {
		fmt.Println(err.Error())
	}
	content = []byte(os.ExpandEnv(string(content)))
	config := Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		fmt.Println(err.Error())
	}
	return config
}
