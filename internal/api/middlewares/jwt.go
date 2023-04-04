package middlewares

import (
	"andromeda/internal/api/models"
	"andromeda/internal/api/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var userModel = new(models.User)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := utils.GetClaimsFromJWT(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if claims["exp"] == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "JWT has no expiration",
			})
			return
		}

		if _, ok := claims["exp"].(float64); !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid JWT expiration",
			})
			return
		}

		if int64(claims["exp"].(float64)) < time.Now().Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "JWT expired",
			})
			return
		}

		userID := uint(claims["ID"].(float64))
		user, err := userModel.GetByID(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
				"error": "user not found",
			})
		}

		c.Set("user", user)

		c.Next()
	}
}
