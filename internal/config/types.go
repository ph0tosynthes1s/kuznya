package config

type Config struct {
	Version      int           `yaml:"version"`
	Project      Project       `yaml:"project"`
	Instructions []Instruction `yaml:"instructions"`
}

type Project struct {
	Name   string `yaml:"name"`
	Output string `yaml:"output"`
}

type Instruction struct {
	Type    string `yaml:"type"`
	Path    string `yaml:"path"`
	Content string `yaml:"content,omitempty"`
}
