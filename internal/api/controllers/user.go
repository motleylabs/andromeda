package controllers

import (
	"andromeda/internal/api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{}

// GetNFTs godoc
//
// @Summary         Get user NFTs
// @Description     get the nfts of the wallet
// @Tags            users
// @Accept          json
// @Produce         json
// @Param           address          query         string  true         "Wallet address"
// @Success		    200	             {object}	   types.UserNFT
// @Failure         500
// @Router          /users/nfts     [get]
func (ctrl User) GetNFTs(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=60, stale-while-revalidate")
	address := c.Query("address")

	dataProvider := utils.GetProvider()
	nftRes, err := dataProvider.GetUserNFTs(address)
	if err != nil {
		log.Printf("User GetNFTs >> DataProvder GetUserNFTs; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
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
// @Param           address          query         string  true         "Wallet address"
// @Param           limit            query         int     true         "Limit"
// @Param           offset           query         int     true         "Offset"
// @Param           activity_types   query         string  true         "Activity types (['listing'])"
// @Success		    200	             {object}	   types.ActivityRes
// @Failure		    400
// @Failure         500
// @Router          /users/activities     [get]
func (ctrl User) GetActivities(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=10, stale-while-revalidate")
	params, err := utils.GetActivityParams(c, false)
	if err != nil {
		log.Printf("User GetActivities >> Util GetActivityParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dataProvider := utils.GetProvider()
	activityRes, err := dataProvider.GetUserActivities(&params)

	if err != nil {
		log.Printf("User GetActivities >> DataProvder GetUserActivities; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
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
// @Param           address          query         string  true         "Wallet address"
// @Param           limit            query         int     true         "Limit"
// @Param           offset           query         int     true         "Offset"
// @Success		    200	             {object}	   types.ActivityRes
// @Failure		    400
// @Failure         500
// @Router          /users/offers     [get]
func (ctrl User) GetOffers(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=10, stale-while-revalidate")
	params, err := utils.GetActivityParams(c, true)
	if err != nil {
		log.Printf("User GetOffers >> Util GetActivityParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dataProvider := utils.GetProvider()
	activityRes, err := dataProvider.GetUserOffers(&params)

	if err != nil {
		log.Printf("User GetOffers >> DataProvder GetUserOffers; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, activityRes)
}
