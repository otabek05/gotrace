package capture

import (
	"fmt"
	"gotracer/internal/model"
	"gotracer/internal/parser"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)


type Engine struct {
	handle *pcap.Handle
	packet chan gopacket.Packet
	parser *parser.PacketParser
	isStarted bool
}

func New() *Engine {

	return &Engine{
		packet: make(chan gopacket.Packet, 200),
		parser: parser.New(),
	}
}

func (e *Engine) Start(wsChan *chan []byte, msg *model.WSReceiveMessage) error {
	ifaces, _ := pcap.FindAllDevs()
	iface := ifaces[0]

	h, err := pcap.OpenLive(msg.NetworkInterface.Name, 65535, true, pcap.BlockForever)
	if err != nil {
		return err
	}

	 

	e.handle = h
	//e.handle.SetBPFFilter("tcp and port 80")
	e.handle.SetBPFFilter(buildBPFFilter(msg))
	go e.loop(wsChan, string(iface.Addresses[0].IP))
	e.isStarted = true

	return nil
}


func buildBPFFilter(msg *model.WSReceiveMessage) string  {
	var filters []string

	switch msg.Transport {
	case "tcp" , "udp":
		filters = append(filters, msg.Transport)
	}

	switch msg.NetworkLayer {
	case "ipv4":
		filters = append(filters, "ip")
	case "ipv6":
		filters = append(filters, "ip6")
	}

	if len(msg.ApplicationServices) > 0 {
		var ports []string
		for _, p := range msg.ApplicationServices {
			ports = append(ports, fmt.Sprintf("port %s", p))
		}
		filters = append(filters, "("+strings.Join(ports, " or ")+")")
	}

	if len(filters) == 0 {
		return ""
	}

	return strings.Join(filters, " and ")

}


func (e *Engine) IsRunning() bool {
	return e.isStarted
}