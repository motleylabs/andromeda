package routers

import (
	"andromeda/internal/api/controllers"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func RPCRouter(r *gin.RouterGroup) {
	rpcStore := persistence.NewInMemoryStore(time.Second)
	rpcController := new(controllers.RPC)

	c := r.Group("/rpc")
	c.Use()
	{
		c.GET("/report", cache.CachePage(rpcStore, time.Minute*5, rpcController.GetReport))
	}
}
