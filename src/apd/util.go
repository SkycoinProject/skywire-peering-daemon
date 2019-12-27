package apd

import (
	"fmt"
	"net"

	"github.com/SkycoinProject/skycoin/src/util/logging"
)

var logger = func() *logging.Logger {
	masterLogger := logging.NewMasterLogger()
	return masterLogger.PackageLogger("apd.broadcast")
}

// BroadCastPubKey broadcasts a UDP packet containing the public key of the local visor.
// Broadcasts is sent on the local network broadcasts address.
func BroadCastPubKey(pubkey, broadCastIP string, port int) error {
	address := fmt.Sprintf("%s:%d", broadCastIP, port)
	bAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		logger().Errorf("Couldn't resolve broadcast address: %v", err)
		return err
	}

	conn, err := net.DialUDP("udp", nil, bAddr)
	if err != nil {
		return err
	}

	defer conn.Close()

	packet := []byte(pubkey)
	_, err = conn.Write(packet)
	if err != nil {
		return err
	}

	return nil
}

func getLocalIP() string {
	var localIP string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logger().Errorf("Couldn't get device unicast addresses: %v", err)
		return ""
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIP = ipnet.IP.String()
			}
		}
	}
	return localIP
}
