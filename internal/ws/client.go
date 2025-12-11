package ws

import (
	"fmt"
	"gotrace/internal/capture"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	engine *capture.Engine
	IP     string
	Send   chan []byte
}

func (c *Client) read() {
	defer func() {
		DefaultHub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		message := string(msg)

		switch message {
		case "Start":
			c.engine.Start(&c.Send)
		case "Stop":
			c.engine.Stop()
		}

		fmt.Println(string(msg))
	}
}

func (c *Client) write() {
	for msg := range c.Send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}
