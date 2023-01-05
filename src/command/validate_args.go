package command

import (
	"errors"

	"github.com/spf13/cobra"
)

func ValidateArgs(cmd *cobra.Command, args []string) error {
	checklists := cmd.Flag("checklists").Value.String()
	if len(args) < 1 && len(checklists) <= 2 {
		return errors.New("requires a path to the checklist file or a list of presets")
	}
	return nil
}
