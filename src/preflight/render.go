package preflight

import "preflight/src/styles"

func (p PreflightModel) RenderConclusion() string {
	hasFail := false
	hasWarning := false
	for _, systemCheck := range p.checks {
		if !systemCheck.Check {
			if systemCheck.Optional {
				hasWarning = true
			} else {
				hasFail = true
				break
			}
		}
	}

	if hasFail {
		return styles.KoMark.Render("\n\n No go, no go! Check above for more details. 🛬\n")
	}

	if hasWarning {
		return styles.WarningMark.Render("\n\n You're good to go, but check above, some checks were unsuccessful 🎫\n")
	}

	return styles.CheckMark.Render("\n\nDone! You're good to go 🛫\n")
}
