package programs

import (
	"fmt"
	"preflight/src/io"
	"preflight/src/styles"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type responseMsg []byte

type errMsg struct{ error }

type loadOverHttpModel struct {
	url      string
	spinner  spinner.Model
	body     []byte
	quitting bool
	err      error
}

func initialModel(url string) loadOverHttpModel {
	s := spinner.New()
	s.Spinner = spinner.MiniDot
	s.Style = lipgloss.NewStyle().Foreground(styles.Honey)
	return loadOverHttpModel{spinner: s, url: url}
}

func (m loadOverHttpModel) Init() tea.Cmd {
	return tea.Batch(m.checkServer, m.spinner.Tick)
}

func (m loadOverHttpModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		m.body = msg
		return m, tea.Quit
	case errMsg:
		m.err = msg
		return m, tea.Tick(time.Duration(2*time.Second), func(t time.Time) tea.Msg {
			m.quitting = true
			return nil
		})

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		if m.quitting {
			return m, tea.Quit
		}

		return m, cmd
	}
}

func (m loadOverHttpModel) View() string {
	str := fmt.Sprintf("%s Fetching file from %s...\n", m.spinner.View(), m.url)
	if m.err != nil {
		return m.err.Error()
	}
	if m.quitting {
		return str + string(m.err.Error()) + "\n"
	}
	if len(m.body) != 0 {
		return string(m.body)
	}
	return str
}

func (m loadOverHttpModel) checkServer() tea.Msg {
	body, err := io.ReadHttpFile(m.url)
	if err != nil {
		return errMsg{err}
	}
	return responseMsg(body)
}

func LoadHttpFileFrom(url string) ([]byte, error) {
	model, err := tea.NewProgram(initialModel(url)).Run()
	if err != nil {
		return nil, err
	}
	return model.(loadOverHttpModel).body, nil
}
