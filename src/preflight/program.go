package preflight

import (
	"preflight/src/render"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func (p PreflightModel) Init() tea.Cmd {
	return tea.Batch(
		p.runCheckpoint(),
		p.spinner.Tick,
	)
}

func (p PreflightModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.progress.Width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return p, tea.Quit
		}
	case systemCheckMsg:
		return p.UpdateInternalState(msg)
	case progress.FrameMsg:
		return p.UpdateProgress(msg)
	default:
		p.spinner, cmd = p.spinner.Update(msg)
	}
	return p, cmd
}

func (p PreflightModel) UpdateProgress(msg progress.FrameMsg) (PreflightModel, tea.Cmd) {
	newModel, cmd := p.progress.Update(msg)
	if newModel, ok := newModel.(progress.Model); ok {
		p.progress = newModel
	}
	return p, cmd
}

func (p PreflightModel) View() string {
	view := strings.Builder{}

	if p.done {
		view.WriteString(render.RenderResultFor(p.checks[len(p.checks)-1]))
		view.WriteString(p.RenderConclusion())
		return view.String()
	}

	for i := p.activeIndex; i < len(p.checks); i++ {
		view.WriteString(render.RenderSystemCheck(p.checks[i], i == p.activeIndex, p.spinner))
	}
	view.WriteString("\n")
	view.WriteString(p.progress.View())

	return view.String()
}
