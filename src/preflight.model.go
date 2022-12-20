package preflight

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type PreflightModel struct {
	checks                []SystemCheck
	spinner               spinner.Model
	progress              progress.Model
	activeIndex           int
	activeCheckpointIndex int
	done                  bool
}

func (p PreflightModel) getActive() *SystemCheck {
	return &p.checks[p.activeIndex]
}
func (p PreflightModel) getActiveCheckpoint() Checkpoint {
	return p.getActive().Checkpoints[p.activeCheckpointIndex]
}

type systemCheckMsg struct{ check bool }

func (p PreflightModel) runCheckpoint() tea.Cmd {
	checkpoint := p.getActiveCheckpoint()
	interpreter := GetInterpreterCommand(runtime.GOOS)
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

func (p PreflightModel) UpdateInternalState(msg systemCheckMsg) (tea.Model, tea.Cmd) {
	p.activeCheckpointIndex++
	if msg.check {
		p.getActive().Check = msg.check
	}
	result := p.getActive().RenderResult()
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
			tea.Tick(time.Millisecond*time.Duration(250), func(t time.Time) tea.Msg {
				return p.runCheckpoint()()
			}),
			tea.Printf(result),
		)
	}
	return p, p.runCheckpoint()
}

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
