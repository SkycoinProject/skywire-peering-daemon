package src

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/SkycoinProject/skycoin/src/cipher"
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
}

// NewApd returns an Apd type
func NewApd() *APD {
	pubKey, _ := cipher.GenerateKeyPair()

	return &APD{
		PublicKey: pubKey.Hex(),
		LocalIP:   getLocalIP(),
		PacketMap: make(map[string]string),
		DoneCh:    make(chan error),
		PacketCh:  make(chan Packet, packetLength),
	}
}

// BroadCastPubKey broadcasts a UDP packet which contains a public key
// to the local network's broadcast address.
func (apd *APD) BroadCastPubKey(broadCastIP string, timer *time.Ticker, port int) {
	log.Printf("Auto-peering Daemon broadcasting on address %s:%d", defaultBroadCastIP, port)

	for _ = range timer.C {
		log.Println("[UDP BROADCAST] Broadcasting public key")
		err := BroadCastPubKey(apd.PublicKey, broadCastIP, port)
		if err != nil {
			log.Println(err)
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
		apd.DoneCh <- err
		return
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		apd.DoneCh <- err
		return
	}

	defer conn.Close()
	log.Printf("Auto-peering Daemon listening on address %s", address)

	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err)
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

	// send broadcasts at ten minute intervals
	go apd.BroadCastPubKey(defaultBroadCastIP, t, port)

	// listen for incoming broadcasts
	go apd.Listen(port)
}

// RegisterPubKey checks if a public key received from a broadcast is already registered.
// It adds only new public keys to a map.
func (apd *APD) RegisterPubKey(packet Packet) {
	if _, ok := apd.PacketMap[packet.PublicKey]; !ok {
		apd.PacketMap[packet.PublicKey] = packet.IP
		log.Printf("Received => %s: %s", packet.PublicKey, packet.IP)
	}
}
