package main

import (
	apd2 "github.com/SkycoinProject/skywire-peering-daemon/src/apd"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	shutDownCh := make(chan os.Signal)
	signal.Notify(shutDownCh, syscall.SIGTERM, syscall.SIGINT)

	apd := apd2.NewApd()

	// Run the daemon
	apd.Run()

	for {
		select {
		case <-apd.DoneCh:
			os.Exit(1)
		case packet := <-apd.PacketCh:
			apd.RegisterPubKey(packet)
		case <-shutDownCh:
			log.Println("Shutting down daemon")
			os.Exit(1)
		}
	}
}
