package systemcheck

import (
	"fmt"
	"preflight/src/render"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
)

func (s SystemCheck) RenderSystemCheck(active bool, spinner spinner.Model) string {
	icon := render.PkgNameStyle.Render("-")
	checkName := render.PkgNameStyle.Render(s.Name)

	if active {
		icon = spinner.View()
		checkName = render.CurrentPkgNameStyle.Render(s.Name)
	}

	return fmt.Sprintf("%s %s\n", icon, checkName)
}

func (s SystemCheck) RenderResult() string {
	icon := render.CheckMark.String()
	name := render.CheckMark.Render(s.Name)
	desc := strings.Builder{}

	if !s.Check {
		style := render.KoMark
		if s.Optional {
			style = render.WarningMark
		}
		icon = style.String()
		name = style.Render(s.Name)
		desc.WriteString(fmt.Sprintf("\n\t%s", s.Description))
		for _, checkpoint := range s.Checkpoints {
			desc.WriteString(fmt.Sprintf("\n\t%s\t%s", checkpoint.Name, checkpoint.Documentation))
		}
	}

	return fmt.Sprintf("%s %s%s", icon, name, render.PkgNameStyle.Render(desc.String()))
}
