package routers

import (
	"andromeda/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func CollectionRouter(r *gin.RouterGroup) {
	collectionController := new(controllers.Collection)
	c := r.Group("/collections")
	c.Use()
	{
		c.GET("/detail/:address", collectionController.GetDetail)
		c.GET("/trend", collectionController.GetTrends)
		c.POST("/nfts", collectionController.GetNFTs)
		c.POST("/series", collectionController.GetTimeSeries)
		c.POST("/activities", collectionController.GetActivities)
	}
}
