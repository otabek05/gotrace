package capture

import (
	"context"
	"encoding/hex"
	"fmt"
	"gotrace/backend/internal/types"
	"gotrace/backend/internal/ws"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)



type Scanner struct {
	mu  sync.Mutex
	handle *pcap.Handle
	cancel context.CancelFunc
	running bool
	iface string 
	snapshot int32
	promisc  bool
	filter string 
}

var DefaultScanner = &Scanner{}

func (s *Scanner) Start(iface string, snaplen int32, promisc bool, filter string) error {
	 s.mu.Lock()
    defer s.mu.Unlock()
    if s.running {
        return fmt.Errorf("already running on %s", s.iface)
    }

    handle, err := pcap.OpenLive(iface, snaplen, promisc, pcap.BlockForever)
    if err != nil {
        return err
    }
    if filter != "" {
        if err := handle.SetBPFFilter(filter); err != nil {
            handle.Close()
            return fmt.Errorf("failed to set filter: %w", err)
        }
    }

    ctx, cancel := context.WithCancel(context.Background())
    s.handle = handle
    s.cancel = cancel
    s.running = true
    s.iface = iface
    s.snapshot = snaplen
    s.promisc = promisc
    s.filter = filter

    go s.readLoop(ctx, handle)
    ws.DefaultHub.BroadcastMessage(types.WSMessage{
        Type: types.WSMessageStatus,
        Data: map[string]any{"running": true, "interface": iface},
    })
    return nil
}


func (s *Scanner) Stop() {
    s.mu.Lock()
    if !s.running {
        s.mu.Unlock()
        return
    }
    s.cancel()
    if s.handle != nil {
        s.handle.Close()
    }
    s.running = false
    iface := s.iface
    s.iface = ""
    s.mu.Unlock()

    ws.DefaultHub.BroadcastMessage(types.WSMessage{
        Type: types.WSMessageStatus,
        Data: map[string]any{"running": false, "interface": iface},
    })
}

func (s *Scanner) IsRunning() (bool, string) {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.running, s.iface
}


func (s *Scanner) readLoop(ctx context.Context, handle *pcap.Handle) {
	src := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := src.Packets()
	
	for {
		select{
		case <- ctx.Done():
			return 
		case packet, ok :=  <- packets:
			if !ok {
				return 
			}

			cp := mapPacket(packet, s.iface)
			ws.DefaultHub.BroadcastMessage(types.WSMessage{
				Type: types.WSMessagePacket,
				Data: cp,
			})

		
		}

	}
}


func mapPacket(p gopacket.Packet, iface string) types.CapturedPacket {
	var srcIP, dstIP, srcPort, dstPort, payloadStr string
	var netProto  types.NetworkProtocol = "UNKNOWN"
	var transportProto types.TransportProtocol = "UNKNOWN"

	   if ip4 := p.Layer(layers.LayerTypeIPv4); ip4 != nil {
        ip := ip4.(*layers.IPv4)
        srcIP = ip.SrcIP.String()
        dstIP = ip.DstIP.String()
        netProto = "IPv4"
    } else if ip6 := p.Layer(layers.LayerTypeIPv6); ip6 != nil {
        ip := ip6.(*layers.IPv6)
        srcIP = ip.SrcIP.String()
        dstIP = ip.DstIP.String()
        netProto = "IPv6"
    } else {
        if arp := p.Layer(layers.LayerTypeARP); arp != nil {
            a := arp.(*layers.ARP)
            srcIP = net.IP(a.SourceProtAddress).String()
            dstIP = net.IP(a.DstProtAddress).String()
            netProto = "ARP"
        }
    }

	if tcp := p.Layer(layers.LayerTypeTCP); tcp != nil {
        t := tcp.(*layers.TCP)
        srcPort = fmt.Sprintf("%d", t.SrcPort)
        dstPort = fmt.Sprintf("%d", t.DstPort)
        transportProto = "TCP"
        payloadStr = toHexString(t.Payload)
    } else if udp := p.Layer(layers.LayerTypeUDP); udp != nil {
        u := udp.(*layers.UDP)
        srcPort = fmt.Sprintf("%d", u.SrcPort)
        dstPort = fmt.Sprintf("%d", u.DstPort)
        transportProto = "UDP"
        payloadStr = toHexString(u.Payload)
    } else {
        // raw payload
        app := p.ApplicationLayer()
        if app != nil {
            payloadStr = toHexString(app.Payload())
            transportProto = "RAW"
        }
    }

    if payloadStr == "" {
        if app := p.ApplicationLayer(); app != nil {
            s := strings.TrimSpace(string(app.Payload()))
            if len(s) > 0 {
                payloadStr = s
            }
        }
    }

    length := len(p.Data())

    return types.CapturedPacket{
        Timestamp: time.Now(),
        Interface: iface,
        SrcIP:     srcIP,
        DstIP:     dstIP,
        SrcPort:   srcPort,
        DstPort:   dstPort,
        Network:   netProto,
        Transport: transportProto,
        App:       types.ApplicationProtocol("UNKNOWN"),
        Length:    length,
        Payload:   payloadStr,
    }


	
}

func toHexString(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	max := 4096
	for len(b ) > max {
		return  hex.EncodeToString(b[:max]) + "...(truncated)"
	}

	return hex.EncodeToString(b)
}