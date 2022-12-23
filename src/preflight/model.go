package preflight

import (
	"fmt"
	"preflight/src/domain"
	"preflight/src/render"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

type PreflightModel struct {
	checks                []domain.SystemCheck
	spinner               spinner.Model
	progress              progress.Model
	activeIndex           int
	activeCheckpointIndex int
	done                  bool
}

func (p PreflightModel) getActive() *domain.SystemCheck {
	return &p.checks[p.activeIndex]
}
func (p PreflightModel) getActiveCheckpoint() domain.Checkpoint {
	return p.getActive().Checkpoints[p.activeCheckpointIndex]
}

func InitPreflightModel(systemCheck []domain.SystemCheck) PreflightModel {
	fmt.Println(render.Greetings.String())
	p := progress.New(
		progress.WithGradient(string(render.Ocean), string(render.White)),
	)
	s := spinner.New()
	s.Spinner = spinner.Jump
	s.Style = lipgloss.NewStyle().Foreground(render.Honey)
	return PreflightModel{
		checks:   systemCheck,
		spinner:  s,
		progress: p,
	}
}
