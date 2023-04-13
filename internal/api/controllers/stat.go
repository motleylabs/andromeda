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
	dataProvider := utils.GetProvider()
	stat, err := dataProvider.GetStatOverall()
	if err != nil {
		log.Printf("Stat GetOverallStat >> DataProvider GetStatOverall; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, stat)
}
