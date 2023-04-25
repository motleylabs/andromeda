package nft

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

func GetOffers(address string) ([]types.NFTActivity, error) {
	tokenAddresses := []string{address}
	actionType := "BID"
	projectStatParams := common.StatParams{
		Condition: &common.Condition{
			TokenAddresses: &tokenAddresses,
			ActionType:     &actionType,
		},
	}

	payload, err := json.Marshal(projectStatParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-token-state", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var offers []NFTOffers
	if err := json.Unmarshal(res, &offers); err != nil {
		return nil, err
	}

	if len(offers) == 0 {
		return nil, fmt.Errorf("unexpected error")
	}

	return convertNFTActivitySnapshots(offers[0].MarketPlaceStates), nil
}

func convertNFTActivitySnapshots(snapshots []common.MarketPlaceSnapshot) []types.NFTActivity {
	activities := make([]types.NFTActivity, len(snapshots))

	for index := range snapshots {
		if snapshots[index].MarketPlaceState != nil {
			price := common.GetLamportsFromPointer(snapshots[index].MarketPlaceState.Price)
			activities[index].MarketPlaceProgramAddress = snapshots[index].MarketPlaceState.MarketPlaceProgramID
			activities[index].AuctionHouseAddress = snapshots[index].MarketPlaceState.MarketPlaceInstanceID
			activities[index].Signature = snapshots[index].MarketPlaceState.Signature
			activities[index].BlockTimestamp = snapshots[index].MarketPlaceState.BlockTimestamp
			activities[index].Seller = snapshots[index].MarketPlaceState.SellerAddress
			activities[index].Buyer = snapshots[index].MarketPlaceState.BuyerAddress
			activities[index].ActivityType = snapshots[index].MarketPlaceState.Type
			activities[index].Price = &price
		}

	}
	return activities
}
