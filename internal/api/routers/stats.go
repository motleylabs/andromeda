package routers

import (
	"andromeda/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func StatRouter(r *gin.RouterGroup) {
	statController := new(controllers.Stat)

	c := r.Group("/stat")
	c.Use()
	{
		c.GET("/overall", statController.GetOverallStat)
		c.GET("/search", statController.Search)
	}
}
