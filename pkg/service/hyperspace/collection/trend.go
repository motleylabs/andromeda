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
	projectStatParams := GetProjectStatParams(params)
	payload, err := json.Marshal(projectStatParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-project-stats", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var projectStats ProjectStatRes
	if err := json.Unmarshal(res, &projectStats); err != nil {
		return nil, err
	}

	trendRes := types.TrendRes{
		HasNextPage: projectStats.PaginationInfo.HasNextPage,
		Trends:      ConvertStatistics(projectStats.ProjectStats),
	}

	return &trendRes, nil
}

func ConvertStatistics(stats []ProjectStat) []types.Trend {
	trends := make([]types.Trend, len(stats))

	for index := range stats {
		// volume
		trends[index].Volume1D = common.GetFromIntPointer(stats[index].Volume1Day)
		trends[index].Volume7D = common.GetFromIntPointer(stats[index].Volume7Day)
		trends[index].Volume30D = common.GetFromIntPointer(stats[index].Volume1M)

		// floor price
		trends[index].Floor1D = common.GetLamportsFromPointer(stats[index].FloorPrice1Day)
		trends[index].Floor7D = common.GetLamportsFromPointer(stats[index].FloorPrice7Day)
		trends[index].Floor30D = common.GetLamportsFromPointer(stats[index].FloorPrice1M)

		// listing
		trends[index].Listed1D = common.GetFromIntPointer(stats[index].Listed1Day)

		// collection
		trends[index].Collection = types.Collection{
			ID:          stats[index].ProjectID,
			Name:        stats[index].Project.DisplayName,
			Description: stats[index].Project.Description,
			Image:       stats[index].Project.ImgURL,
		}
	}

	return trends
}

func GetProjectStatParams(input *types.TrendParams) *ProjectStatParam {
	return &ProjectStatParam{
		OrderBy:        GetOrderField(input),
		PaginationInfo: GetPaginationInfo(input),
	}
}

func GetOrderField(input *types.TrendParams) common.OrderConfig {
	orderFieldName := "floor_price"

	switch input.OrderBy {
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

	return common.OrderConfig{
		FieldName: fmt.Sprintf("%s_%s", orderFieldName, periodSuffix),
		SortOrder: input.SortBy,
	}
}

func GetPaginationInfo(input *types.TrendParams) *common.PaginationConfig {
	pageNumber := input.Offset/input.Limit + 1
	pageSize := input.Limit

	paginationInfo := common.PaginationConfig{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}

	return &paginationInfo
}
