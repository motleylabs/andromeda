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
	activityParams := getActivityParams(params)
	payload, err := json.Marshal(activityParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-project-history", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var snapshots ProjectSnapshotsRes
	if err := json.Unmarshal(res, &snapshots); err != nil {
		return nil, err
	}

	activityRes := types.ActivityRes{
		Activities:  convertSnapshots(snapshots.MarketPlaceSnapshots),
		HasNextPage: snapshots.PaginationInfo.HasNextPage,
	}
	return &activityRes, nil
}

func convertSnapshots(snapshots []MarketPlaceSnapshot) []types.Activity {
	activities := make([]types.Activity, len(snapshots))

	for index := range snapshots {
		activities[index].Name = snapshots[index].Name
		activities[index].Image = snapshots[index].MetadataImg
		activities[index].Mint = snapshots[index].TokenAddress
		activities[index].Owner = snapshots[index].Owner

		if snapshots[index].MarketPlaceState != nil {
			price := common.GetLamportsFromPointer(snapshots[index].MarketPlaceState.Price)
			activities[index].MarketPlaceProgramAddress = snapshots[index].MarketPlaceState.MarketPlaceProgramID
			activities[index].Signature = snapshots[index].MarketPlaceState.Signature
			activities[index].CreatedAt = snapshots[index].MarketPlaceState.CreatedAt
			activities[index].ActivityType = snapshots[index].MarketPlaceState.Type
			activities[index].Price = &price
		}

	}
	return activities
}

func getActivityParams(input *types.ActivityParams) *common.ActivityParams {
	projectIDs := []common.ProjectIDItem{
		{
			ProjectID: input.Address,
		},
	}
	pageNumber := input.Offset/input.Limit + 1

	var activityParams = common.ActivityParams{
		ActivityCondition: common.ActivityCondition{
			Projects:   projectIDs,
			ByMPATypes: input.ActivityTypes,
		},
		PaginationInfo: &common.PaginationConfig{
			PageNumber: &pageNumber,
			PageSize:   &input.Limit,
		},
	}
	return &activityParams
}
