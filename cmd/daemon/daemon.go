package main

import (
	"flag"
	"log"

	spd "github.com/SkycoinProject/skywire-peering-daemon/pkg/daemon"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 3 {
		log.Fatalf("Invalid number of arguments: found %d - requires:\n\n [publickey] [remote address] [named pipe]", len(flag.Args()))
	}

	pubKey := flag.Args()[0]
	lADDR := flag.Args()[1]
	namedPipe := flag.Args()[2]
	daemon := spd.NewDaemon(pubKey, lADDR, namedPipe)

	// Run the daemon
	daemon.Run()
}
