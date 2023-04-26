package nft

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"fmt"

	"github.com/gin-contrib/cache/persistence"
)

func GetDetail(address string, store *persistence.InMemoryStore) (*types.NFT, error) {
	tokenAddresses := []string{address}
	nftRes, err := common.GetNFTsFromAddresses(tokenAddresses, 1, 10)
	if err != nil {
		return nil, err
	}

	if len(nftRes.MarketPlaceSnapshots) == 0 {
		return nil, fmt.Errorf("no nft information found")
	}

	curSlug := nftRes.MarketPlaceSnapshots[0].ProjectSlug
	curID := nftRes.MarketPlaceSnapshots[0].ProjectID
	if err := store.Set(curSlug, curID, -1); err != nil {
		return nil, err
	}

	return common.ConvertNFTSnapshot(&nftRes.MarketPlaceSnapshots[0], true, nil), nil
}
