package src

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	defaultBroadCastIP = "255.255.255.255"
)

type Packet struct {
	PublicKey string
	IP        string
}

type Apd struct {
	PublicKey string
	LocalIP   string
	Port      int
	PacketMap map[string]string
	DoneCh    chan error
	PacketCh  chan Packet
}

// NewApd returns an Apd type
func NewApd() *Apd {
	return &Apd{
		Port:      3000,
		LocalIP:   getLocalIP(),
		PacketMap: make(map[string]string),
		DoneCh:    make(chan error),
		PacketCh:  make(chan Packet, 10),
	}
}

// BroadCastPubKey broadcasts a UDP packet which contains a public key
// to the local network's broadcast address.
func (apd *Apd) BroadCastPubKey(broadCastIP string, timer *time.Ticker) {
	for {
		select {
		case <-timer.C:
			log.Println("Broadcasting public key...")
			err := BroadCastPubKey(apd.PublicKey, broadCastIP, apd.Port)
			if err != nil {
				log.Println(err)
				apd.DoneCh <- err
				return
			}
		}
	}
}

// Listen listens for incoming broadcasts on a local network, and reads incoming UDP broadcasts.
func (apd *Apd) Listen() {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", apd.Port))
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		apd.DoneCh <- err
		return
	}

	defer conn.Close()
	log.Printf("Auto-peering Daemon up and running on port %d", apd.Port)

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
func (apd *Apd) Run() {
	t := time.NewTicker(10 * time.Second)

	// send broadcasts at ten minute intervals
	go apd.BroadCastPubKey(defaultBroadCastIP, t)

	// listen for incoming broadcasts
	go apd.Listen()
}

// RegisterPubKey checks if a public key received from a broadcast is already registered.
// It adds only new public keys to a map.
func (apd *Apd) RegisterPubKey(packet Packet) {
	if _, ok := apd.PacketMap[packet.PublicKey]; !ok {
		apd.PacketMap[packet.PublicKey] = packet.IP
		log.Printf("Received => %s: %s", packet.PublicKey, packet.IP)
	}
}
