package preflight

import "preflight/src/render"

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
		return render.KoMark.Render("\n\n No go, no go! Check above for more details. 🛬\n")
	}

	if hasWarning {
		return render.WarningMark.Render("\n\n You're good to go, but check above, some checks were unsuccessful 🎫\n")
	}

	return render.CheckMark.Render("\n\nDone! You're good to go 🛫\n")
}
