package main

import (
	"flag"

	"github.com/SkycoinProject/skywire-peering-daemon/src/apd"
)

var (
	pubKey string
)

func main() {
	flag.Parse()
	pubKey := flag.Args()[0]
	namedPipe := flag.Args()[1]
	daemon := apd.NewApd(pubKey, namedPipe)
	daemon.Logger.Infof("$s: %s", pubKey, namedPipe)

	// Run the daemon
	daemon.Run()
}
