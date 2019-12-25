package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SkycoinProject/skywire-peering-daemon/src/apd"
)

func main() {
	shutDownCh := make(chan os.Signal)
	signal.Notify(shutDownCh, syscall.SIGTERM, syscall.SIGINT)

	daemon := apd.NewApd()

	// Run the daemon
	daemon.Run()

	for {
		select {
		case <-daemon.DoneCh:
			os.Exit(1)
		case packet := <-daemon.PacketCh:
			daemon.RegisterPubKey(packet)
		case <-shutDownCh:
			log.Println("Shutting down daemon")
			os.Exit(1)
		}
	}
}
