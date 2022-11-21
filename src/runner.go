package preflight

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	honey               = lipgloss.Color("#febe3c")
	ocean               = lipgloss.Color("#1686cb")
	pkgNameStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(honey)
	checkMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	koMark              = lipgloss.NewStyle().Foreground(lipgloss.Color("160")).SetString("✕")
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
		progress.WithoutPercentage(),
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
		p.spinner.Tick,
	)
}

func (p preflightModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return p, tea.Quit
		}
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
		view.WriteString(checkMark.Render("Done! 🛫\n"))
		return view.String()
	}

	for i := p.activeIndex; i < len(p.checks); i++ {
		view.WriteString(p.checks[i].Render(i == p.activeIndex, p.spinner))
	}

	view.WriteString(p.progress.View())
	view.WriteString(fmt.Sprintf(" %d / %d", p.activeIndex+1, len(p.checks)))

	return view.String()
}
