package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DbConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Database string `yaml:"db"`
}

func GetConfig() *DbConfig {
	c := &DbConfig{}
	info, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading data from conf file: %v\n", err)
		os.Exit(1)
	}
	yaml.Unmarshal(info, c)
	return c
}
