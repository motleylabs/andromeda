package nft

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"fmt"
)

func GetDetail(address string) (*types.NFT, error) {
	tokenAddresses := []string{address}
	nftRes, err := common.GetNFTsFromAddresses(tokenAddresses, 1, 10)
	if err != nil {
		return nil, err
	}

	if len(nftRes.MarketPlaceSnapshots) == 0 {
		return nil, fmt.Errorf("no nft information found")
	}

	return common.ConvertNFTSnapshot(&nftRes.MarketPlaceSnapshots[0]), nil
}
