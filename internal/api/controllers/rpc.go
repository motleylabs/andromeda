package controllers

import (
	"andromeda/internal/api/state"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RPC struct{}

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
	c.Writer.Header().Set("Cache-Control", "public, max-age=10, stale-while-revalidate")
	// _ := c.Param("address")

	c.JSON(http.StatusOK, state.GetReport())
}
