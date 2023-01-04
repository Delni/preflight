package command

import (
	"fmt"
	"os"
	"preflight/src/preflight"
	"preflight/src/programs"
	"preflight/src/styles"
	"preflight/src/systemcheck"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Println()

	var systemChecks []systemcheck.SystemCheck

	if cmd.Flags().Lookup("checklists").Changed {
		systemChecks = programs.UsePresets(strings.Split(Checklist[0], ","))
	} else {
		systemChecks = ReadFile(args[0])
	}

	sort.SliceStable(systemChecks, func(a, b int) bool {
		return systemChecks[a].Name < systemChecks[b].Name
	})

	if len(systemChecks) == 0 {
		fmt.Println(styles.WarningMark.String() + styles.WarningMark.Render(" Your checklist is empty! Weird but why not?"))
		fmt.Println(styles.CheckMark.Render("Done! You're good to go 🛫"))
		os.Exit(0)
	}

	if _, err := tea.NewProgram(preflight.InitPreflightModel(systemChecks)).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
