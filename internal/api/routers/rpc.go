package routers

import (
	"andromeda/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func RPCRouter(r *gin.RouterGroup) {
	rpcController := new(controllers.RPC)

	c := r.Group("/rpc")
	c.Use()
	{
		c.GET("/report", rpcController.GetReport)
	}
}
