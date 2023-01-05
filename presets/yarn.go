package presets

import "preflight/src/systemcheck"

var yarn = systemcheck.SystemCheck{
	Name:        "Yarn",
	Description: "Package Manager and tooling for managing your node projects",
	Checkpoints: []systemcheck.Checkpoint{
		{
			Name:          "yarn",
			Command:       "yarn",
			Documentation: "See https://yarnpkg.com",
		},
	},
}
