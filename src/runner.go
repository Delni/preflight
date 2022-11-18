package preflight

import tea "github.com/charmbracelet/bubbletea"

type preflightModel struct {
}

func PreflighModel(systemCheck []SystemCheck) preflightModel {
	return preflightModel{}
}

func (p preflightModel) Init() tea.Cmd {
	return nil
}

func (p preflightModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return p, tea.Quit
		}
	}
	return p, nil
}

func (p preflightModel) View() string {
	return ``
}
