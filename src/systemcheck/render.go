package systemcheck

import (
	"fmt"
	"preflight/src/styles"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
)

func (s SystemCheck) RenderSystemCheck(active bool, spinner spinner.Model) string {
	icon := styles.PkgNameStyle.Render("-")
	checkName := styles.PkgNameStyle.Render(s.Name)

	if active {
		icon = spinner.View()
		checkName = styles.CurrentPkgNameStyle.Render(s.Name)
	}

	return fmt.Sprintf("%s %s\n", icon, checkName)
}

func (s SystemCheck) RenderResult() string {
	icon := styles.CheckMark.String()
	name := styles.CheckMark.Render(s.Name)
	desc := strings.Builder{}

	if !s.Check {
		style := styles.KoMark
		if s.Optional {
			style = styles.WarningMark
		}
		icon = style.String()
		name = style.Render(s.Name)
		desc.WriteString(fmt.Sprintf("\n\t%s", s.Description))
		for _, checkpoint := range s.Checkpoints {
			desc.WriteString(fmt.Sprintf("\n\t%s\t%s", checkpoint.Name, checkpoint.Documentation))
		}
	}

	return fmt.Sprintf("%s %s%s", icon, name, styles.PkgNameStyle.Render(desc.String()))
}
