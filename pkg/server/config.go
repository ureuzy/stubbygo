package server

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	*Endpoints `yaml:"endpoints"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}
	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(b, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
