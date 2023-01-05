package styles

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	Ocean = lipgloss.Color("#1686cb")
	Honey = lipgloss.Color("#febe3c")
	White = lipgloss.Color("#ffffff")
	// Writings
	Greetings           = lipgloss.NewStyle().Foreground(Ocean).SetString("Checking preflight conditions:\n")
	CurrentPkgNameStyle = lipgloss.NewStyle().Foreground(Honey)
	PkgNameStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	// Icons
	Heart       = lipgloss.NewStyle().Foreground(lipgloss.Color("161")).SetString("❤️")
	CheckMark   = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	WarningMark = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).SetString("!")
	KoMark      = lipgloss.NewStyle().Foreground(lipgloss.Color("197")).SetString("✕")
	Golor       = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8")).SetString("Go")
)
