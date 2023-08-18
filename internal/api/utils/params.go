package utils

import (
	"andromeda/pkg/service/entrance/types"
	"encoding/base64"
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

	name := c.Query("name")
	if name != "" {
		rawDecodedText, err := base64.StdEncoding.DecodeString(name)
		if err != nil {
			return params, fmt.Errorf("name param is not valid")
		}
		nameStr := string(rawDecodedText)
		params.Name = &nameStr
	}

	minStr := c.Query("min")
	if minStr != "" {
		min, err := strconv.ParseFloat(minStr, 64)
		if err != nil {
			return params, fmt.Errorf("min param is not valid")
		}
		params.PriceMin = &min
	}

	maxStr := c.Query("max")
	if maxStr != "" {
		max, err := strconv.ParseFloat(maxStr, 64)
		if err != nil {
			return params, fmt.Errorf("max param is not valid")
		}
		params.PriceMax = &max
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

func GetActivityParams(c *gin.Context, emptyActivity bool) (types.ActivityParams, error) {
	var params types.ActivityParams
	params.Address = c.Query("address")

	activityTypes := c.Query("activity_types")
	if activityTypes == "" {
		if !emptyActivity {
			return params, fmt.Errorf("activity types are missing")
		}
	} else {
		if err := json.Unmarshal([]byte(activityTypes), &params.ActivityTypes); err != nil {
			return params, err
		}

		for _, activityType := range params.ActivityTypes {
			isValid := false
			for _, stringType := range types.ActivityStringTypes {
				if strings.ToLower(activityType) == stringType {
					isValid = true
				}
			}
			if !isValid {
				return params, fmt.Errorf("invalid activity type: %s", activityType)
			}
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

func GetSearchParams(c *gin.Context) (types.SearchParams, error) {
	var params types.SearchParams
	params.Mode = c.Query("mode")
	params.Keyword = c.Query("keyword")

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		return params, fmt.Errorf("offset param is not valid")
	}
	params.Offset = offset

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return params, fmt.Errorf("limit param is not valid")
	}
	params.Limit = limit

	return params, nil
}

func GetWebsocketParams(c *gin.Context) (types.WebsocketParams, error) {
	var params types.WebsocketParams
	params.CollectionID = c.Query("collection_id")

	return params, nil
}
