package routers

import (
	"andromeda/internal/api/controllers"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func StatRouter(r *gin.RouterGroup) {
	statStore := persistence.NewInMemoryStore(time.Second)
	statController := new(controllers.Stat)

	c := r.Group("/stat")
	c.Use()
	{
		c.GET("/overall", cache.CachePage(statStore, time.Minute*5, statController.GetOverallStat))
	}
}
