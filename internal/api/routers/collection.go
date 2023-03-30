package routers

import (
	"andromeda/internal/api/controllers"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"

	"github.com/gin-gonic/gin"
)

func CollectionRouter(r *gin.RouterGroup) {
	collectionStore := persistence.NewInMemoryStore(time.Second)
	collectionController := new(controllers.Collection)

	c := r.Group("/collections")
	c.Use()
	{
		c.GET("/trend", cache.CachePage(collectionStore, time.Minute, collectionController.GetTrends))
		c.GET("/series", cache.CachePage(collectionStore, time.Minute*30, collectionController.GetTimeSeries))
		c.GET("/nfts", cache.CachePage(collectionStore, time.Minute, collectionController.GetNFTs))
		c.GET("/activities", cache.CachePage(collectionStore, time.Minute, collectionController.GetActivities))
		c.GET("/:address", cache.CachePage(collectionStore, time.Hour, collectionController.GetDetail))
	}
}
