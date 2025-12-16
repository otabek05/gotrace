package capture

import (
	"fmt"
	"gotracer/internal/model"
	"gotracer/internal/parser"
	"strings"
	"sync"

	//"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/gorilla/websocket"
)

type Engine struct {
	handle *pcap.Handle
	conn   *websocket.Conn
	parser *parser.PacketParser
	mux *sync.Mutex
}

func New(conn *websocket.Conn) *Engine {

	return &Engine{
		mux: &sync.Mutex{},
		parser: parser.New(),
		conn: conn,
	}
}

func (e *Engine) Start(msg *model.WSReceiveMessage) error {
	h, err := pcap.OpenLive(msg.NetworkInterface.Name, 65535, true, pcap.BlockForever)
	if err != nil {
		return err
	}

	e.handle = h
	filters := buildBPFFilter(msg)
	if filters != "" {
		e.handle.SetBPFFilter(filters)
	}

	isOutgoingTraffic := msg.TrafficOptions == string(model.OUTGOING)
	isIncomingTraffic := msg.TrafficOptions == string(model.INCOMING)
	go e.loop(msg.NetworkInterface.Addresses[0].IP, isIncomingTraffic, isOutgoingTraffic)

	return nil
}

func buildBPFFilter(msg *model.WSReceiveMessage) string {
	var filters []string

	switch msg.Transport {
	case "tcp", "udp":
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


func (e *Engine) write(buff []byte) error {
	e.mux.Lock()
	defer e.mux.Unlock()
	return  e.conn.WriteMessage(websocket.TextMessage, buff)
}