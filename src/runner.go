package preflight

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	honey = lipgloss.Color("#febe3c")
	ocean = lipgloss.Color("#1686cb")
)

type preflightModel struct {
	checks      []SystemCheck
	spinner     spinner.Model
	progress    progress.Model
	activeIndex int
	done        bool
}

func PreflighModel(systemCheck []SystemCheck) preflightModel {
	fmt.Println(lipgloss.NewStyle().Foreground(ocean).Render("Checking preflight conditions:"))
	p := progress.New(
		progress.WithDefaultGradient(),
		// progress.WithWidth(40),
	)
	s := spinner.New()
	s.Spinner = spinner.Jump
	s.Style = lipgloss.NewStyle().Foreground(honey)
	return preflightModel{
		checks:   systemCheck,
		spinner:  s,
		progress: p,
	}
}

func (p preflightModel) Init() tea.Cmd {
	return tea.Batch(
		p.checks[p.activeIndex].run(),
		p.spinner.Tick,
	)
}

func (p preflightModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		if p.activeIndex >= len(p.checks)-1 {
			// Everything's been installed. We're done!
			p.done = true
			return p, tea.Quit
		}
		p.checks[p.activeIndex].Check = msg.check
		p.activeIndex++
		// Update progress bar
		progressCmd := p.progress.SetPercent(float64(p.activeIndex) / float64(len(p.checks)))
		return p, tea.Batch(
			progressCmd,
			tea.Printf(p.checks[p.activeIndex-1].RenderResult()),
			p.checks[p.activeIndex].run(),
		)
	case progress.FrameMsg:
		newModel, cmd := p.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			p.progress = newModel
		}
		return p, cmd
	default:
		p.spinner, cmd = p.spinner.Update(msg)
	}
	return p, cmd
}

func (p preflightModel) View() string {
	view := strings.Builder{}

	if p.done {
		view.WriteString(p.checks[len(p.checks)-1].RenderResult())
		view.WriteString(checkMark.Render("\nDone! 🛫\n"))
		return view.String()
	}

	for i := p.activeIndex; i < len(p.checks); i++ {
		view.WriteString(p.checks[i].Render(i == p.activeIndex, p.spinner))
	}

	view.WriteString(p.progress.View())

	return view.String()
}

type systemCheckMsg struct{ check bool }

func (s SystemCheck) run() tea.Cmd {
	for _, checkpoint := range s.Checkpoints {
		tea.Println(checkpoint.Name)
	}
	d := time.Millisecond * time.Duration(500)
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return systemCheckMsg{check: false}
	})
}
