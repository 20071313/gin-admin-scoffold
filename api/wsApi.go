package api

import (
	"github.com/gin-admin-scoffold/utils/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"net/http"
)

// WsPage is a websocket handler
func WsPage(c *gin.Context) {
	// change the request to websocket model
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	// websocket connect
	client := &ws.Client{ID: uuid.NewV4().String(), Socket: conn, Send: make(chan []byte)}

	ws.Manager.Register <- client

	go client.Read()
	go client.Write()
}
