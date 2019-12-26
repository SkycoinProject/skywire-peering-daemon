package main

import (
	"github.com/SkycoinProject/skywire-peering-daemon/src/apd"
)

func main() {
	daemon := apd.NewApd()

	// Run the daemon
	daemon.Run()
}
