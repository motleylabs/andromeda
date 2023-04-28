package hyperspace

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/collection"
	"andromeda/pkg/service/hyperspace/nft"
	"andromeda/pkg/service/hyperspace/stat"
	"andromeda/pkg/service/hyperspace/user"
	"fmt"
	"time"

	"github.com/gin-contrib/cache/persistence"
)

type Hyperspace struct{}

// store that saves the following information
// slug data
// sol price
var store = persistence.NewInMemoryStore(0)

func getCollectionIDFromSlug(slug string) (string, error) {
	collectionID := ""
	if err := store.Get(slug, &collectionID); err != nil {
		col, _ := collection.GetDetail(slug, store)
		if col == nil {
			store.Set(slug, "NOT_FOUND", 30*time.Second)
			return "", fmt.Errorf("slug %s not registered", slug)
		} else {
			collectionID = col.ID
		}
	} else if collectionID == "NOT_FOUND" {
		return "", fmt.Errorf("slug %s not registered", slug)
	}

	return collectionID, nil
}

func (Hyperspace) GetStatOverall() (*types.StatRes, error) {
	return stat.GetOverall()
}

func (Hyperspace) GetStatSearch(params *types.SearchParams) (*types.SearchRes, error) {
	return stat.GetHits(params, store)
}

func (Hyperspace) GetCollectionTrends(params *types.TrendParams) (*types.TrendRes, error) {
	return collection.GetTrends(params, store)
}

func (Hyperspace) GetCollectionDetail(slug string) (*types.Collection, error) {
	collectionID := ""
	if err := store.Get(slug, &collectionID); err != nil {
		collectionID = slug
	}
	return collection.GetDetail(collectionID, store)
}

func (Hyperspace) GetCollectionTimeSeries(params *types.TimeSeriesParams) (*types.TimeSeriesRes, error) {
	collectionID, err := getCollectionIDFromSlug(params.Address)
	if err != nil {
		return nil, err
	}

	params.Address = collectionID
	return collection.GetTimeSeries(params)
}

func (Hyperspace) GetCollectionNFTs(params *types.NFTParams) (*types.NFTRes, error) {
	collectionID, err := getCollectionIDFromSlug(params.Address)
	if err != nil {
		return nil, err
	}

	params.Address = collectionID
	return collection.GetNFTs(params)
}

func (Hyperspace) GetCollectionActivities(params *types.ActivityParams) (*types.ActivityRes, error) {
	collectionID, err := getCollectionIDFromSlug(params.Address)
	if err != nil {
		return nil, err
	}

	params.Address = collectionID
	return collection.GetActivities(params)
}

func (Hyperspace) GetNFTDetail(address string) (*types.NFT, error) {
	return nft.GetDetail(address, store)
}

func (Hyperspace) GetNFTActivities(params *types.ActivityParams) (*types.NFTActivityRes, error) {
	return nft.GetActivities(params)
}

func (Hyperspace) GetNFTOffers(address string) ([]types.NFTActivity, error) {
	return nft.GetOffers(address)
}

func (Hyperspace) GetUserNFTs(address string) (*types.UserNFT, error) {
	return user.GetNFTs(address, store)
}

func (Hyperspace) GetUserActivities(params *types.ActivityParams) (*types.ActivityRes, error) {
	return user.GetActivities(params)
}

func (Hyperspace) GetUserOffers(params *types.ActivityParams) (*types.ActivityRes, error) {
	return user.GetOffers(params)
}
