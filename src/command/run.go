package command

import (
	"fmt"
	"os"
	"preflight/src/preflight"
	"preflight/src/systemcheck"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println()

	var systemChecks []systemcheck.SystemCheck

	systemChecks = ReadFile(args[0])

	sort.SliceStable(systemChecks, func(a, b int) bool {
		return systemChecks[a].Name < systemChecks[b].Name
	})

	if _, err := tea.NewProgram(preflight.InitPreflightModel(systemChecks)).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
