package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skywire-peering-daemon/src"
)

func main() {
	shutDownCh := make(chan os.Signal)
	signal.Notify(shutDownCh, syscall.SIGTERM, syscall.SIGINT)

	apd := src.NewApd()
	pubKey, _ := cipher.GenerateKeyPair()

	apd.PublicKey = pubKey.Hex()

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
