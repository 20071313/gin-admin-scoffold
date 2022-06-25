package api

import (
	"github.com/gin-admin-scoffold/utils/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
	"net/http"
)

// WsPage is a websocket handler
func WsPage(c *gin.Context) {
	// change the request to websocket model
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//for {
	//err = conn.WriteMessage(1, []byte("Hi this is from server!"+time.Now().Format("2006-01-02 15:04:05")))
	//if err != nil {
	//	log.Println(err)
	//}
	//time.Sleep(1 * time.Second)
	//}

	// websocket connect

	client := &ws.Client{ID: uuid.NewV4().String(), Socket: conn, Msg: make(chan []byte)}

	ws.Manager.Register <- client

	ws.Manager.Send([]byte("test client id: "+cast.ToString(client.ID)), nil)

	go client.Read()
	go client.Write()
}
