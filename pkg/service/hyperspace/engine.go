package hyperspace

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/collection"
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
	return nil, nil
}

func (Hyperspace) GetNFTActivities(params *types.ActivityParams) (*types.ActivityRes, error) {
	return nil, nil
}

func (Hyperspace) GetWalletNFTs(address string) (*types.WalletNFT, error) {
	return nil, nil
}

func (Hyperspace) GetWalletActivities(address string) ([]types.Activity, error) {
	return []types.Activity{}, nil
}
