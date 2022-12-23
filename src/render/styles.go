package render

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	ocean = lipgloss.Color("#1686cb")
	Honey = lipgloss.Color("#febe3c")
	White = lipgloss.Color("#ffffff")
	// Writings
	greetings           = lipgloss.NewStyle().Foreground(ocean).SetString("Checking preflight conditions:\n")
	currentPkgNameStyle = lipgloss.NewStyle().Foreground(Honey)
	pkgNameStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("242"))
	// Icons
	checkMark   = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
	warningMark = lipgloss.NewStyle().Foreground(lipgloss.Color("214")).SetString("!")
	koMark      = lipgloss.NewStyle().Foreground(lipgloss.Color("197")).SetString("✕")
)
