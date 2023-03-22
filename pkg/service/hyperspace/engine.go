package hyperspace

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/collection"
	"andromeda/pkg/service/hyperspace/nft"
	"andromeda/pkg/service/hyperspace/user"
)

type Hyperspace struct{}

func (Hyperspace) GetCollectionTrends(params *types.TrendParams) (*types.TrendRes, error) {
	return collection.GetTrends(params)
}

func (Hyperspace) GetCollectionDetail(address string) (*types.Collection, error) {
	return collection.GetDetail(address)
}

func (Hyperspace) GetCollectionTimeSeries(params *types.TimeSeriesParams) (*types.TimeSeriesRes, error) {
	return collection.GetTimeSeries(params)
}

func (Hyperspace) GetCollectionNFTs(params *types.NFTParams) (*types.NFTRes, error) {
	return collection.GetNFTs(params)
}

func (Hyperspace) GetCollectionActivities(params *types.ActivityParams) (*types.ActivityRes, error) {
	return collection.GetActivities(params)
}

func (Hyperspace) GetNFTDetail(address string) (*types.NFT, error) {
	return nft.GetDetail(address)
}

func (Hyperspace) GetNFTActivities(params *types.ActivityParams) (*types.ActivityRes, error) {
	return nft.GetActivities(params)
}

func (Hyperspace) GetUserNFTs(address string) (*types.UserNFT, error) {
	return nil, nil
}

func (Hyperspace) GetUserActivities(params *types.ActivityParams) (*types.ActivityRes, error) {
	return user.GetActivities(params)
}
