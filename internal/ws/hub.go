package ws

import (
	"fmt"
	"gotrace/internal/capture"
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
			close(client.Send)
			print("client disconnected: ", client.IP)

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

	engine := capture.New()

	client := &Client{
		conn: conn,
		IP:  ip,
		engine: engine,
		Send: make(chan []byte),
	}

	h.register <- client

	go client.read()
	go client.write()
	
}

