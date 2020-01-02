package main

import (
	"flag"
	"github.com/SkycoinProject/skywire-peering-daemon/src/apd"
)

func main() {
	flag.Parse()
	pubKey := flag.Args()[0]
	namedPipe := flag.Args()[1]
	daemon := apd.NewDaemon(pubKey, namedPipe)

	// Run the daemon
	daemon.Run()
}
