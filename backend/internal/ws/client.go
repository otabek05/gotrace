package ws

import (
	"encoding/json"
	"fmt"
	"gotracer/internal/capture"
	"gotracer/internal/model"

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
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		var wsMessage model.WSReceiveMessage
		err = json.Unmarshal(msg, &wsMessage)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch wsMessage.Type {
		case "start_capturing":
			c.engine.Stop()
			c.engine.Start(&wsMessage)
		case "Stop":
			c.engine.Stop()
		}

		fmt.Println(wsMessage)
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

func (c *Client) stop() {
	c.conn.Close()
	c.engine.Stop()
	close(c.Send)
}
