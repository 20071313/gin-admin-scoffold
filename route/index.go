package route

import (
	"github.com/gin-admin-scoffold/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(r *gin.Engine) {
	RegisterWS(r)

	r.GET("hello", api.Hello)
}
