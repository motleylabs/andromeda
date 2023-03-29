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
		c.GET("/series", collectionController.GetTimeSeries)
		c.GET("/nfts", collectionController.GetNFTs)
		c.POST("/activities", collectionController.GetActivities)
	}
}
