package main

import (
	"fmt"
	"os"
	io "preflight/src/io"
	"preflight/src/preflight"
	"preflight/src/programs"
	"preflight/src/styles"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var Remote bool

var rootCmd = &cobra.Command{
	Use:   "preflight [flags] [checklist file]",
	Short: "Automate checklist to ensure you are ready to go",
	Long: fmt.Sprintf(`A small CLI that will run some commands for you, depending on the chosen config. 

Written with %s in %s`, styles.Heart, styles.Golor),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			dataBytes []byte
			err       error
		)
		if Remote {
			dataBytes, err = programs.LoadHttpFileFrom(args[0])
		} else {
			dataBytes, err = io.ReadFile(args[0])
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		systemCheck, err := io.ReadChecklist(dataBytes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sort.SliceStable(systemCheck, func(a, b int) bool {
			return systemCheck[a].Name < systemCheck[b].Name
		})
		if _, err := tea.NewProgram(preflight.InitPreflightModel(systemCheck)).Run(); err != nil {
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
