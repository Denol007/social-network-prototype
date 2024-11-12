package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	data, err := os.ReadFile("backend/config/config.yaml")
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
