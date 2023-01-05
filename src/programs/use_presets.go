package programs

import (
	"fmt"
	p "preflight/presets"
	"preflight/src/styles"
	"preflight/src/systemcheck"
	"sort"
	"strings"
)

func displayPresetError(err_presets []string) {
	fmt.Println(styles.KoMark.Render("Unknown presets:"))
	for _, err := range err_presets {
		fmt.Println(styles.KoMark.Render("   - " + err))
	}

	fmt.Println(styles.KoMark.Render("\n" + AvailablePresets(p.Presets)))
	fmt.Println(styles.PkgNameStyle.Render("Thinks that something is missing? Please open an issue on https://github.com/delni/preflight/issues/new"))

	fmt.Println()
}

func UsePresets(presetsCandidate []string, knownPresets map[string]systemcheck.SystemCheck) []systemcheck.SystemCheck {
	var err_presets = []string{}
	var systemcheck = []systemcheck.SystemCheck{}
	for _, presetName := range presetsCandidate {
		preset, exists := knownPresets[presetName]
		if !exists {
			err_presets = append(err_presets, presetName)
		} else {
			systemcheck = append(systemcheck, preset)
		}
	}
	if len(err_presets) > 0 {
		displayPresetError(err_presets)
	}
	return systemcheck
}

func AvailablePresets(knownPresets map[string]systemcheck.SystemCheck) string {
	var (
		builder = strings.Builder{}
		keys    = []string{}
	)

	for name := range knownPresets {
		keys = append(keys, name)
	}

	sort.Strings(keys)
	builder.WriteString("Available presets are: ")
	builder.WriteString(strings.Join(keys, ", "))
	return builder.String()
}
