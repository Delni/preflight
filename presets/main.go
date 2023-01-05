package presets

import "preflight/src/systemcheck"

// Please keep this map in alphabetical order
var Presets = map[string]systemcheck.SystemCheck{
	"node": node,
}
