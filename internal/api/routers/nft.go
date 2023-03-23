package routers

import (
	"andromeda/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func NFTRouter(r *gin.RouterGroup) {
	nftController := new(controllers.NFT)
	c := r.Group("/nfts")
	c.Use()
	{
		c.GET("/detail/:address", nftController.GetDetail)
		c.GET("/offers", nftController.GetOffers)
		c.POST("/activities", nftController.GetActivities)
	}
}
