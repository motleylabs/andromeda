package hyperspace

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/collection"
	"andromeda/pkg/service/hyperspace/nft"
	"andromeda/pkg/service/hyperspace/user"
	"fmt"
	"time"

	"github.com/gin-contrib/cache/persistence"
)

type Hyperspace struct{}

var slugStore = persistence.NewInMemoryStore(time.Second)

func (Hyperspace) GetCollectionTrends(params *types.TrendParams) (*types.TrendRes, error) {
	return collection.GetTrends(params, slugStore)
}

func (Hyperspace) GetCollectionDetail(slug string) (*types.Collection, error) {
	collectionID := ""
	if err := slugStore.Get(slug, &collectionID); err != nil {
		return nil, fmt.Errorf("slug %s not registered", slug)
	}
	return collection.GetDetail(collectionID)
}

func (Hyperspace) GetCollectionTimeSeries(params *types.TimeSeriesParams) (*types.TimeSeriesRes, error) {
	collectionID := ""
	if err := slugStore.Get(params.Address, &collectionID); err != nil {
		return nil, fmt.Errorf("slug %s not registered", params.Address)
	}
	params.Address = collectionID
	return collection.GetTimeSeries(params)
}

func (Hyperspace) GetCollectionNFTs(params *types.NFTParams) (*types.NFTRes, error) {
	collectionID := ""
	if err := slugStore.Get(params.Address, &collectionID); err != nil {
		return nil, fmt.Errorf("slug %s not registered", params.Address)
	}
	params.Address = collectionID
	return collection.GetNFTs(params)
}

func (Hyperspace) GetCollectionActivities(params *types.ActivityParams) (*types.ActivityRes, error) {
	collectionID := ""
	if err := slugStore.Get(params.Address, &collectionID); err != nil {
		return nil, fmt.Errorf("slug %s not registered", params.Address)
	}
	params.Address = collectionID
	return collection.GetActivities(params)
}

func (Hyperspace) GetNFTDetail(address string) (*types.NFT, error) {
	return nft.GetDetail(address)
}

func (Hyperspace) GetNFTActivities(params *types.ActivityParams) (*types.NFTActivityRes, error) {
	return nft.GetActivities(params)
}

func (Hyperspace) GetNFTOffers(address string) ([]types.NFTActivity, error) {
	return nft.GetOffers(address)
}

func (Hyperspace) GetUserNFTs(address string) (*types.UserNFT, error) {
	return user.GetNFTs(address)
}

func (Hyperspace) GetUserActivities(params *types.ActivityParams) (*types.ActivityRes, error) {
	return user.GetActivities(params)
}

func (Hyperspace) GetUserOffers(params *types.ActivityParams) (*types.ActivityRes, error) {
	return user.GetOffers(params)
}
