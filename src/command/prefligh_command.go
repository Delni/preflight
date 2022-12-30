package command

import (
	"fmt"
	"preflight/src/styles"
	"strings"

	"github.com/spf13/cobra"
)

var Remote bool

var rootCmd = &cobra.Command{
	Use:   "preflight [flags] [checklist file]",
	Short: fmt.Sprintf("Automate checklist to ensure you are ready to %s 🛫", styles.Golor),
	Args:  cobra.ExactArgs(1),
	Run:   Run,
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

	builder.WriteString("A small CLI that will run some commands for you, depending on the chosen config, to make sure you are ready to go 🛫")
	builder.WriteString("\n\n")
	builder.WriteString(fmt.Sprintf("Written with %s in %s", styles.Heart, styles.Golor))

	return builder.String()
}
