package preflight

import (
	"fmt"
	"os"
	"os/exec"
	"preflight/src/io"
	"runtime"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type systemCheckMsg struct{ check bool }

func (p PreflightModel) runCheckpoint() tea.Cmd {
	checkpoint := p.getActiveCheckpoint()
	interpreter, err := io.GetInterpreterCommand(runtime.GOOS)
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
			tea.Tick(time.Millisecond*time.Duration(150), func(t time.Time) tea.Msg {
				return p.runCheckpoint()()
			}),
			tea.Printf(result),
		)
	}
	return p, p.runCheckpoint()
}
