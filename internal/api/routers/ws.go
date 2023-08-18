package routers

import (
	"andromeda/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

// WSRouter function will routes websocket connections
func WSRouter(r *gin.Engine) {
	wsController := new(controllers.WS)
	wsServer := wsController.InitWS()

	r.GET("/ws", func(ctx *gin.Context) {
		wsController.GetWS(wsServer, ctx)
	})
}
