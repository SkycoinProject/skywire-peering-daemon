package src

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Packet struct {
	PublicKey string
	IP        string
}

type Apd struct {
	PublicKey   string
	BroadCastIP string
	LocalIP     string
	Port        int
	Host        string
	PacketMap   map[string]string
	DoneCh      chan error
	PacketCh    chan Packet
}

func NewApd() *Apd {
	return &Apd{
		BroadCastIP: "255.255.255.255",
		Port:        3000,
		LocalIP:     getLocalIP(),
		Host:        "localhost",
		PacketMap:   make(map[string]string),
		DoneCh:      make(chan error),
		PacketCh:    make(chan Packet, 10),
	}
}

func (apd *Apd) BroadCastPubKey(timer *time.Ticker) {
	for {
		select {
		case <-timer.C:
			log.Println("Broadcasting public key...")
			err := BroadCastPubKey(apd.PublicKey, apd.BroadCastIP, apd.Port)
			if err != nil {
				log.Println(err)
				apd.DoneCh <- err
				return
			}
		}
	}
}

func (apd *Apd) Listen() {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", apd.Port))
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		apd.DoneCh <- err
		return
	}

	defer conn.Close()
	log.Printf("Auto-peering Daemon up and listening on port %d", apd.Port)

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

func (apd *Apd) Run() {
	t := time.NewTicker(10 * time.Second)

	go apd.BroadCastPubKey(t)
	go apd.Listen()
}

func (apd *Apd) RegisterPubKey(packet Packet) {
	if _, ok := apd.PacketMap[packet.PublicKey]; !ok {
		apd.PacketMap[packet.PublicKey] = packet.IP
		log.Printf("Received => %s: %s", packet.PublicKey, packet.IP)
	}
}
