package preflight

import (
	"fmt"
	"os"
	"os/exec"
	domain "preflight/src/domain"
	"preflight/src/render"
	"runtime"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
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

type systemCheckMsg struct{ check bool }

func (p PreflightModel) runCheckpoint() tea.Cmd {
	checkpoint := p.getActiveCheckpoint()
	interpreter, err := GetInterpreterCommand(runtime.GOOS)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	interpreterArg := interpreter.InterpreterArgs
	if checkpoint.UseInteractive {
		interpreterArg = interpreter.InterpreterInteractiveArgs
	}
	arg := fmt.Sprintf("%s %s %s", interpreter.Command, interpreter.CommandArgs, checkpoint.Command)
	command := exec.Command(interpreter.Interpreter, interpreterArg, arg)

	// Run check only
	return func() tea.Msg {
		err := command.Run()
		return systemCheckMsg{check: err == nil}
	}
}

func (p PreflightModel) UpdateInternalState(msg systemCheckMsg) (PreflightModel, tea.Cmd) {
	p.activeCheckpointIndex++
	if msg.check {
		p.getActive().Check = msg.check
	}
	result := render.RenderResultFor(*p.getActive())
	if p.activeCheckpointIndex >= len(p.getActive().Checkpoints) {
		p.activeCheckpointIndex = 0
		p.activeIndex++
		if p.activeIndex >= len(p.checks) {
			// Everything's been installed. We're done!
			p.done = true
			return p, tea.Quit
		}
		progressCmd := p.progress.SetPercent(float64(p.activeIndex) / float64(len(p.checks)))
		return p, tea.Batch(
			progressCmd,
			tea.Tick(time.Millisecond*time.Duration(150), func(t time.Time) tea.Msg {
				return p.runCheckpoint()()
			}),
			tea.Printf(result),
		)
	}
	return p, p.runCheckpoint()
}

var (
	checkMark   = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	warningMark = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).SetString("!")
	koMark      = lipgloss.NewStyle().Foreground(lipgloss.Color("197")).SetString("✕")
)

func (p PreflightModel) RenderConclusion() string {
	hasFail := false
	hasWarning := false
	for _, systemCheck := range p.checks {
		if !systemCheck.Check {
			if systemCheck.Optional {
				hasWarning = true
			} else {
				hasFail = true
				break
			}
		}
	}

	if hasFail {
		return koMark.Render("\n\n No go, no go! Check above for more details. 🛬\n")
	}

	if hasWarning {
		return warningMark.Render("\n\n You're good to go, but check above, some checks were unsuccessful 🎫\n")
	}

	return checkMark.Render("\n\nDone! You're good to go 🛫\n")
}
