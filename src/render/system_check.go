package render

import (
	"fmt"
	domain "preflight/src/domain"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
)

func RenderSystemCheck(s domain.SystemCheck, active bool, spinner spinner.Model) string {
	icon := pkgNameStyle.Render("-")
	checkName := pkgNameStyle.Render(s.Name)

	if active {
		icon = spinner.View()
		checkName = currentPkgNameStyle.Render(s.Name)
	}

	return fmt.Sprintf("%s %s\n", icon, checkName)
}

func RenderResultFor(s domain.SystemCheck) string {
	icon := checkMark.String()
	name := checkMark.Render(s.Name)
	desc := strings.Builder{}

	if !s.Check {
		style := koMark
		if s.Optional {
			style = warningMark
		}
		icon = style.String()
		name = style.Render(s.Name)
		desc.WriteString(fmt.Sprintf("\n\t%s", s.Description))
		for _, checkpoint := range s.Checkpoints {
			desc.WriteString(fmt.Sprintf("\n\t%s\t%s", checkpoint.Name, checkpoint.Documentation))
		}
	}

	return fmt.Sprintf("%s %s%s", icon, name, pkgNameStyle.Render(desc.String()))
}
