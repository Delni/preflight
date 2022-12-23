package domain

type Checkpoint struct {
	Name           string `yaml:"name"`
	Command        string `yaml:"command"`
	Documentation  string `yaml:"documentation"`
	UseInteractive bool   `yaml:"use_interactive"`
}
