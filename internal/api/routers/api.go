package routers

import "github.com/gin-gonic/gin"

func APIRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		CollectionRouter(api)
	}
}
