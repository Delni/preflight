package programs

import (
	"fmt"
	p "preflight/presets"
	"preflight/src/styles"
	"preflight/src/systemcheck"
	"strings"
)

func UsePresets(presets []string) []systemcheck.SystemCheck {
	var err_presets = []string{}
	var systemcheck = []systemcheck.SystemCheck{}
	for _, presetName := range presets {
		preset, exists := p.Presets[presetName]
		if !exists {
			err_presets = append(err_presets, presetName)
		} else {
			systemcheck = append(systemcheck, preset)
		}
	}
	if len(err_presets) > 0 {
		fmt.Println(styles.KoMark.Render("Unknown presets:"))
		for _, err := range err_presets {
			fmt.Println(styles.KoMark.Render("   - " + err))
		}

		fmt.Println(styles.KoMark.Render("\n" + AvailablePresets()))
		fmt.Println(styles.PkgNameStyle.Render("Thinks that something is missing? Please open an issue on https://github.com/delni/preflight/issues/new"))

		fmt.Println()
	}
	return systemcheck
}

func AvailablePresets() string {
	var (
		builder = strings.Builder{}
		keys    = []string{}
	)
	for name := range p.Presets {
		keys = append(keys, name)
	}
	builder.WriteString("Available presets are: ")
	builder.WriteString(strings.Join(keys, ", "))
	return builder.String()
}
