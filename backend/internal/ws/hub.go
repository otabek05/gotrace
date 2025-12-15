package ws

import (
	"fmt"
	"gotracer/internal/capture"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)


type Hub struct {
    clients []Client
	register chan *Client
	unregister chan *Client
}

var DefaultHub = New()

func New() *Hub {
	return &Hub{
		clients: make([]Client, 10),
		register:  make(chan *Client),
		unregister: make(chan *Client),
	}
}


func (h *Hub) Run() {
	for {
		select{
		case client := <- h.register:
			h.clients = append(h.clients, *client)
			fmt.Println("client connected: ", client.IP)
		case client := <- h.unregister:
			client.stop()
			h.removeClient(client)
			fmt.Println("client disconnected: ", client.IP)

		}
	}
}


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {return true},
}


func (h *Hub) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	engine := capture.New(conn)

	client := &Client{
		conn: conn,
		IP:  ip,
		engine: engine,
		Send: make(chan []byte),
	}

	h.register <- client

	go client.read()
	
}

func (h *Hub) removeClient( client *Client) {
	for i, c := range h.clients {
		if c == *client {
			h.clients = append(h.clients[:i], h.clients[i+1:]... )
		}
	}
}