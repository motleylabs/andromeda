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
