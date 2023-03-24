package controllers

import (
	"andromeda/internal/api/utils"
	"andromeda/pkg/service/entrance/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (ctrl User) GetNFTs(c *gin.Context) {
	address := c.Query("address")

	dataProvider := utils.GetProvider()
	nftRes, err := dataProvider.GetUserNFTs(address)
	if err != nil {
		log.Printf("User GetNFTs >> DataProvder GetUserNFTs; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, nftRes)
}

// GetActivities godoc
//
// @Summary         Get user activities
// @Description     get the activities with related to the wallet
// @Tags            users
// @Accept          json
// @Produce         json
// @Param           params   body          types.ActivityParams true        "Search parameters"
// @Success		    200	     {object}	   types.ActivityRes
// @Failure		    400
// @Failure         500
// @Router          /users/activities     [post]
func (ctrl User) GetActivities(c *gin.Context) {
	var params types.ActivityParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("User GetActivities >> ShouldBindJSON; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dataProvider := utils.GetProvider()
	activityRes, err := dataProvider.GetUserActivities(&params)

	if err != nil {
		log.Printf("User GetActivities >> DataProvder GetUserActivities; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, activityRes)
}

// Getoffers godoc
//
// @Summary         Get user offers
// @Description     get the offers with related to the wallet
// @Tags            users
// @Accept          json
// @Produce         json
// @Param           params   body          types.ActivityParams true        "Search parameters"
// @Success		    200	     {object}	   types.ActivityRes
// @Failure		    400
// @Failure         500
// @Router          /users/offers     [post]
func (ctrl User) GetOffers(c *gin.Context) {
	var params types.ActivityParams
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Printf("User GetOffers >> ShouldBindJSON; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	dataProvider := utils.GetProvider()
	activityRes, err := dataProvider.GetUserOffers(&params)

	if err != nil {
		log.Printf("User GetOffers >> DataProvder GetUserOffers; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, activityRes)
}
