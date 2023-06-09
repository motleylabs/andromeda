package controllers

import (
	"andromeda/internal/api/utils"
	"andromeda/pkg/service/entrance/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NFT struct{}

// GetDetail godoc
//
// @Summary         Get NFT detail
// @Description     get detail information about the NFT
// @Tags            nfts
// @Accept          json
// @Produce         json
// @Param           address  path          string true          "NFT address"
// @Success		    200	     {object}	   types.NFT
// @Failure         500
// @Router          /nfts/{address}     [get]
func (ctrl NFT) GetDetail(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=10, stale-while-revalidate")
	address := c.Param("address")

	dataProvider := utils.GetProvider()
	nft, err := dataProvider.GetNFTDetail(address)
	if err != nil {
		log.Printf("NFT GetDetail >> DataProvder GetNFTDetail; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nft)
}

// GetOffers godoc
//
// @Summary         Get NFT offers
// @Description     get the offers with related to the NFT
// @Tags            nfts
// @Accept          json
// @Produce         json
// @Param           address  query         string true         "NFT address"
// @Success		    200	     {object}	   []types.NFTActivity
// @Failure         500
// @Router          /nfts/offers     [get]
func (ctrl NFT) GetOffers(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=10, stale-while-revalidate")
	address := c.Query("address")

	dataProvider := utils.GetProvider()
	offers, err := dataProvider.GetNFTOffers(address)
	if err != nil {
		log.Printf("NFT GetOffers >> DataProvder GetNFTOffers; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, offers)
}

// GetActivities godoc
//
// @Summary         Get NFT activities
// @Description     get the activities with related to the NFT
// @Tags            nfts
// @Accept          json
// @Produce         json
// @Param           address          query         string  true         "NFT address"
// @Param           activity_types   query         string  false        "Activity Types (['listing'])"
// @Success		    200	             {object}	   types.NFTActivityRes
// @Failure		    400
// @Failure         500
// @Router          /nfts/activities     [get]
func (ctrl NFT) GetActivities(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=10, stale-while-revalidate")
	params, err := utils.GetNFTActivityParams(c)
	if err != nil {
		log.Printf("NFT GetActivities >> Util GetNFTActivityParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dataProvider := utils.GetProvider()
	activityRes, err := dataProvider.GetNFTActivities(&params)

	if err != nil {
		log.Printf("NFT GetTimeSeries >> DataProvder GetNFTActivities; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, activityRes)
}

// GetBuyNowTx godoc
//
// @Summary         Get transaction data for buying NFT
// @Description     get the byte array for transaction to buy the NFT
// @Tags            nfts
// @Accept          json
// @Produce         json
// @Param           request          body		   types.BuyParams true     "Input data"
// @Success		    200	             {object}	   types.BuyRes
// @Failure		    400
// @Failure         500
// @Router          /nfts/buy        [post]
func (ctrl NFT) GetBuyNowTx(c *gin.Context) {
	var buyParams types.BuyParams
	if err := c.ShouldBindJSON(&buyParams); err != nil {
		log.Printf("NFT GetBuyNowTx >> ShouldBindJSON; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dataProvider := utils.GetProvider()
	buyNowTx, err := dataProvider.GetNFTBuyNowTx(&buyParams)

	if err != nil {
		log.Printf("NFT GetBuyNowTx >> DataProvder GetNFTBuyNowTx; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, buyNowTx)
}
