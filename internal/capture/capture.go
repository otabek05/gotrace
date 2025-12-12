package capture

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotrace/internal/model"
	"gotrace/internal/parser"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type Engine struct {
	handle *pcap.Handle
	packet chan gopacket.Packet
	parser *parser.PacketParser
}

func New() *Engine {

	return &Engine{
		packet: make(chan gopacket.Packet, 200),
		parser: parser.New(),
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
	e.handle.SetBPFFilter("tcp and port 80")
	go e.loop(wsChan)

	return nil
}

func (e *Engine) loop(wsChan *chan []byte) {
	src := gopacket.NewPacketSource(e.handle, e.handle.LinkType())
	for p := range src.Packets() {
		var parsedLayers model.ParsedPacket

		fmt.Println(p)
		if ethLayer := p.Layer(layers.LayerTypeEthernet); ethLayer != nil {
			eth := ethLayer.(*layers.Ethernet)
			e.parser.ParseEthernet(eth, &parsedLayers)
		}

		if ipLayer := p.Layer(layers.LayerTypeIPv4); ipLayer != nil {
			ip := ipLayer.(*layers.IPv4)
			e.parser.ParseIPv4(ip, &parsedLayers)
		}

		if tcpLayer := p.Layer(layers.LayerTypeTCP); tcpLayer != nil {
			tcp := tcpLayer.(*layers.TCP)
			e.parser.ParseTCP(tcp, &parsedLayers)
		}

		if udpLayer := p.Layer(layers.LayerTypeUDP); udpLayer != nil {
			udp := udpLayer.(*layers.UDP)
			e.parser.ParseUDP(udp, &parsedLayers)
		}

		if dnsLayer := p.Layer(layers.LayerTypeDNS); dnsLayer != nil {
			dns := dnsLayer.(*layers.DNS)
			e.parser.ParseDNS(dns, &parsedLayers)
		}

		if app := p.ApplicationLayer(); app != nil {
			payload := app.Payload()
			if isHTTPPayload(payload) {
				e.parser.ParseHTTP(payload, &parsedLayers)
			}
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		enc.Encode(parsedLayers)
		
		*wsChan <- buf.Bytes()

	}
}

func (e *Engine) Stop() {
	if e.handle != nil {
		e.handle.Close()
	}
}

func isHTTPPayload(b []byte) bool {
	s := string(b)
	return strings.HasPrefix(s, "GET ") ||
		strings.HasPrefix(s, "POST ") ||
		strings.HasPrefix(s, "PUT ") ||
		strings.HasPrefix(s, "DELETE ") ||
		strings.HasPrefix(s, "HEAD ") ||
		strings.HasPrefix(s, "OPTIONS ") ||
		strings.HasPrefix(s, "HTTP/")
}
