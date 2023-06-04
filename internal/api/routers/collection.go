package routers

import (
	"andromeda/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func CollectionRouter(r *gin.RouterGroup) {
	collectionController := new(controllers.Collection)
	wsServer := collectionController.NewWebsocketServer()

	// fmt.Println("______WEBSOCKET STARTED_____")
	go wsServer.Run()

	c := r.Group("/collections")
	c.Use()
	{
		c.GET("/trend", collectionController.GetTrends)
		c.GET("/series", collectionController.GetTimeSeries)
		c.GET("/nfts", collectionController.GetNFTs)
		c.GET("/activities", collectionController.GetActivities)
		c.GET("/:address", collectionController.GetDetail)
		c.GET("/ws", func(ctx *gin.Context) {
			collectionController.GetWs(wsServer, ctx)
		})
	}
}
