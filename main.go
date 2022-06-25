package main

import (
	"github.com/gin-admin-scoffold/route"
	"github.com/gin-admin-scoffold/utils/ws"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	route.RegisterRoute(r)
	//r.GET("/ws", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"msg": "ws msg",
	//		"":    "",
	//	})
	//})

	var websocketClient ws.ClientManager
	go websocketClient.Start()

	err := r.Run()

	if err != nil {
		return
	} // 监听并在 0.0.0.0:8080 上启动服务
}
