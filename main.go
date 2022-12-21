package main

import (
	"fmt"
	"os"
	preflight "preflight/src"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	heart = lipgloss.NewStyle().Foreground(lipgloss.Color("161"))
	golor = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8"))
)
var rootCmd = &cobra.Command{
	Use:   "preflight [checklist file]",
	Short: "Automate checklist to ensure you are ready to go",
	Long: fmt.Sprintf(`A small CLI that will run some commands for you, depending on the chosen config. 
Written with %s in %s.`, heart.Render("❤️"), golor.Render("Go")),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		systemCheck, err := preflight.ReadChecklistFile(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sort.SliceStable(systemCheck, func(a, b int) bool {
			return systemCheck[a].Name < systemCheck[b].Name
		})
		if _, err := tea.NewProgram(preflight.PreflighModel(systemCheck)).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
