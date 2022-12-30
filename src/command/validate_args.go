package command

import (
	"errors"

	"github.com/spf13/cobra"
)

func ValidateArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 1 && len(Checklist) < 1 {
		return errors.New("requires a path to the checklist file or a list of presets")
	}
	return nil
}
