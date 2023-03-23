package nft

import "andromeda/pkg/service/hyperspace/common"

type NFTActivities struct {
	TokenAddress       string                    `json:"token_address"`
	MarketPlaceActions []common.MarketPlaceState `json:"market_place_actions"`
}

type NFTOffers struct {
	TokenAddress      string                       `json:"token_address"`
	MarketPlaceStates []common.MarketPlaceSnapshot `json:"market_place_states"`
}
