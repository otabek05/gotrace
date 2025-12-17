package capture

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotracer/internal/model"
	"gotracer/internal/utils"
	"net"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type PacketProp struct {
	incomingTraffic bool
	outgoingTraffic bool
	iface net.IP
	targetIP *string
	bytesIn uint64
	bytesOut uint64

}

func (e *Engine) loop(msg *model.WebSocketRX) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	src := gopacket.NewPacketSource(e.handle, e.handle.LinkType())

	prop := PacketProp{
		incomingTraffic: msg.TrafficOptions == string(model.INCOMING),
		outgoingTraffic: msg.TrafficOptions == string(model.OUTGOING),
		iface: msg.NetworkInterface.Addresses[0].IP,
		targetIP: msg.IPv4,
	}

	packetChan := src.Packets()

	for {

		select {
		case <-e.ctx.Done():
			return
		case <-ticker.C:
			e.handleSpeed(&prop)
		case p, ok := <- packetChan:
			if !ok {
			   return 
			}

			e.handlePacket(p, &prop)
		}
	}
}

func (e *Engine) handlePacket(p gopacket.Packet, prop *PacketProp)  {
	var parsedLayers model.ParsedPacket
	parsedLayers.TimeStamp = time.Now().Format("2006-01-02 15:04:05.000")
	network := p.NetworkLayer()
	if network == nil {
		return
	}

	fmt.Println("Request has been arrived to handlePack")
	size := uint64(len(p.Data()))
	src := net.ParseIP(network.NetworkFlow().Src().String())
	dst := net.ParseIP(network.NetworkFlow().Dst().String())

	if src.Equal(net.IP(prop.iface)) {
		prop.bytesIn += size
		parsedLayers.Direction = model.OUTGOING
		if prop.incomingTraffic {
			return
		}

		if prop.targetIP != nil && !strings.EqualFold(dst.String(), *prop.targetIP) {
			return
		}

	} else if dst.Equal(net.IP(prop.iface)) {
		prop.bytesOut += size
		parsedLayers.Direction = model.INCOMING

		if prop.outgoingTraffic {
			return
		}

		if prop.targetIP != nil && !strings.EqualFold(src.String(), *prop.targetIP) {
			return 
		}

	}

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

		parsedLayers.App = &model.AppLayer{
			Protocol: "raw",
			Length:   len(payload),
			Data:     utils.BytesToSafeString(payload),
		}

		if utils.IsHTTPPayload(payload) {
			e.parser.ParseHTTP(payload, &parsedLayers)
		}
	}

	data := &model.WebSocketTX{
		Type:    "packets",
		Packets: &parsedLayers,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.Encode(data)

	if err := e.write(data); err != nil {
		fmt.Println(err)
		return
	}
}




func (e *Engine) handleSpeed(prop *PacketProp) {
	speed := &model.InternetSpeed{
		BytesIn:  utils.FormatBytesPerSec(prop.bytesIn),
		BytesOut: utils.FormatBytesPerSec(prop.bytesOut),
	}

	data := &model.WebSocketTX{
		Type:          "speed",
		InternetSpeed: speed,
	}

	if err := e.write(data); err != nil {
		return 
	}

	prop.bytesIn = 0
	prop.bytesOut = 0
}
