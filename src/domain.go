package main

type Checkpoint struct {
	Name          string `yaml:"name"`
	Command       string `yaml:"command"`
	Documentation string `yaml:"documentation"`
	Check         bool
}

type SystemCheck struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Optional    bool         `yaml:"optional"`
	Checkpoints []Checkpoint `yaml:"options"`
	Check       bool
}
