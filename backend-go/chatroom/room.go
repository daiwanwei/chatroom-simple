package chatroom

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Room struct {
	BroadcastChan chan *Message
	joinChan      chan *Client
	leaveChan     chan *Client
	clients       map[*Client]bool
}

func NewChatRoom() *Room {
	return &Room{
		BroadcastChan: make(chan *Message),
		joinChan:      make(chan *Client),
		leaveChan:     make(chan *Client),
		clients:       make(map[*Client]bool),
	}
}

func (r *Room) Start() {
	for {
		select {
		case client := <-r.joinChan:
			r.clients[client] = true
		case client := <-r.leaveChan:
			delete(r.clients, client)
			close(client.ReceiveChan)
		case msg := <-r.BroadcastChan:
			fmt.Println("broadcast to client in the room")
			for client := range r.clients {
				select {
				case client.ReceiveChan <- msg:
					fmt.Println("send msg to client")
				default:
					delete(r.clients, client)
					close(client.ReceiveChan)
				}
			}
		}
	}
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("server:", err)
	}

	client := &Client{
		WsConn:      conn,
		ReceiveChan: make(chan *Message, 256),
		Room:        r,
		userData:    map[string]interface{}{"name": fmt.Sprintf("user%d", time.Now().Unix())},
	}
	r.joinChan <- client

	defer func() {
		r.leaveChan <- client
	}()
	go client.Send()
	client.Receive()
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}
