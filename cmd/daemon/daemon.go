package main

import (
	"flag"
	"log"

	spd "github.com/SkycoinProject/skywire-peering-daemon/src/daemon"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		log.Fatalf("Invalid number of arguments: found %d - requires 2", len(flag.Args()))
	}

	pubKey := flag.Args()[0]
	namedPipe := flag.Args()[1]
	daemon := spd.NewDaemon(pubKey, namedPipe)

	// Run the daemon
	daemon.Run()
}
