package controllers

import (
	"andromeda/internal/api/utils"
	"andromeda/pkg/service/entrance/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Collection struct{}

func (ctrl Collection) GetTrend(c *gin.Context) {
	var params types.TrendParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Collection GetTrend >> ShouldBindJSON; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dataProvider := utils.GetProvider()
	trends, err := dataProvider.GetCollectionTrends(&params)

	if err != nil {
		log.Printf("Collection GetTrend >> DataProvder GetCollectionTrends; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, trends)
}
