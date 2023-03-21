package entrance

import (
	"andromeda/pkg/service/entrance/types"
)

type DataProvider interface {
	// for collection
	GetCollectionTrends(*types.TrendParams) (*types.TrendRes, error)
	GetCollectionDetail(string) (*types.Collection, error)
	GetCollectionTimeSeries(*types.TimeSeriesParams) (*types.TimeSeriesRes, error)
	GetCollectionNFTs(*types.NFTParams) (*types.NFTRes, error)
	GetCollectionActivities(*types.ActivityParams) (*types.ActivityResult, error)

	// for nft
	GetNFTDetail(string) (*types.NFT, error)
	GetNFTActivities(*types.ActivityParams) (*types.ActivityResult, error)

	// for wallet
	GetWalletNFTs(string) (*types.WalletNFT, error)
	GetWalletActivities(string) ([]types.Activity, error)
}
