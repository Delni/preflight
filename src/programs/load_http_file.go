package programs

import (
	"fmt"
	"preflight/src/io"
	"preflight/src/styles"

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

func (m loadOverHttpModel) getSentence(prefix interface{}, done bool) string {
	verb := "Fetching"
	if done {
		verb = "Fetched"
	}
	sentence := styles.PkgNameStyle.Render(fmt.Sprintf("%s file from %s", verb, m.url))
	return styles.PkgNameStyle.Render(fmt.Sprintf("%s %s\n", prefix, sentence))
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

		m.quitting = true

		return m, tea.Printf(m.getSentence(styles.CheckMark, true))
	case errMsg:
		m.err = msg
		m.quitting = true
		return m, nil

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
	str := m.getSentence(m.spinner.View(), false)
	if m.err != nil {
		return m.err.Error()
	}
	if len(m.body) != 0 {
		return ""
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
