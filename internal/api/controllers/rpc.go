package controllers

import (
	"andromeda/internal/api/utils"
	"andromeda/pkg/service/web3"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RPC struct{}

type ReportRes struct {
	Volume   float64 `json:"volume"`
	TPS      uint64  `json:"tps"`
	SOLPrice float64 `json:"solPrice"`
}

// GetReport godoc
//
// @Summary         Get Report
// @Description     get information like solana tps, price
// @Tags            rpc
// @Accept          json
// @Produce         json
// @Param           address  path          string true          "auction house address"
// @Success		    200	     {object}	   ReportRes
// @Failure         500
// @Router          /rpc/report     [get]
func (ctrl RPC) GetReport(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "public, max-age=300, stale-while-revalidate")
	address := c.Param("address")

	solPrice, err := web3.GetSOLPrice()
	if err != nil {
		log.Printf("RPC GetReport >> Web3 GetSOLPrice; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	tps, err := web3.GetSOLTPS()
	if err != nil {
		log.Printf("RPC GetReport >> Web3 GetSOLPrice; %s", err.Error())
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	volume, _ := web3.GetVolume(address)

	c.JSON(http.StatusOK, ReportRes{
		Volume:   volume,
		TPS:      tps,
		SOLPrice: solPrice,
	})
}
