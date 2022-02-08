package chatroom

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	WsConn      *websocket.Conn
	ReceiveChan chan *Message
	Room        *Room
	userData    map[string]interface{}
}

func (c *Client) Send() {
	for {
		var msg *Message
		if err := c.WsConn.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			c.Room.BroadcastChan <- msg
		} else {
			fmt.Println(err.Error())
			break
		}
	}
	c.WsConn.Close()
}

func (c *Client) Receive() {
	for msg := range c.ReceiveChan {
		if err := c.WsConn.WriteJSON(msg); err != nil {
			break
		}
	}
	c.WsConn.Close()
}
