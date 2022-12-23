package preflight

import (
	"fmt"
	domain "preflight/src/domain"
	"preflight/src/render"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	honey     = lipgloss.Color("#febe3c")
	ocean     = lipgloss.Color("#1686cb")
	white     = lipgloss.Color("#ffffff")
	greetings = lipgloss.NewStyle().Foreground(ocean).SetString("Checking preflight conditions:\n")
)

func InitPreflightModel(systemCheck []domain.SystemCheck) PreflightModel {
	fmt.Println(greetings.String())
	p := progress.New(
		progress.WithGradient(string(ocean), string(white)),
	)
	s := spinner.New()
	s.Spinner = spinner.Jump
	s.Style = lipgloss.NewStyle().Foreground(honey)
	return PreflightModel{
		checks:   systemCheck,
		spinner:  s,
		progress: p,
	}
}

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
