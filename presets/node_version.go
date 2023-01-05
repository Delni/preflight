package presets

import "preflight/src/systemcheck"

var nodeVersion systemcheck.SystemCheck = systemcheck.SystemCheck{
	Name:        "Node Version Manager",
	Description: "tools to be able to seamlessly change your current node/npm version",
	Checkpoints: []systemcheck.Checkpoint{
		{
			Name:          "nvm",
			Command:       "nvm",
			Documentation: "See installation on github: https://github.com/nvm-sh/nvm#installing-and-updating",
		},
		{
			Name:          "n",
			Command:       "n",
			Documentation: "See installation on github: https://github.com/tj/n#installation",
		},
		{
			Name:          "volta",
			Command:       "volta",
			Documentation: "See https://volta.sh",
		},
	},
}
