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

	go ws.Manager.Start()

	// 监听并在 0.0.0.0:8080 上启动服务
	err := r.Run("127.0.0.1:8080")

	if err != nil {
		return
	}

}
