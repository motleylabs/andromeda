package routers

import (
	"andromeda/internal/api/controllers"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func NFTRouter(r *gin.RouterGroup) {
	nftStore := persistence.NewInMemoryStore(time.Second)
	nftController := new(controllers.NFT)

	c := r.Group("/nfts")
	c.Use()
	{
		c.GET("/activities", cache.CachePage(nftStore, time.Second*10, nftController.GetActivities))
		c.GET("/offers", cache.CachePage(nftStore, time.Second*10, nftController.GetOffers))
		c.GET("/:address", cache.CachePage(nftStore, time.Second*10, nftController.GetDetail))
	}
}
