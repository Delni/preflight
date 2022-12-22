package render

import "github.com/charmbracelet/lipgloss"

var (
	Honey               = lipgloss.Color("#febe3c")
	Ocean               = lipgloss.Color("#1686cb")
	White               = lipgloss.Color("#ffffff")
	Greetings           = lipgloss.NewStyle().Foreground(Ocean).SetString("Checking preflight conditions:\n")
	CurrentPkgNameStyle = lipgloss.NewStyle().Foreground(Honey)
	PkgNameStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
	CheckMark           = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	WarningMark         = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).SetString("!")
	KoMark              = lipgloss.NewStyle().Foreground(lipgloss.Color("197")).SetString("✕")
)
