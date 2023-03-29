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
		c.GET("/activities", userController.GetActivities)
		c.GET("/nfts", userController.GetNFTs)
		c.GET("/offers", userController.GetOffers)
	}
}
