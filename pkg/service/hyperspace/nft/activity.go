package nft

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

func GetActivities(params *types.ActivityParams) (*types.NFTActivityRes, error) {
	if params == nil {
		return nil, fmt.Errorf("no activity params")
	}
	activityParams := getNFTActivityParams(params)
	payload, err := json.Marshal(activityParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-market-place-actions-by-token", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var activities []NFTActivities
	if err := json.Unmarshal(res, &activities); err != nil {
		return nil, err
	}

	if len(activities) == 0 {
		return nil, fmt.Errorf("unexpected error")
	}

	activityRes := types.NFTActivityRes{
		Activities:  common.ConvertNFTActivity(activities[0].MarketPlaceActions),
		HasNextPage: false,
	}
	return &activityRes, nil
}

func getNFTActivityParams(input *types.ActivityParams) *common.StatParams {
	tokenAddress := []string{input.Address}
	orderConfig := common.OrderConfig{
		FieldName: "block_timestamp",
		SortOrder: "DESC",
	}

	var actionType *string
	if len(input.ActivityTypes) > 0 {
		actionType = &input.ActivityTypes[0]
	}

	var activityParams = common.StatParams{
		Condition: &common.Condition{
			TokenAddresses: &tokenAddress,
			ActionType:     actionType,
		},
		OrderBy: &orderConfig,
	}
	return &activityParams
}
