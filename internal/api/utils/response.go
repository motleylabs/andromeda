package utils

import "github.com/gin-gonic/gin"

func SendError(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, gin.H{
		"error": msg,
	})
}
