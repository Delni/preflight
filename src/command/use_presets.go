package command

import (
	"fmt"
	"preflight/src/systemcheck"
)

func UsePresets(presets []string) []systemcheck.SystemCheck {
	fmt.Println(presets)
	return []systemcheck.SystemCheck{}
}
