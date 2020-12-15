package build

import (
	"runtime/debug"
)

// Version is dynamically set by the toolchain.
var Version = "DEV"

// Date is dynamically set at build time.
var Date = "1970-01-01" // YYYY-MM-DD

func init() {
	if Version == "DEV" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
			Version = info.Main.Version
		}
	}
}
