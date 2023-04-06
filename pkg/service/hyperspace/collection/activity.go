package collection

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

func GetActivities(params *types.ActivityParams) (*types.ActivityRes, error) {
	if params == nil {
		return nil, fmt.Errorf("no activity params")
	}
	activityParams := getCollectionActivityParams(params)
	payload, err := json.Marshal(activityParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-project-history", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var snapshots common.ProjectSnapshotsRes
	if err := json.Unmarshal(res, &snapshots); err != nil {
		return nil, err
	}

	activityRes := types.ActivityRes{
		Activities:  common.ConvertActivitySnapshots(snapshots.MarketPlaceSnapshots),
		HasNextPage: snapshots.PaginationInfo.HasNextPage,
	}
	return &activityRes, nil
}

func getCollectionActivityParams(input *types.ActivityParams) *common.ActivityParams {
	projectIDs := []common.ProjectIDItem{
		{
			ProjectID: input.Address,
		},
	}
	pageNumber := input.Offset/input.Limit + 1

	activityTypes := make([]string, len(input.ActivityTypes))
	for index := range input.ActivityTypes {
		activityTypes[index] = common.ConvertToActivityType(input.ActivityTypes[index])
	}

	var activityParams = common.ActivityParams{
		ActivityCondition: common.ActivityCondition{
			Projects:   &projectIDs,
			ByMPATypes: activityTypes,
		},
		PaginationInfo: &common.PaginationConfig{
			PageNumber: &pageNumber,
			PageSize:   &input.Limit,
		},
	}
	return &activityParams
}
