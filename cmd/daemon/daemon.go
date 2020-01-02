package main

import (
	"flag"
	"log"

	spd "github.com/SkycoinProject/skywire-peering-daemon/src/daemon"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 3 {
		log.Fatalf("Invalid number of arguments: found %d - requires 3", len(flag.Args()))
	}

	pubKey := flag.Args()[0]
	remoteAddress := flag.Args()[1]
	namedPipe := flag.Args()[2]
	daemon := spd.NewDaemon(pubKey, remoteAddress, namedPipe)

	// Run the daemon
	daemon.Run()
}
