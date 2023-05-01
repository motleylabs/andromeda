package controllers

import (
	"andromeda/internal/api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Stat struct{}

// GetReport godoc
//
// @Summary         Get Overall Stats
// @Description     get information like total market cap, volume
// @Tags            stats
// @Accept          json
// @Produce         json
// @Success		    200	     {object}	   types.StatRes
// @Failure         500
// @Router          /stat/overall     [get]
func (ctrl Stat) GetOverallStat(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=300, stale-while-revalidate")
	dataProvider := utils.GetProvider()
	stat, err := dataProvider.GetStatOverall()
	if err != nil {
		log.Printf("Stat GetOverallStat >> DataProvider GetStatOverall; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stat)
}

// Search godoc
//
// @Summary         Search collections, users
// @Description     Search collections, users by name
// @Tags            stats
// @Accept          json
// @Produce         json
// @Param           mode     query         string true          "Search mode(collection)"
// @Param           keyword  query         string true          "Search keyword"
// @Param           limit    query         string true          "Page limit"
// @Param           offset   query         string true          "Page offset"
// @Success		    200	     {object}	   types.SearchRes
// @Failure         500
// @Router          /stat/search     [get]
func (ctrl Stat) Search(c *gin.Context) {
	params, err := utils.GetSearchParams(c)
	if err != nil {
		log.Printf("Stat Search >> Util GetSearchParams; %s", err.Error())
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	dataProvider := utils.GetProvider()
	res, err := dataProvider.GetStatSearch(&params)
	if err != nil {
		log.Printf("Stat Search >> DataProvider GetStatSearch; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
