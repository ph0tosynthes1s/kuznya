package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Read(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %q: %w", path, err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse YAML %q: %w", path, err)
	}

	return &cfg, nil
}
