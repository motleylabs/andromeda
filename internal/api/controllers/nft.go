package controllers

import (
	"andromeda/internal/api/utils"
	"andromeda/pkg/service/entrance/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NFT struct{}

func (ctrl NFT) GetDetail(c *gin.Context) {
	address := c.Param("address")

	dataProvider := utils.GetProvider()
	collection, err := dataProvider.GetNFTDetail(address)
	if err != nil {
		log.Printf("NFT GetDetail >> DataProvder GetDetail; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, collection)
}

func (ctrl NFT) GetActivities(c *gin.Context) {
	var params types.ActivityParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("NFT GetActivities >> ShouldBindJSON; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dataProvider := utils.GetProvider()
	activityRes, err := dataProvider.GetNFTActivities(&params)

	if err != nil {
		log.Printf("NFT GetTimeSeries >> DataProvder GetNFTTimeSeries; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, activityRes)
}
