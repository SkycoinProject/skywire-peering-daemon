package apd

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SkycoinProject/skycoin/src/util/logging"
)

const (
	defaultBroadCastIP = "255.255.255.255"
	port               = 3000
	packetLength       = 10
)

type Packet struct {
	PublicKey string
	IP        string
}

type APD struct {
	PublicKey string
	LocalIP   string
	PacketMap map[string]string
	DoneCh    chan error
	PacketCh  chan Packet
	logger    *logging.Logger
}

// NewApd returns an Apd type
func NewApd(pubKey string, masterLogger *logging.MasterLogger) *APD {
	logger := masterLogger.PackageLogger("auto-peering-daemon")

	return &APD{
		PublicKey: pubKey,
		LocalIP:   getLocalIP(),
		PacketMap: make(map[string]string),
		DoneCh:    make(chan error),
		PacketCh:  make(chan Packet, packetLength),
		logger:    logger,
	}
}

// BroadCastPubKey broadcasts a UDP packet which contains a public key
// to the local network's broadcast address.
func (apd *APD) BroadCastPubKey(broadCastIP string, timer *time.Ticker, port int) {
	apd.logger.Infof("broadcasting on address %s:%d", defaultBroadCastIP, port)
	for range timer.C {
		apd.logger.Infof("[UDP BROADCAST] broadcasting public key")
		err := BroadCastPubKey(apd.PublicKey, broadCastIP, port)
		if err != nil {
			apd.logger.Error(err)
			apd.DoneCh <- err
			return
		}
	}
}

// Listen listens for incoming broadcasts on a local network, and reads incoming UDP broadcasts.
func (apd *APD) Listen(port int) {
	address := fmt.Sprintf(":%d", port)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		apd.logger.Error(err)
		apd.DoneCh <- err
		return
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		apd.logger.Error(err)
		apd.DoneCh <- err
		return
	}

	defer conn.Close()
	apd.logger.Infof("listening on address %s", address)

	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			apd.logger.Error(err)
			apd.DoneCh <- err
			return
		}

		message := Packet{
			PublicKey: string(buffer[:n]),
			IP:        addr.String(),
		}

		apd.PacketCh <- message
	}
}

// Run starts an auto-peering daemon process in two goroutines.
// The daemon broadcasts a public key in a goroutine, and listens
// for incoming broadcasts in another goroutine.
func (apd *APD) Run() {
	t := time.NewTicker(10 * time.Second)

	shutDownCh := make(chan os.Signal)
	signal.Notify(shutDownCh, syscall.SIGTERM, syscall.SIGINT)

	// send broadcasts at ten minute intervals
	go apd.BroadCastPubKey(defaultBroadCastIP, t, port)

	// listen for incoming broadcasts
	go apd.Listen(port)

	for {
		select {
		case <-apd.DoneCh:
			apd.logger.Fatal("Shutting down daemon")
			os.Exit(1)
		case packet := <-apd.PacketCh:
			apd.RegisterPubKey(packet)
		case <-shutDownCh:
			apd.logger.Print("Shutting down daemon")
			os.Exit(1)
		}
	}
}

// RegisterPubKey checks if a public key received from a broadcast is already registered.
// It adds only new public keys to a map.
func (apd *APD) RegisterPubKey(packet Packet) {
	if apd.PublicKey != packet.PublicKey {
		if _, ok := apd.PacketMap[packet.PublicKey]; !ok {
			apd.PacketMap[packet.PublicKey] = packet.IP
			apd.logger.Infof("Received packet %s: %s", packet.PublicKey, packet.IP)
		}
	}
}
