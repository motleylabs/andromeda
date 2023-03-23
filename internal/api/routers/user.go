package routers

import (
	"andromeda/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	userController := new(controllers.User)
	c := r.Group("/users")
	c.Use()
	{
		c.GET("/nfts", userController.GetNFTs)
		c.POST("/activities", userController.GetActivities)
		c.POST("/offers", userController.GetOffers)
	}
}
