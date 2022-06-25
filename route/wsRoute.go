package route

import (
	"github.com/gin-admin-scoffold/api"
	"github.com/gin-gonic/gin"
)

func RegisterWS(r *gin.Engine) {
	websocket := r.Group("websocket")
	{
		websocket.GET("wsPage", api.WsPage)
	}
}
