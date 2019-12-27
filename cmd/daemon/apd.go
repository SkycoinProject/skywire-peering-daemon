package main

import (
	"github.com/SkycoinProject/skycoin/src/cipher"
	"github.com/SkycoinProject/skycoin/src/util/logging"
	"github.com/SkycoinProject/skywire-peering-daemon/src/apd"
)

func main() {
	masterLogger := logging.NewMasterLogger()
	logger := masterLogger.PackageLogger("auto-peering-daemon")
	pubKey, _ := cipher.GenerateKeyPair()
	daemon := apd.NewApd(pubKey.Hex(), logger)

	// Run the daemon
	daemon.Run()
}
