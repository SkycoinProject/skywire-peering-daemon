package main

import (
	"flag"
	apd "github.com/SkycoinProject/skywire-peering-daemon/src/daemon"
)

func main() {
	flag.Parse()
	pubKey := flag.Args()[0]
	namedPipe := flag.Args()[1]
	daemon := apd.NewDaemon(pubKey, namedPipe)

	// Run the daemon
	daemon.Run()
}
