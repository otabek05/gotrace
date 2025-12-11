package capture

import (
	"encoding/json"
	"fmt"
	"gotrace/internal/model"
	"gotrace/internal/parser/ethernet"
	"gotrace/internal/parser/ipv4"
	"gotrace/internal/parser/tcp"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Engine struct {
	handle    *pcap.Handle
	packet    chan gopacket.Packet
	ethParser *ethernet.EthernetParser
	ipParser *ipv4.IPV4Parser
	tcpParser *tcp.TCPParser
}

func New() *Engine {

	return &Engine{
		packet:    make(chan gopacket.Packet, 200),
		tcpParser: tcp.New(),
		ethParser: ethernet.New(),
		ipParser:  ipv4.Net(),
	}
}

func (e *Engine) Start(wsChan *chan []byte) error {
	ifaces, _ := pcap.FindAllDevs()
	iface := ifaces[0]

	h, err := pcap.OpenLive(iface.Name, 65535, true, pcap.BlockForever)
	if err != nil {
		return err
	}

	e.handle = h
	go e.loop(wsChan)

	return nil
}

func (e *Engine) loop(wsChan *chan []byte) {
	src := gopacket.NewPacketSource(e.handle, e.handle.LinkType())
	for p := range src.Packets() {
		if ethLayer := p.Layer(layers.LayerTypeEthernet); ethLayer != nil {
			eth := ethLayer.(*layers.Ethernet)
			send(wsChan, e.ethParser.Parse(eth))
		}

		if ipLayer := p.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			ip := ipLayer.(*layers.IPv4)
			send(wsChan, e.ipParser.Parse(ip))
		}

	}
}

func (e *Engine) Stop() {
	if e.handle != nil {
		e.handle.Close()
	}
}

func send(wsChan *chan []byte, data *model.LayerInfo) {
	jsonData, _ := json.Marshal(data)
	fmt.Println(data)
	*wsChan <- jsonData
}
