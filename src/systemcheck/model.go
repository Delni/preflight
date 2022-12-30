package systemcheck

type SystemCheck struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Optional    bool         `yaml:"optional"`
	Checkpoints []Checkpoint `yaml:"checkpoints"`
	Check       bool
}

type Checkpoint struct {
	Name           string `yaml:"name"`
	Command        string `yaml:"command"`
	Documentation  string `yaml:"documentation"`
	UseInteractive bool   `yaml:"use_interactive"`
}
