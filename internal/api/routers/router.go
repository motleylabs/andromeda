package routers

import (
	"andromeda/internal/api/middlewares"

	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORS())
	APIRouter(r)
	WSRouter(r)

	return r
}
