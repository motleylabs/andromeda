package nft

import "andromeda/pkg/service/hyperspace/common"

type NFTActivities struct {
	TokenAddress       string                    `json:"token_address"`
	MarketPlaceActions []common.MarketPlaceState `json:"market_place_actions"`
}
