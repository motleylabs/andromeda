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
	GetCollectionActivities(*types.ActivityParams) (*types.ActivityRes, error)

	// for nft
	GetNFTDetail(string) (*types.NFT, error)
	GetNFTActivities(*types.ActivityParams) (*types.NFTActivityRes, error)
	GetNFTOffers(string) ([]types.NFTActivity, error)

	// for wallet
	GetUserNFTs(string) (*types.UserNFT, error)
	GetUserActivities(*types.ActivityParams) (*types.ActivityRes, error)
	GetUserOffers(*types.ActivityParams) (*types.NFTActivityRes, error)
}
