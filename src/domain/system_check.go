package domain

type SystemCheck struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Optional    bool         `yaml:"optional"`
	Checkpoints []Checkpoint `yaml:"checkpoints"`
	Check       bool
}
