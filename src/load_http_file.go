package render

import (
	"fmt"
	"preflight/src/styles"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type responseMsg []byte

type errMsg struct{ error }

type model struct {
	url      string
	spinner  spinner.Model
	quitting bool
	err      error
}

func initialModel(url string) model {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = lipgloss.NewStyle().Foreground(styles.Honey)
	return model{spinner: s, url: url}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.checkServer, m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case responseMsg:
		return m, nil
	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Loading forever...press q to quit\n\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func (m model) checkServer() tea.Msg {
	// body, _ := preflight.read.ReadHttpFile(m.url)
	// return responseMsg(body)
	return responseMsg([]byte{})
}

func LoadHttpFileProgram(url string) *tea.Program {
	return tea.NewProgram(initialModel(url))
}
