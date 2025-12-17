package capture

import (
	"bytes"
	"context"
	"encoding/json"
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

	ctx    context.Context
	cancel context.CancelFunc
}

func New(conn *websocket.Conn) *Engine {
	return &Engine{
		mux: &sync.Mutex{},
		parser: parser.New(),
		conn: conn,
	}
}

func (e *Engine) Start(msg *model.WebSocketRX) error {
	e.Stop()

	h, err := pcap.OpenLive(msg.NetworkInterface.Name, 65535, true, pcap.BlockForever)
	if err != nil {
		return err
	}

	e.handle = h
	filters := buildBPFFilter(msg)
	if filters != "" {
		e.handle.SetBPFFilter(filters)
	}

	ctx, cancel := context.WithCancel(context.Background())
	e.ctx = ctx
	e.cancel = cancel

	go e.loop(msg)

	return nil
}

func buildBPFFilter(msg *model.WebSocketRX) string {
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


func (e *Engine) Stop() {
	if e.cancel != nil {
		e.cancel()
	}

	if e.handle != nil {
		e.handle.Close()
	}
}


func (e *Engine) write(data any) error {
	e.mux.Lock()
	defer e.mux.Unlock()

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.Encode(data)
	
	return  e.conn.WriteMessage(websocket.TextMessage, buf.Bytes())
}