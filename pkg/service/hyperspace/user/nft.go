package user

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"andromeda/pkg/service/web3"
	"fmt"
	"sort"
	"strconv"

	"github.com/gin-contrib/cache/persistence"
)

func GetNFTs(address string, store *persistence.InMemoryStore) (*types.UserNFT, error) {
	go common.FetchSOLPrice(store)

	mints, err := web3.GetTokensByOwner(address)
	if err != nil {
		return nil, err
	}

	if len(mints) == 0 {
		return &types.UserNFT{
			Collections: []types.CollectedCollection{},
			NFTs:        []types.NFT{},
		}, nil
	}

	marketplaceSnapshots := []common.MarketPlaceSnapshot{}
	pageNumber := 1
	pageSize := 50
	for {
		nftRes, err := common.GetNFTsFromAddresses(mints, pageNumber, pageSize)
		if err != nil {
			return nil, err
		}

		marketplaceSnapshots = append(marketplaceSnapshots, nftRes.MarketPlaceSnapshots...)

		if nftRes.PaginationInfo.HasNextPage {
			pageNumber += 1
		} else {
			break
		}
	}

	if len(marketplaceSnapshots) == 0 {
		return &types.UserNFT{
			Collections: []types.CollectedCollection{},
			NFTs:        []types.NFT{},
		}, nil
	}

	// get project ids
	projectIDs := []string{}
	for index := range marketplaceSnapshots {
		projectID := marketplaceSnapshots[index].ProjectID
		alreadyExists := false

		for _, curID := range projectIDs {
			if projectID == curID {
				alreadyExists = true
			}
		}

		if !alreadyExists && projectID != "missing_first_creator" {
			projectIDs = append(projectIDs, projectID)
		}
	}

	solPrice, err := common.GetSOLPrice(store)
	if err != nil {
		return nil, err
	}

	// get collected collections
	collectedCollections := []types.CollectedCollection{}
	if len(projectIDs) > 0 {
		pageNumber = 1
		pageSize = 50
		for {
			projectStatRes, err := common.GetProjectsFromAddresses(projectIDs, true, pageNumber, pageSize)
			if err != nil {
				return nil, err
			}

			for index := range projectStatRes.ProjectStats {
				collection := common.ConvertProjectStat(&projectStatRes.ProjectStats[index], solPrice)
				collectedCollection := types.CollectedCollection{
					ID:        collection.ID,
					Name:      collection.Name,
					Image:     collection.Image,
					NFTsOwned: 0,
				}
				if collection.Statistics != nil {
					collectedCollection.FloorPrice = collection.Statistics.Floor1D
				}
				collectedCollections = append(collectedCollections, collectedCollection)
			}

			if projectStatRes.PaginationInfo.HasNextPage {
				pageNumber += 1
			} else {
				break
			}
		}
	}

	nfts := []types.NFT{}
	for index := range marketplaceSnapshots {
		projectID := marketplaceSnapshots[index].ProjectID

		for collectionIndex := range collectedCollections {
			if projectID == collectedCollections[collectionIndex].ID {
				collectedCollections[collectionIndex].NFTsOwned += 1
			}
		}

		nfts = append(nfts, *common.ConvertNFTSnapshot(&marketplaceSnapshots[index]))
	}

	for index := range collectedCollections {
		floorPrice, err := strconv.ParseInt(collectedCollections[index].FloorPrice, 10, 64)
		if err != nil {
			collectedCollections[index].EstimatedValue = "0"
		}
		estimatedValue := floorPrice * int64(collectedCollections[index].NFTsOwned)
		collectedCollections[index].EstimatedValue = fmt.Sprint(estimatedValue)
	}

	sort.Slice(collectedCollections, func(i, j int) bool {
		floorPriceI, _ := strconv.ParseInt(collectedCollections[i].FloorPrice, 10, 64)
		floorPriceJ, _ := strconv.ParseInt(collectedCollections[j].FloorPrice, 10, 64)
		return floorPriceI > floorPriceJ
	})

	return &types.UserNFT{
		Collections: collectedCollections,
		NFTs:        nfts,
	}, nil
}
