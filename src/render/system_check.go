package render

import (
	"fmt"
	domain "preflight/src/domain"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
)

func RenderSystemCheck(s domain.SystemCheck, active bool, spinner spinner.Model) string {
	icon := PkgNameStyle.Render("-")
	checkName := PkgNameStyle.Render(s.Name)

	if active {
		icon = spinner.View()
		checkName = CurrentPkgNameStyle.Render(s.Name)
	}

	return fmt.Sprintf("%s %s\n", icon, checkName)
}

func RenderResultFor(s domain.SystemCheck) string {
	icon := CheckMark.String()
	name := CheckMark.Render(s.Name)
	desc := strings.Builder{}

	if !s.Check {
		style := KoMark
		if s.Optional {
			style = WarningMark
		}
		icon = style.String()
		name = style.Render(s.Name)
		desc.WriteString(fmt.Sprintf("\n\t%s", s.Description))
		for _, checkpoint := range s.Checkpoints {
			desc.WriteString(fmt.Sprintf("\n\t%s\t%s", checkpoint.Name, checkpoint.Documentation))
		}
	}

	return fmt.Sprintf("%s %s%s", icon, name, PkgNameStyle.Render(desc.String()))
}
