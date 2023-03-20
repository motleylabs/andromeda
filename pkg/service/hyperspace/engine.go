package hyperspace

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/collection"
)

type Hyperspace struct{}

func (Hyperspace) GetCollectionTrends(params *types.TrendParams) (*types.TrendRes, error) {
	return collection.GetTrends(params)
}

func (Hyperspace) GetCollectionStat(address string) (*types.Statistics, error) {
	return nil, nil
}

func (Hyperspace) GetCollectionTimeSeries(params *types.TimeSeriesParams) ([]types.TimeSeries, error) {
	return []types.TimeSeries{}, nil
}

func (Hyperspace) GetCollectionNFTs(string, string) ([]types.NFT, error) {
	return []types.NFT{}, nil
}

func (Hyperspace) GetCollectionActivities(*types.ActivityParams) (*types.ActivityResult, error) {
	return nil, nil
}

func (Hyperspace) GetNFTDetail(string) (*types.NFT, error) {
	return nil, nil
}

func (Hyperspace) GetNFTActivities(*types.ActivityParams) (*types.ActivityResult, error) {
	return nil, nil
}

func (Hyperspace) GetWalletNFTs(string) (*types.WalletNFT, error) {
	return nil, nil
}

func (Hyperspace) GetWalletActivities(string) ([]types.Activity, error) {
	return []types.Activity{}, nil
}
