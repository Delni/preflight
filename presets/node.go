package presets

import "preflight/src/systemcheck"

var node systemcheck.SystemCheck = systemcheck.SystemCheck{
	Name:        "Node",
	Description: "",
	Checkpoints: []systemcheck.Checkpoint{
		{
			Name:          "node",
			Command:       "node",
			Documentation: "See installation: https://nodejs.org/",
		},
	},
}
