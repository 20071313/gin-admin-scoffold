package api

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cast"

	"github.com/gin-admin-scoffold/utils/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
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

	client := &ws.Client{ID: uuid.NewV4().String(), Socket: conn, Msg: make(chan []byte)}

	ws.Manager.Register <- client

	go client.Read()
	go client.Write()

	go sendMsg(client)

	ws.Manager.Send([]byte("test client id: "+cast.ToString(client.ID)), client)

	//for i := 0; i < 3; i++ {
	//	time.Sleep(1 * time.Second)
	//	ws.Manager.Send([]byte(time.Now().Format("2006-01-02 15:04:05")), client)
	//}

}

func sendMsg(client *ws.Client) {
	for {
		//var err interface{}
		select {
		case _, ok := <-client.Msg:
			if !ok {
				fmt.Println("socket client closed, send data stop")
				return
			}
		default:
			fmt.Println("nothing in client.Err")
		}
		ws.Manager.Send([]byte(generateData()), client)
		time.Sleep(2 * time.Second)
	}
}

// 发送给客户端的消息样例
func generateData() (result string) {
	type msg struct {
		Id      string
		Content string
		Date    string
	}

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	resultByte, _ := json.Marshal(&msg{
		Id:      "fadsafd",
		Content: generateMd5Str(nowTime),
		Date:    nowTime,
	})
	result = string(resultByte)
	return result
}

func generateMd5Str(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
