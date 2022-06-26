package ws

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

// ClientManager is a websocket manager
type ClientManager struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

// Client is a websocket client
type Client struct {
	ID     string
	Socket *websocket.Conn
	Msg    chan []byte
	Err    chan interface{}
}

// Message is an object for websocket message which is mapped to json type
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Manager define a ws server manager
var Manager = ClientManager{
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

// Start is to start a ws server
func (manager *ClientManager) Start() {
	for {
		select {
		case conn := <-manager.Register:
			manager.Clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "A new socket has connected."})
			manager.Send(jsonMessage, conn)
		case conn := <-manager.Unregister:
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Msg)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "A socket has disconnected."})
				manager.Send(jsonMessage, conn)
			}
		case message := <-manager.Broadcast:
			for conn := range manager.Clients {
				select {
				case conn.Msg <- message:
				default:
					close(conn.Msg)
					delete(manager.Clients, conn)
				}
			}
		}
	}
}

// Send is to send ws message to ws client
func (manager *ClientManager) Send(message []byte, client *Client) {
	msg := string(message)
	fmt.Println("Send msg:", msg)
	for conn := range manager.Clients {
		if conn == client {
			conn.Msg <- message
		}
	}
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		err := c.Socket.Close()
		if err != nil {
			fmt.Println("defer ws err:", err.Error())
			return
		}
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.Unregister <- c
			err := c.Socket.Close()
			if err != nil {
				fmt.Println("for ws err:", err.Error())
				return
			}
			break
		}
		fmt.Println("socket read msg: ", string(message))
		//jsonMessage, _ := json.Marshal(&Message{Sender: "Broadcast Read c.ID: " + c.ID, Content: string(message)})
		//Manager.Broadcast <- jsonMessage
	}
}

func (c *Client) Write() {
	defer func() {
		err := c.Socket.Close()
		if err != nil {
			fmt.Println("ws err:", err.Error())
			return
		}
	}()

	for {
		select {
		case message, ok := <-c.Msg:
			if !ok {
				err := c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					fmt.Println("write ws err:", err.Error())
					c.Err <- 1
					close(c.Msg)
					return
				}
				c.Err <- 0
				return
			}
			msg := string(message)
			fmt.Println("Write msg:", msg)
			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println("ws err:", err.Error())
				return
			}
		}
	}
}
