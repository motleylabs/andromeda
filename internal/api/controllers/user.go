package controllers

import (
	"andromeda/internal/api/models"
	"andromeda/internal/api/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct{}

type LoginPayload struct {
	Address   string `json:"address" binding:"required"`
	SignedMsg string `json:"msg" binding:"required"`
}

type LoginRes struct {
	User         models.User `json:"user"`
	Token        string      `json:"token"`
	RefreshToken string      `json:"refresh_token"`
}

type TokenRes struct {
	Token string `json:"token"`
}

var userModel = new(models.User)
var nonceModel = new(models.Nonce)
var refreshTokenModel = new(models.RefreshToken)

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

// GetNonce godoc
//
// @Summary         Get nonce for message signing
// @Description     get the temporary none for message signing
// @Tags            users
// @Accept          json
// @Produce         json
// @Param           address          query         string  true         "Wallet address"
// @Success		    200	             {string}      string               "Nonce"
// @Failure		    400
// @Failure         500
// @Router          /users/nonce     [get]
func (ctrl User) GetNonce(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	nonce, err := nonceModel.FirstOrCreate(&models.Nonce{
		Address: address,
	})
	if err != nil {
		log.Printf("User GetNonce >> Nonce FirstOrCreate with address %s; %s", address, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newNonce, err := utils.GenerateRandomString(32)
	if err != nil {
		log.Printf("User GetNonce >> Util GenerateRandomString with address %s; %s", address, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	expiredAt := time.Now().Unix() + 60
	nonce.Nonce = &newNonce
	nonce.ExpiredAt = &expiredAt
	if err := nonceModel.Update(address, nonce); err != nil {
		log.Printf("User GetNonce >> Nonce Update with address %s; %s", address, err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, newNonce)
}

// Login godoc
//
// @Summary         User login
// @Description     login with user wallet address
// @Tags            users
// @Accept          json
// @Produce         json
// @Param           request          body          LoginPayload  true     "login payload"
// @Success		    200	             {object}      LoginRes
// @Failure		    400              {object}      utils.ErrorRes               "invalid payload"
// @Failure		    403              {object}      utils.ErrorRes               "invalid message signing"
// @Failure		    409              {object}      utils.ErrorRes               "nonce is expired"
// @Failure         500
// @Router          /users/login     [post]
func (ctrl User) Login(c *gin.Context) {
	var inputData LoginPayload
	if err := c.ShouldBindJSON(&inputData); err != nil {
		log.Printf("User Login >> ShouldBindJSON; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// get nonce data
	nonce, err := nonceModel.GetByAddress(inputData.Address)
	if err != nil {
		log.Printf("User Login >> Nonce GetByAddress; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// check if the nonce is valid
	if nonce.ExpiredAt == nil || nonce.Nonce == nil {
		log.Printf("User Login >> Invalid nonce with address %s", inputData.Address)
		utils.SendError(c, http.StatusConflict, "invalid nonce")
		return
	}
	if nonce.ExpiredAt != nil {
		curTime := time.Now().Unix()
		if curTime > *nonce.ExpiredAt {
			log.Printf("User Login >> Nonce is expired with address %s", inputData.Address)
			utils.SendError(c, http.StatusConflict, "nonce is expired")
			return
		}
	}

	// check message signing
	if ok := utils.ValidateMessage(inputData.Address, inputData.SignedMsg, *nonce.Nonce); !ok {
		log.Printf("User Login >> Util ValidateMessage false with address %s", inputData.Address)
		utils.SendError(c, http.StatusForbidden, "message signing is invalid")
		return
	}

	// get or create user
	user := models.User{
		Address: inputData.Address,
	}
	err = userModel.FirstOrCreate(&user)
	if err != nil {
		log.Printf("User Login >> User Create with address %s; %s", inputData.Address, err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	// get jwt token, refresh token
	token, err := utils.GenerateToken(user)
	if err != nil {
		log.Printf("User Login >> Util GenerateToken for user ID %d; %s", user.ID, err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, refreshTokenExpire, err := utils.GenerateRefreshToken(user)
	if err != nil {
		log.Printf("User Login >> Util GenerateRefreshToken for user ID %d; %s", user.ID, err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = refreshTokenModel.Create(user.ID, refreshToken, refreshTokenExpire)
	if err != nil {
		log.Printf("User Login >> Refresh Token Create for user ID %d; %s", user.ID, err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Set("user", user)

	c.JSON(http.StatusOK, LoginRes{
		User:         user,
		Token:        token,
		RefreshToken: refreshToken,
	})
}

// GetRefreshToken godoc
//
// @Summary         Get refreshed token
// @Description     get a new token using refresh token
// @Tags            users
// @Accept          json
// @Produce         json
// @Param           Authorization    header        string true                  "Bearer {token}"
// @Success		    200	             {object}      TokenRes
// @Failure		    400              {object}      utils.ErrorRes               "invalid payload"
// @Failure		    403              {object}      utils.ErrorRes               "invalid refresh token"
// @Failure		    409              {object}      utils.ErrorRes               "expired refresh token"
// @Failure         500
// @Router          /users/refresh_token     [get]
func (ctrl User) GetRefreshToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		authHeader = c.Query("token")
	}
	if authHeader == "" {
		utils.SendError(c, http.StatusBadRequest, "auth header is missing")
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		utils.SendError(c, http.StatusBadRequest, "auth header is invalid")
		return
	}

	bearerToken := parts[1]
	refreshToken, err := refreshTokenModel.GetByRefreshToken(bearerToken)
	if err != nil {
		utils.SendError(c, http.StatusForbidden, "invalid refresh token")
		return
	}

	if refreshToken.ExpiredAt.Unix() < time.Now().Unix() {
		utils.SendError(c, http.StatusConflict, "expired refresh token")
		return
	}

	token, err := utils.GenerateToken(&models.User{
		ID: refreshToken.UserID,
	})
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "token generation failed")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (ctrl User) GetMe(c *gin.Context) {

}
