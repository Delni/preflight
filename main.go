package main

import (
	"fmt"
	"os"
	preflight "preflight/src"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "preflight [checklist file]",
	Short: "Automate checklist to ensure you are ready to go",
	Long: `A small CLI that will run some commands for you, depending on the chosen config. 
Written with ❤️ in Go.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		systemCheck := preflight.ReadChecklistFile(args[0])
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
