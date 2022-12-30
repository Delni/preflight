package preflight

import (
	"fmt"
	"preflight/src/styles"
	"preflight/src/systemcheck"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

type PreflightModel struct {
	checks                []systemcheck.SystemCheck
	spinner               spinner.Model
	progress              progress.Model
	activeIndex           int
	activeCheckpointIndex int
	done                  bool
}

func (p PreflightModel) getActive() *systemcheck.SystemCheck {
	return &p.checks[p.activeIndex]
}
func (p PreflightModel) getActiveCheckpoint() systemcheck.Checkpoint {
	return p.getActive().Checkpoints[p.activeCheckpointIndex]
}

func InitPreflightModel(systemCheck []systemcheck.SystemCheck) PreflightModel {
	fmt.Println(styles.Greetings.String())
	p := progress.New(
		progress.WithGradient(string(styles.Ocean), string(styles.White)),
	)
	s := spinner.New()
	s.Spinner = spinner.Jump
	s.Style = lipgloss.NewStyle().Foreground(styles.Honey)
	return PreflightModel{
		checks:   systemCheck,
		spinner:  s,
		progress: p,
	}
}
