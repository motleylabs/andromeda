package controllers

import (
	"andromeda/internal/api/utils"
	"andromeda/pkg/service/entrance/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Collection struct{}

// GetTrends godoc
//
// @Summary         Get collection trends
// @Description     get trending collections
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           period   query         string  true         "Period (1d|7d|1m)"
// @Param           sort_by  query         string  true         "Sort by (volume)"
// @Param           order    query         string  true         "Order (ASC|DESC)"
// @Param           limit    query         string  true         "Limit"
// @Param           offset   query         string  true         "Offset"
// @Success		    200	     {object}	   types.TrendRes
// @Failure		    400
// @Failure         500
// @Router          /collections/trend     [get]
func (ctrl Collection) GetTrends(c *gin.Context) {
	params, err := utils.GetTrendParams(c)
	if err != nil {
		log.Printf("Collection GetTrends >> Util GetTrendParams; %s", err.Error())
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

// GetNFTs godoc
//
// @Summary         Get collection NFTs
// @Description     get the list of NFTs of the collection
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           params   body          types.NFTParams true        "Search parameters"
// @Success		    200	     {object}	   types.NFTRes
// @Failure		    400
// @Failure         500
// @Router          /collections/nfts     [post]
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

// GetTimeSeries godoc
//
// @Summary         Get collection historical data
// @Description     get the historical stats for the collection
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           params   body          types.TimeSeriesParams true        "Search parameters"
// @Success		    200	     {object}	   types.TimeSeriesRes
// @Failure		    400
// @Failure         500
// @Router          /collections/series     [post]
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

// GetDetail godoc
//
// @Summary         Get collection detail
// @Description     get collection detail information with the address
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           address  path          string true                     "Collection Address"
// @Success		    200	     {object}	   types.Collection
// @Failure         500
// @Router          /collections/detail/{address} [get]
func (ctrl Collection) GetDetail(c *gin.Context) {
	address := c.Param("address")

	dataProvider := utils.GetProvider()
	collection, err := dataProvider.GetCollectionDetail(address)
	if err != nil {
		log.Printf("Collection GetDetail >> DataProvder GetDetail; %s", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, collection)
}

// GetActivities godoc
//
// @Summary         Get collection activities
// @Description     get the activities with related to the collection
// @Tags            collections
// @Accept          json
// @Produce         json
// @Param           params   body          types.ActivityParams true        "Search parameters"
// @Success		    200	     {object}	   types.ActivityRes
// @Failure		    400
// @Failure         500
// @Router          /collections/activities     [post]
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
