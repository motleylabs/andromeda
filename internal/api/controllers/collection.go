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
// @Param           limit    query         int     true         "Limit"
// @Param           offset   query         int     true         "Offset"
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
// @Param           address         query         string  true         "Collection Address"
// @Param           attributes      query         string  false        "NFT attributes to filter ([{'name': 'Tattoos', 'type': 'CATEGORY', 'values': ['Barbwire']}])"
// @Param           listing_only    query         string  false        "Only listed NFTs? (true|false)"
// @Param           program         query         string  false        "Marketplace program address"
// @Param           auction_house   query         string  false        "Auction house address"
// @Param           sort_by         query         string  true         "Sort By (lowest_listing_block_timestamp)"
// @Param           order           query         string  true         "Order (ASC|DESC)"
// @Param           limit           query         int     true         "Limit"
// @Param           offset          query         int     true         "Offset"
// @Success		    200	            {object}	  types.NFTRes
// @Failure		    400
// @Failure         500
// @Router          /collections/nfts     [get]
func (ctrl Collection) GetNFTs(c *gin.Context) {
	params, err := utils.GetNFTParams(c)
	if err != nil {
		log.Printf("Collection GetNFTs >> Util GetNFTParams; %s", err.Error())
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
// @Param           address       query         string  true         "Collection address"
// @Param           from_time     query         int     true         "Start timestamp"
// @Param           to_time       query         int     true         "End timestamp"
// @Param           granularity   query         string  true         "Granularity (PER_HOUR|PER_DAY)"
// @Param           limit         query         int     true         "Limit"
// @Param           offset        query         int     true         "Offset"
// @Success		    200	          {object}	    types.TimeSeriesRes
// @Failure		    400
// @Failure         500
// @Router          /collections/series     [get]
func (ctrl Collection) GetTimeSeries(c *gin.Context) {
	params, err := utils.GetTimeSeriesParams(c)
	if err != nil {
		log.Printf("Collection GetTimeSeries >> Util GetTimeSeriesParams; %s", err.Error())
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
