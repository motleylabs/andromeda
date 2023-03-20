package entrance

import (
	"andromeda/pkg/service/entrance/types"
)

type DataProvider interface {
	// for collection
	GetCollectionTrends(*types.TrendParams) (*types.TrendRes, error)
	GetCollectionStat(string) (*types.Statistics, error)
	GetCollectionTimeSeries(*types.TimeSeriesParams) ([]types.TimeSeries, error)
	GetCollectionNFTs(string, string) ([]types.NFT, error)
	GetCollectionActivities(*types.ActivityParams) (*types.ActivityResult, error)

	// for nft
	GetNFTDetail(string) (*types.NFT, error)
	GetNFTActivities(*types.ActivityParams) (*types.ActivityResult, error)

	// for wallet
	GetWalletNFTs(string) (*types.WalletNFT, error)
	GetWalletActivities(string) ([]types.Activity, error)
}
