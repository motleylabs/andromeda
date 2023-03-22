package nft

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

func GetDetail(address string) (*types.NFT, error) {
	tokenAddresses := []string{address}
	projectStatParams := common.StatParams{
		Condition: &common.Condition{
			TokenAddresses: &tokenAddresses,
		},
	}

	payload, err := json.Marshal(projectStatParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-market-place-snapshots", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var nftRes common.ProjectSnapshotsRes
	if err := json.Unmarshal(res, &nftRes); err != nil {
		return nil, err
	}

	if len(nftRes.MarketPlaceSnapshots) == 0 {
		return nil, fmt.Errorf("no nft information found")
	}

	return common.ConvertNFTSnapshot(&nftRes.MarketPlaceSnapshots[0]), nil
}
