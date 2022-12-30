package command

import (
	"fmt"
	"os"
	"preflight/src/io"
	"preflight/src/preflight"
	"preflight/src/programs"
	"preflight/src/styles"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var Remote bool

var rootCmd = &cobra.Command{
	Use:   "preflight [flags] [checklist file]",
	Short: fmt.Sprintf("Automate checklist to ensure you are ready to %s 🛫", styles.Golor),
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println()
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

func PreflightCommand() *cobra.Command {
	// Add long description
	rootCmd.Long = makeLongDescription()
	// Add flags
	rootCmd.
		Flags().
		BoolVarP(&Remote, "remote", "r", false, "Fetch your checklist file from a remote server.")

	return rootCmd
}

func makeLongDescription() string {
	builder := strings.Builder{}

	builder.WriteString("A small CLI that will run some commands for you, depending on the chosen config.")
	builder.WriteString("\n\n")
	builder.WriteString(fmt.Sprintf("Written with %s in %s", styles.Heart, styles.Golor))

	return builder.String()
}
