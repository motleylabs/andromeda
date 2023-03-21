package collection

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

func GetTimeSeries(params *types.TimeSeriesParams) (*types.TimeSeriesRes, error) {
	if params == nil {
		return nil, fmt.Errorf("no trend params")
	}
	projectStatParams := GetProjectHistParams(params)
	payload, err := json.Marshal(projectStatParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-project-stat-hist", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var historyRes ProjectStatHistRes
	if err := json.Unmarshal(res, &historyRes); err != nil {
		return nil, err
	}

	seriesRes := types.TimeSeriesRes{
		Series:      ConvertHistEntries(historyRes.HistEntries),
		HasNextPage: historyRes.PaginationInfo.HasNextPage,
	}
	return &seriesRes, nil
}

func ConvertHistEntries(entries []ProjectStatHistEntry) []types.TimeSeries {
	series := make([]types.TimeSeries, len(entries))

	for index := range entries {
		listed := 0
		if entries[index].NumOfTokenListed != nil {
			listed = *entries[index].NumOfTokenListed
		}
		series[index] = types.TimeSeries{
			Timestamp:  fmt.Sprintf("%d", entries[index].Timestamp),
			FloorPrice: common.GetLamportsFromPointer(entries[index].FloorPrice),
			Listed:     listed,
		}
	}

	return series
}

func GetProjectHistParams(input *types.TimeSeriesParams) *common.StatParams {
	projectIDs := []string{input.Address}
	pageNumber := input.Offset/input.Limit + 1

	var statParams common.StatParams
	statParams.Conditions = &common.Conditions{
		ProjectIDs:      &projectIDs,
		StartTimestamp:  &input.FromTime,
		EndTimestamp:    &input.ToTime,
		TimeGranularity: &input.Granularity,
	}
	statParams.PaginationInfo = &common.PaginationConfig{
		PageNumber: &pageNumber,
		PageSize:   &input.Limit,
	}
	return &statParams
}
