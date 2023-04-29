package entrance

import (
	"andromeda/pkg/service/entrance/types"
)

type DataProvider interface {
	// for overall stats
	GetStatOverall() (*types.StatRes, error)
	GetStatSearch(*types.SearchParams) (*types.SearchRes, error)

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
	GetNFTBuyNowTx(*types.BuyParams) ([]byte, error)

	// for wallet
	GetUserNFTs(string) (*types.UserNFT, error)
	GetUserActivities(*types.ActivityParams) (*types.ActivityRes, error)
	GetUserOffers(*types.ActivityParams) (*types.ActivityRes, error)
}
