package preflight

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

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

var (
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(honey)
	pkgNameStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	warningMark         = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).SetString("⚠")
	koMark              = lipgloss.NewStyle().Foreground(lipgloss.Color("197")).SetString("✕")
)

func (s SystemCheck) Render(active bool, spinner spinner.Model) string {
	icon := pkgNameStyle.Render("-")
	checkName := pkgNameStyle.Render(s.Name)

	if active {
		icon = spinner.View()
		checkName = currentPkgNameStyle.Render(s.Name)
	}

	return fmt.Sprintf("%s %s\n", icon, checkName)
}

func (s SystemCheck) RenderResult() string {
	icon := checkMark.String()
	name := pkgNameStyle.Render(s.Name)
	desc := strings.Builder{}

	if !s.Check {
		style := koMark
		if s.Optional {
			style = warningMark
		}
		icon = style.String()
		name = style.Render(s.Name)
		desc.WriteString(fmt.Sprintf("\n\t%s", s.Description))
		for _, checkpoint := range s.Checkpoints {
			desc.WriteString(fmt.Sprintf("\n\t%s\t%s", checkpoint.Name, checkpoint.Documentation))
		}
	}

	return fmt.Sprintf("%s %s%s", icon, name, pkgNameStyle.Render(desc.String()))
}
