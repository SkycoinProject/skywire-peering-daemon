package main

import (
	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skywire-peering-daemon/src/apd"
)

func main() {
	pubKey, _ := cipher.GenerateKeyPair()
	daemon := apd.NewApd(pubKey.Hex())

	// Run the daemon
	daemon.Run()
}
