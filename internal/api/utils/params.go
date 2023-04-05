package utils

import (
	"andromeda/pkg/service/entrance/types"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTrendParams(c *gin.Context) (types.TrendParams, error) {
	var params types.TrendParams
	params.Period = c.Query("period")
	params.SortBy = c.Query("sort_by")
	params.Order = c.Query("order")
	if strings.ToLower(params.Order) != "asc" && strings.ToLower(params.Order) != "desc" {
		return params, fmt.Errorf("order param is not valid")
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return params, fmt.Errorf("limit param is not valid")
	}
	params.Limit = limit

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		return params, fmt.Errorf("offset param is not valid")
	}
	params.Offset = offset

	return params, nil
}

func GetTimeSeriesParams(c *gin.Context) (types.TimeSeriesParams, error) {
	var params types.TimeSeriesParams
	params.Address = c.Query("address")
	params.Granularity = c.Query("granularity")
	if strings.ToLower(params.Granularity) != "per_day" && strings.ToLower(params.Granularity) != "per_hour" {
		return params, fmt.Errorf("granularity param is not valid")
	}

	fromTime, err := strconv.Atoi(c.Query("from_time"))
	if err != nil {
		return params, fmt.Errorf("from_time param is not valid")
	}
	params.FromTime = fromTime

	toTime, err := strconv.Atoi(c.Query("to_time"))
	if err != nil {
		return params, fmt.Errorf("to_time param is not valid")
	}
	params.ToTime = toTime

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return params, fmt.Errorf("limit param is not valid")
	}
	params.Limit = limit

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		return params, fmt.Errorf("offset param is not valid")
	}
	params.Offset = offset

	return params, nil
}

func GetNFTParams(c *gin.Context) (types.NFTParams, error) {
	var params types.NFTParams
	params.Address = c.Query("address")
	params.SortBy = c.Query("sort_by")
	params.Order = c.Query("order")
	if strings.ToLower(params.Order) != "asc" && strings.ToLower(params.Order) != "desc" {
		return params, fmt.Errorf("order param is not valid")
	}
	params.ListingOnly = false
	if c.Query("listing_only") == "true" {
		params.ListingOnly = true
	}

	program := c.Query("program")
	if program != "" {
		params.Program = &program
	}

	auctionHouse := c.Query("auction_house")
	if auctionHouse != "" {
		params.AuctionHouse = &auctionHouse
	}

	attributes := c.Query("attributes")
	if attributes != "" {
		if err := json.Unmarshal([]byte(attributes), &params.Attributes); err != nil {
			return params, err
		}
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return params, fmt.Errorf("limit param is not valid")
	}
	params.Limit = limit

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		return params, fmt.Errorf("offset param is not valid")
	}
	params.Offset = offset

	return params, nil
}

func GetActivityParams(c *gin.Context) (types.ActivityParams, error) {
	var params types.ActivityParams
	params.Address = c.Query("address")

	activityTypes := c.Query("activity_types")
	if activityTypes != "" {
		if err := json.Unmarshal([]byte(activityTypes), &params.ActivityTypes); err != nil {
			return params, err
		}
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return params, fmt.Errorf("limit param is not valid")
	}
	params.Limit = limit

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		return params, fmt.Errorf("offset param is not valid")
	}
	params.Offset = offset

	return params, nil
}

func GetNFTActivityParams(c *gin.Context) (types.ActivityParams, error) {
	var params types.ActivityParams
	params.Address = c.Query("address")

	activityTypes := c.Query("activity_types")
	if activityTypes != "" {
		if err := json.Unmarshal([]byte(activityTypes), &params.ActivityTypes); err != nil {
			return params, err
		}
	}

	return params, nil
}
