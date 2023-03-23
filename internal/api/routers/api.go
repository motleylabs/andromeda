package routers

import (
	"andromeda/docs"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

func APIRouter(r *gin.Engine) {

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Andromeda"
	docs.SwaggerInfo.Description = "This is a data layer for the nightmarket server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	api := r.Group("/api")
	{
		CollectionRouter(api)
		NFTRouter(api)
		UserRouter(api)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
