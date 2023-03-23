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
		log.Printf("NFT GetDetail >> DataProvder GetNFTDetail; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, collection)
}

func (ctrl NFT) GetOffers(c *gin.Context) {
	address := c.Query("address")

	dataProvider := utils.GetProvider()
	offers, err := dataProvider.GetNFTOffers(address)
	if err != nil {
		log.Printf("NFT GetOffers >> DataProvder GetNFTOffers; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, offers)
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
		log.Printf("NFT GetTimeSeries >> DataProvder GetNFTActivities; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, activityRes)
}
