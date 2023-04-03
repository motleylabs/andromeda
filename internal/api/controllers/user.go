package controllers

import (
	"andromeda/internal/api/models"
	"andromeda/internal/api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{}

type LoginPayload struct {
	Address string  `json:"address"`
	Signed  *[]byte `json:"signed,omitempty"`
}

var userModel = new(models.User)
var nonceModel = new(models.Nonce)

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
	address := c.Query("address")

	dataProvider := utils.GetProvider()
	nftRes, err := dataProvider.GetUserNFTs(address)
	if err != nil {
		log.Printf("User GetNFTs >> DataProvder GetUserNFTs; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
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
// @Param           activity_types   query         string  false        "Activity types (['LISTING'])"
// @Success		    200	             {object}	   types.ActivityRes
// @Failure		    400
// @Failure         500
// @Router          /users/activities     [get]
func (ctrl User) GetActivities(c *gin.Context) {
	params, err := utils.GetActivityParams(c)
	if err != nil {
		log.Printf("User GetActivities >> Util GetActivityParams; %s", err.Error())
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
// @Param           address          query         string  true         "Wallet address"
// @Param           limit            query         int     true         "Limit"
// @Param           offset           query         int     true         "Offset"
// @Success		    200	             {object}	   types.ActivityRes
// @Failure		    400
// @Failure         500
// @Router          /users/offers     [get]
func (ctrl User) GetOffers(c *gin.Context) {
	params, err := utils.GetActivityParams(c)
	if err != nil {
		log.Printf("User GetOffers >> Util GetActivityParams; %s", err.Error())
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

func (ctrl User) GetNonce(c *gin.Context) {
	address := c.Query("address")

	nonce, err := nonceModel.FirstOrCreate(&models.Nonce{
		Address: address,
	})
	if err != nil {
		log.Printf("User GetNonce >> Nonce FirstOrCreate with address %s; %s", address, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newNonce := "test_nonce"

	nonce.Nonce = newNonce
	if err := nonceModel.Update(nonce); err != nil {
		log.Printf("User GetNonce >> Nonce Update with address %s; %s", address, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, newNonce)
}

func (ctrl User) Login(c *gin.Context) {
	var inputData LoginPayload

	if err := c.ShouldBindJSON(&inputData); err != nil {
		log.Printf("User Login >> ShouldBindJSON; %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user, err := userModel.FirstOrCreate(&models.User{
		Address: inputData.Address,
	})

	if err != nil {
		log.Printf("User Login >> User Create with address %s; %s", inputData.Address, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}
