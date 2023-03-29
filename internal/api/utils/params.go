package utils

import (
	"andromeda/pkg/service/entrance/types"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTrendParams(c *gin.Context) (types.TrendParams, error) {
	var params types.TrendParams
	params.Period = c.Query("period")
	params.SortBy = c.Query("sort_by")
	params.Order = c.Query("order")
	if params.Order != "ASC" && params.Order != "DESC" {
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
