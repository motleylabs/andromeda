package controllers

import (
	"andromeda/internal/api/utils"
	"andromeda/pkg/service/entrance/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Collection struct{}

func (ctrl Collection) GetTrends(c *gin.Context) {
	var params types.TrendParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Collection GetTrends >> ShouldBindJSON; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dataProvider := utils.GetProvider()
	trends, err := dataProvider.GetCollectionTrends(&params)

	if err != nil {
		log.Printf("Collection GetTrends >> DataProvder GetCollectionTrends; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, trends)
}

func (ctrl Collection) GetNFTs(c *gin.Context) {
	var params types.NFTParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Collection GetNFTs >> ShouldBindJSON; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dataProvider := utils.GetProvider()
	nfts, err := dataProvider.GetCollectionNFTs(&params)

	if err != nil {
		log.Printf("Collection GetNFTs >> DataProvder GetCollectionNFTs; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, nfts)
}

func (ctrl Collection) GetTimeSeries(c *gin.Context) {
	var params types.TimeSeriesParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Collection GetTimeSeries >> ShouldBindJSON; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dataProvider := utils.GetProvider()
	series, err := dataProvider.GetCollectionTimeSeries(&params)

	if err != nil {
		log.Printf("Collection GetTimeSeries >> DataProvder GetCollectionTimeSeries; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, series)
}

func (ctrl Collection) GetDetail(c *gin.Context) {
	address := c.Param("address")

	dataProvider := utils.GetProvider()
	collection, err := dataProvider.GetCollectionDetail(address)
	if err != nil {
		log.Printf("Collection GetDetail >> DataProvder GetDetail; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, collection)
}

func (ctrl Collection) GetActivities(c *gin.Context) {
	var params types.ActivityParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("Collection GetActivities >> ShouldBindJSON; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dataProvider := utils.GetProvider()
	activityRes, err := dataProvider.GetCollectionActivities(&params)

	if err != nil {
		log.Printf("Collection GetTimeSeries >> DataProvder GetCollectionTimeSeries; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, activityRes)
}
