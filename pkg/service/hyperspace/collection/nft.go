package collection

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
	"strings"
)

func GetNFTs(params *types.NFTParams) (*types.NFTRes, error) {
	if params == nil {
		return nil, fmt.Errorf("no nft params")
	}
	nftParams := getNFTParams(params)
	payload, err := json.Marshal(nftParams)
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

	nftsRes := types.NFTRes{
		HasNextPage: nftRes.PaginationInfo.HasNextPage && len(nftRes.MarketPlaceSnapshots) == params.Limit,
		NFTs:        convertNFTSnapshots(nftRes.MarketPlaceSnapshots),
	}

	return &nftsRes, nil
}

func convertNFTSnapshots(snapshots []common.MarketPlaceSnapshot) []types.NFT {
	nfts := make([]types.NFT, len(snapshots))

	for index := range snapshots {
		nfts[index] = *common.ConvertNFTSnapshot(&snapshots[index], false, nil)
	}

	return nfts
}

func getNFTParams(input *types.NFTParams) *common.StatParams {
	// set program ids
	var attributes *[]types.AttributeInput
	if len(input.Attributes) > 0 {
		attributes = &input.Attributes
	}
	projectIDs := []common.ProjectIDItem{
		{
			ProjectID:  input.Address,
			Attributes: attributes,
		},
	}

	// set name
	var nameCond *common.NameParam
	if input.Name != nil {
		nameCond = &common.NameParam{
			Operation: "FUZZY",
			Value:     *input.Name,
		}
	}

	// set price
	var priceFilter *common.PriceFilter
	if input.PriceMin != nil || input.PriceMax != nil {
		priceFilter = &common.PriceFilter{
			Min:   input.PriceMin,
			Max:   input.PriceMax,
			Field: "lowest_listing_price",
		}
	}

	// set marketplace program condition
	var marketplaceProgramCondition *common.MarketPlaceProgramCondition
	if input.Program != nil {
		programCondition := common.MarketPlaceProgramCondition{
			MarketPlacePrograms: []common.MarketPlaceProgram{
				{
					MarketPlaceProgramID: *input.Program,
				},
			},
		}
		if input.AuctionHouse != nil {
			programCondition.MarketPlacePrograms[0].MarketPlaceInstanceID = *input.AuctionHouse
		}
		marketplaceProgramCondition = &programCondition
	}

	// set listing type
	var listingType *string
	if input.ListingOnly {
		listingTypeStr := "LISTING"
		listingType = &listingTypeStr
	}

	return &common.StatParams{
		Condition: &common.Condition{
			ProjectIDs:                  &projectIDs,
			MarketPlaceProgramCondition: marketplaceProgramCondition,
			ListingType:                 listingType,
			Name:                        nameCond,
			PriceFilter:                 priceFilter,
		},
		OrderBy:        getNFTOrderField(input),
		PaginationInfo: getNFTPaginationInfo(input),
	}
}

func getNFTOrderField(input *types.NFTParams) *common.OrderConfig {
	orderFieldName := "lowest_listing_price"

	switch input.SortBy {
	case "timestamp":
		orderFieldName = "lowest_listing_block_timestamp"
	case "price":
		orderFieldName = "lowest_listing_price"
	case "marketplace":
		orderFieldName = "lowest_listing_marketplace_program_id"
	default:
		orderFieldName = input.SortBy
	}

	return &common.OrderConfig{
		FieldName: orderFieldName,
		SortOrder: strings.ToUpper(input.Order),
	}
}

func getNFTPaginationInfo(input *types.NFTParams) *common.PaginationConfig {
	pageNumber := input.Offset/input.Limit + 1
	pageSize := input.Limit

	paginationInfo := common.PaginationConfig{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}

	return &paginationInfo
}
