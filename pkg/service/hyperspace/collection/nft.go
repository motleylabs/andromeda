package collection

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

func GetNFTs(params *types.NFTParams) (*types.NFTRes, error) {
	if params == nil {
		return nil, fmt.Errorf("no nft params")
	}
	nftParams := GetNFTParams(params)
	payload, err := json.Marshal(nftParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-market-place-snapshots", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var nftRes ProjectNFTsRes
	if err := json.Unmarshal(res, &nftRes); err != nil {
		return nil, err
	}

	nftsRes := types.NFTRes{
		HasNextPage: nftRes.PaginationInfo.HasNextPage,
		NFTs:        ConvertNFTSnapshot(nftRes.MarketPlaceSnapshots),
	}

	return &nftsRes, nil
}

func ConvertNFTSnapshot(snapshots []MarketPlaceSnapshot) []types.NFT {
	nfts := make([]types.NFT, len(snapshots))

	for index := range snapshots {
		var lastSold *string
		if snapshots[index].LastSaleMPA != nil {
			lastSoldStr := common.GetLamports(snapshots[index].LastSaleMPA.Price)
			lastSold = &lastSoldStr
		}
		traits := GetTraits(&snapshots[index].Attributes)

		nfts[index].Name = &snapshots[index].Name
		nfts[index].Image = snapshots[index].MetadataImg
		nfts[index].LastSold = lastSold
		nfts[index].MintAddress = snapshots[index].TokenAddress
		nfts[index].MoonRank = snapshots[index].MoonRank
		nfts[index].Royalty = snapshots[index].CreatorRoyalty
		nfts[index].Owner = snapshots[index].Owner
		nfts[index].TokenStandard = snapshots[index].NFTStandard
		nfts[index].Traits = &traits
		nfts[index].URI = snapshots[index].MetadataURI
	}

	return nfts
}

func GetTraits(attributes *map[string]interface{}) []types.Trait {
	res := []types.Trait{}
	if attributes == nil {
		return res
	}

	for key, value := range *attributes {
		res = append(res, types.Trait{
			TraitType: key,
			Value:     fmt.Sprintf("%v", value),
		})
	}

	return res
}

func GetNFTParams(input *types.NFTParams) *common.StatParams {
	var attributes *[]types.Attribute
	if len(input.Attributes) > 0 {
		attributes = &input.Attributes
	}
	projectIDs := []common.ProjectIDItem{
		{
			ProjectID:  input.Address,
			Attributes: attributes,
		},
	}
	return &common.StatParams{
		Condition: &common.Condition{
			ProjectIDs: &projectIDs,
		},
		OrderBy:        GetNFTOrderField(input),
		PaginationInfo: GetNFTPaginationInfo(input),
	}
}

func GetNFTOrderField(input *types.NFTParams) *common.OrderConfig {
	orderFieldName := "lowest_listing_price"

	switch input.SortBy {
	case "listing_timestamp":
		orderFieldName = "lowest_listing_block_timestamp"
	}

	return &common.OrderConfig{
		FieldName: orderFieldName,
		SortOrder: input.Order,
	}
}

func GetNFTPaginationInfo(input *types.NFTParams) *common.PaginationConfig {
	pageNumber := input.Offset/input.Limit + 1
	pageSize := input.Limit

	paginationInfo := common.PaginationConfig{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}

	return &paginationInfo
}
