package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		DB   `yaml:"db"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	DB struct {
		Host     string `yaml:"host"`
		DBPort   string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBname   string `yaml:"dbname"`
	}
)

func NewConfig() (*Config, error) {

	cfg := &Config{}

	file, err := os.ReadFile("../../config/config.yml")

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal([]byte(file), cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
