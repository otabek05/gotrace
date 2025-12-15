package capture

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotracer/internal/model"
	"net"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)


func (e *Engine) loop(ifaceAddr net.IP, incomingDirection, outgoingDirection bool) {
	src := gopacket.NewPacketSource(e.handle, e.handle.LinkType())
	var bytesIn, bytesOut uint64

	ticker := time.NewTicker(time.Second)

	go func ()  {
	  for range ticker.C {
		speed := &model.InternetSpeed{
			BytesIn: formatBytesPerSec(bytesIn),
			BytesOut: formatBytesPerSec(bytesOut),
		}

		data := &model.WSSendingMessage{
			Type: "speed",
			InternetSpeed: speed,
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		enc.Encode(data)
		err := e.write(buf.Bytes())
		if (err != nil) {
			return 
		}
		bytesIn = 0
		bytesOut = 0
	  }	
	}()


	for p := range src.Packets() {
		var parsedLayers model.ParsedPacket
		parsedLayers.TimeStamp = time.Now().Format("2006-01-02 15:04:05.000")

		network := p.NetworkLayer()
		if network == nil {
			continue
		}

		size := uint64(len(p.Data()))
		src := net.ParseIP(network.NetworkFlow().Src().String())
		dst := net.ParseIP(network.NetworkFlow().Dst().String())

		if src.Equal(net.IP(ifaceAddr)) {
			bytesOut += size
			parsedLayers.Direction = model.OUTGOING
			if incomingDirection {
				continue
			}

		} else if dst.Equal(net.IP(ifaceAddr)) {
			bytesIn += size
			parsedLayers.Direction = model.INCOMING

			if outgoingDirection {
				continue
			}
		}


		//fmt.Println(p)
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

		data := &model.WSSendingMessage{
			Type: "packets",
			Packets: &parsedLayers,
		}

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		enc.Encode(data)
		
		err := e.write( buf.Bytes())
		if err != nil {
			break
		}

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


func formatBytesPerSec(b uint64) string {
    const (
        KB = 1024
        MB = 1024 * KB
        GB = 1024 * MB
    )

    switch {
    case b >= GB:
        return fmt.Sprintf("%.2f GB/s", float64(b)/GB)
    case b >= MB:
        return fmt.Sprintf("%.2f MB/s", float64(b)/MB)
    case b >= KB:
        return fmt.Sprintf("%.2f KB/s", float64(b)/KB)
    default:
        return fmt.Sprintf("%d B/s", b)
    }
}
