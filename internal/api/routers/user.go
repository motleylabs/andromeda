package routers

import (
	"andromeda/internal/api/controllers"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	userStore := persistence.NewInMemoryStore(time.Second)
	userController := new(controllers.User)

	c := r.Group("/users")
	c.Use()
	{
		c.GET("/activities", cache.CachePage(userStore, time.Second*30, userController.GetActivities))
		c.GET("/nfts", cache.CachePage(userStore, time.Minute, userController.GetNFTs))
		c.GET("/offers", cache.CachePage(userStore, time.Second*30, userController.GetOffers))
	}
}
