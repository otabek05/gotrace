package ws

import (
	"encoding/json"
	"gotrace/backend/internal/types"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
}


type Hub struct {
	clients map[*Client]bool
	mu sync.Mutex
}


var DefaultHub = &Hub{
	clients: make(map[*Client]bool),
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {return true},
}

func (h *Hub) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
        http.Error(w, "websocket upgrade failed: "+err.Error(), http.StatusInternalServerError)
        return
    }

	client := &Client{conn: conn, send: make(chan []byte, 256)}
	h.register(client)
	go func ()  {
	  ticker := time.NewTicker(30 *time.Second)
	  defer func() {
		ticker.Stop()
		client.conn.Close()
	  }()

	  for {
		select {
		case msg, ok := <- client.send:
			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return 
			}

			w.Write(msg)
			if err := w.Close(); err != nil {
				return 
			}

		case <- ticker.C:
			client.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10 *time.Second))
		}
	  }

	}()

	
	go func (){
		defer h.unregister(client)
		client.conn.SetReadLimit(512)
		client.conn.SetReadDeadline(time.Now().Add(120 *time.Second))
		client.conn.SetPongHandler(func(appData string) error {
			client.conn.SetReadDeadline(time.Now().Add(120 *time.Second))
			return nil
		})

		for {
			if _, _, err := client.conn.NextReader(); err != nil {
				break
			}
		}
	}()
}


func (h *Hub) register(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.clients[c] = true
}

func (h *Hub) unregister(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.clients[c]; ok {
		delete(h.clients, c)
		close(c.send)
		c.conn.Close()
	}

}

func (h *Hub) BroadcastMessage(msg types.WSMessage) {
	b, err := json.Marshal(msg)
	if err != nil {
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	for c := range h.clients {
		select{
		case c.send <- b:
		default:
			delete(h.clients,c)
			close(c.send)
			c.conn.Close()
		}
	}
}

