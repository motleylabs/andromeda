package utils

import "github.com/gin-gonic/gin"

type ErrorRes struct {
	Error string `json:"error"`
}

func SendError(c *gin.Context, code int, errorMsg string) {
	c.AbortWithStatusJSON(code, ErrorRes{
		Error: errorMsg,
	})
}
