package user

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

func GetOffers(params *types.ActivityParams) (*types.ActivityRes, error) {
	if params == nil {
		return nil, fmt.Errorf("no activity params")
	}
	offerParams := getUserOfferParams(params)
	payload, err := json.Marshal(offerParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-user-bids", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var snapshots common.ProjectSnapshotsRes
	if err := json.Unmarshal(res, &snapshots); err != nil {
		return nil, err
	}

	activityRes := types.ActivityRes{
		HasNextPage: snapshots.PaginationInfo.HasNextPage,
		Activities:  common.ConvertActivitySnapshots(snapshots.MarketPlaceSnapshots),
	}

	return &activityRes, nil
}

func getUserOfferParams(input *types.ActivityParams) *OfferParams {
	actionType := "BID"
	pageNumber := input.Offset/input.Limit + 1

	var offerParams = OfferParams{
		Condition: OfferCondition{
			BuyerAddress: input.Address,
			ActionType:   &actionType,
		},
		OrderBy: common.OrderConfig{
			FieldName: "block_timestamp",
			SortOrder: "DESC",
		},
		PaginationInfo: &common.PaginationConfig{
			PageNumber: &pageNumber,
			PageSize:   &input.Limit,
		},
	}

	return &offerParams
}
