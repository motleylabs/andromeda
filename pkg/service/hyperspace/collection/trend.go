package collection

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

func GetTrends(params *types.TrendParams) (*types.TrendRes, error) {
	if params == nil {
		return nil, fmt.Errorf("no trend params")
	}
	projectStatParams := getProjectStatParams(params)
	payload, err := json.Marshal(projectStatParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-project-stats", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var projectStats common.ProjectStatRes
	if err := json.Unmarshal(res, &projectStats); err != nil {
		return nil, err
	}

	trendRes := types.TrendRes{
		HasNextPage: projectStats.PaginationInfo.HasNextPage,
		Trends:      convertStatistics(projectStats.ProjectStats),
	}

	return &trendRes, nil
}

func convertStatistics(stats []common.ProjectStat) []types.Trend {
	trends := make([]types.Trend, len(stats))

	for index := range stats {
		// volume
		trends[index].Volume1D = common.GetFromIntPointer(stats[index].Volume1Day)
		trends[index].Volume7D = common.GetFromIntPointer(stats[index].Volume7Day)
		trends[index].Volume30D = common.GetFromIntPointer(stats[index].Volume1M)
		trends[index].ChangeVolume1D = stats[index].Volume1DayChange

		// floor price
		trends[index].Floor1D = common.GetLamportsFromPointer(stats[index].FloorPrice1Day)
		trends[index].ChangeFloor1D = stats[index].FloorPrice1DayChange

		// listing
		trends[index].Listed1D = common.GetFromIntPointer(stats[index].Listed1Day)

		// collection
		trends[index].Collection = types.Collection{
			ID:          stats[index].ProjectID,
			Name:        stats[index].Project.DisplayName,
			Description: stats[index].Project.Description,
			Image:       stats[index].Project.ImgURL,
			Slug:        stats[index].Project.ProjectSlug,
		}
	}

	return trends
}

func getProjectStatParams(input *types.TrendParams) *common.StatParams {
	excludeProjectAttr := true
	return &common.StatParams{
		Conditions: &common.Conditions{
			ExcludeProjectAttributes: &excludeProjectAttr,
		},
		OrderBy:        getOrderField(input),
		PaginationInfo: getPaginationInfo(input),
	}
}

func getOrderField(input *types.TrendParams) *common.OrderConfig {
	orderFieldName := "floor_price"

	switch input.SortBy {
	case "volume":
		orderFieldName = "volume"
	case "listed":
		orderFieldName = "listed"
	}

	periodSuffix := "1day"
	switch input.Period {
	case "7d":
		periodSuffix = "7day"
	case "1m":
		periodSuffix = "1m"
	}

	return &common.OrderConfig{
		FieldName: fmt.Sprintf("%s_%s", orderFieldName, periodSuffix),
		SortOrder: input.Order,
	}
}

func getPaginationInfo(input *types.TrendParams) *common.PaginationConfig {
	pageNumber := input.Offset/input.Limit + 1
	pageSize := input.Limit

	paginationInfo := common.PaginationConfig{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}

	return &paginationInfo
}
