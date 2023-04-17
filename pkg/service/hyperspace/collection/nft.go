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
		HasNextPage: nftRes.PaginationInfo.HasNextPage,
		NFTs:        convertNFTSnapshots(nftRes.MarketPlaceSnapshots),
	}

	return &nftsRes, nil
}

func convertNFTSnapshots(snapshots []common.MarketPlaceSnapshot) []types.NFT {
	nfts := make([]types.NFT, len(snapshots))

	for index := range snapshots {
		nfts[index] = *common.ConvertNFTSnapshot(&snapshots[index])
	}

	return nfts
}

func getNFTParams(input *types.NFTParams) *common.StatParams {
	// set program ids
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
	}

	fmt.Println(orderFieldName)

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
