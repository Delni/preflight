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

// https://cobra.dev
var (
	heart = lipgloss.NewStyle().Foreground(lipgloss.Color("161"))
	golor = lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8"))
)
var Remote bool

var rootCmd = &cobra.Command{
	Use:   "preflight [flags] [checklist file]",
	Short: "Automate checklist to ensure you are ready to go",
	Long: fmt.Sprintf(`A small CLI that will run some commands for you, depending on the chosen config. 
Written with %s in %s.`, heart.Render("❤️"), golor.Render("Go")),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			dataBytes []byte
			err       error
		)
		if Remote {
			dataBytes, err = preflight.ReadHttpFile(args[0])
		} else {
			dataBytes, err = preflight.ReadFile(args[0])
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		systemCheck, err := preflight.ReadChecklist(dataBytes)
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
	rootCmd.Flags().BoolVarP(&Remote, "remote", "r", false, "Fetch your checklist file from a remote server.")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
