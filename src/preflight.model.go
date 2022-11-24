package preflight

import (
	"os/exec"
	"strings"
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
	cmd_raw := strings.Split(checkpoint.Command, " ")
	cmd_args := []string{"command", "-v"}
	cmd_args = append(cmd_args, cmd_raw...)

	if checkpoint.LiveRun {
		cmd_args = strings.Split(checkpoint.Command, " ")
	}
	args := []string{}
	if len(cmd_args) > 1 {
		args = cmd_args[1 : len(cmd_args)-1]
	}
	command := exec.Command(cmd_args[0], args...)

	if checkpoint.LiveRun {
		// Run in a blocking fashion way

		return tea.Batch(
			tea.EnterAltScreen,
			tea.ExecProcess(command, func(err error) tea.Msg {
				return systemCheckMsg{check: err == nil}
			}),
			tea.ExitAltScreen,
		)
	}
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
