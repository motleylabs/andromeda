package common

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/web3"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"
)

func ConvertLamports(sol float64) int64 {
	return int64(sol * LAMPORTS_PER_SOL)
}

func GetLamports(sol float64) string {
	return fmt.Sprintf("%d", ConvertLamports(sol))
}

func GetLamportsFromPointer(sol *float64) string {
	output := "0"
	if sol != nil {
		output = fmt.Sprintf("%d", ConvertLamports(*sol))
	}
	return output
}

func GetPercentFromRatio(ratio float64) *float64 {
	percent := ratio * float64(100)
	return &percent
}

func GetPercentFromPointer(ratio *float64) *float64 {
	if ratio == nil {
		return nil
	}
	return GetPercentFromRatio(*ratio)
}

func GetLamportsFromUSDIntPointer(val *int, solPrice float64) string {
	output := "0"
	if val != nil && solPrice != 0 {
		output = fmt.Sprintf("%d", int(float64(*val)/solPrice*LAMPORTS_PER_SOL))
	}
	return output
}

func GetFromIntPointer(val *int) string {
	output := "0"
	if val != nil {
		output = fmt.Sprintf("%d", *val)
	}
	return output
}

func GetFromFloatPointer(val *float64) string {
	output := "0"
	if val != nil {
		output = fmt.Sprintf("%f", *val)
	}
	return output
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

	sort.SliceStable(res, func(i, j int) bool {
		if res[i].TraitType == "Attributes Count" {
			return true
		}

		if res[j].TraitType == "Attributes Count" {
			return false
		}

		return strings.Compare(res[i].TraitType, res[j].TraitType) < 0
	})

	return res
}

func convertActionInfo(mpaInfo *MPAInfo) *types.ActionInfo {
	if mpaInfo != nil {
		return &types.ActionInfo{
			User:                  mpaInfo.UserAddress,
			Price:                 GetLamports(mpaInfo.Price),
			Signature:             mpaInfo.Signature,
			BlockTimestamp:        mpaInfo.BlockTimestamp,
			MarketPlaceProgramID:  mpaInfo.MarketPlaceProgramID,
			MarketPlaceInstanceID: mpaInfo.MarketPlaceInstanceID,
			TradeState:            mpaInfo.Metadata.TradeState,
		}
	}
	return nil
}

var MarketPrograms = map[string]string{
	"1BWutmTvYPwDtmw9abTkS4Ssr8no61spGAvW1X6NDix":  "magiceden.v2",
	"4zdNGgAtFsW1cQgHqkiWyRsxaAgxrSRRynnuunxzjxue": "tensorswap.v1",
}

func ConvertNFTSnapshot(snapshot *MarketPlaceSnapshot, isDetail bool, snapshotOwner *string) *types.NFT {
	traits := GetTraits(&snapshot.Attributes)

	var owner *string
	if snapshotOwner != nil {
		owner = snapshotOwner
	} else {
		if snapshot.LowestListingMPA != nil {
			if snapshot.LowestListingMPA.Metadata.Seller != nil {
				owner = snapshot.LowestListingMPA.Metadata.Seller
			} else {
				owner = &snapshot.LowestListingMPA.UserAddress
			}
		}

		if isDetail || owner == nil {
			holder, _ := web3.GetMintOwner(snapshot.TokenAddress)
			if !(isDetail && owner != nil && holder != nil && MarketPrograms[*holder] != "") {
				owner = holder
			}
		}
	}

	nft := types.NFT{
		ProjectID:     snapshot.ProjectID,
		Name:          &snapshot.Name,
		Symbol:        snapshot.ProjectSlug,
		Image:         snapshot.MetadataImg,
		MintAddress:   snapshot.TokenAddress,
		MoonRank:      snapshot.MoonRank,
		Royalty:       snapshot.CreatorRoyalty,
		Owner:         owner,
		TokenStandard: snapshot.NFTStandard,
		Traits:        &traits,
		URI:           snapshot.MetadataURI,
		LastSale:      convertActionInfo(snapshot.LastSaleMPA),
		LatestListing: convertActionInfo(snapshot.LowestListingMPA),
		HighestBid:    convertActionInfo(snapshot.HighestBidMPA),
	}
	return &nft
}

func ConvertActivitySnapshots(snapshots []MarketPlaceSnapshot) []types.Activity {
	activities := make([]types.Activity, len(snapshots))

	for index := range snapshots {
		activities[index].Name = snapshots[index].Name
		activities[index].Image = snapshots[index].MetadataImg
		activities[index].Mint = snapshots[index].TokenAddress
		activities[index].Symbol = snapshots[index].ProjectSlug

		if snapshots[index].MarketPlaceState != nil {
			price := GetLamportsFromPointer(snapshots[index].MarketPlaceState.Price)
			activities[index].MarketPlaceProgramAddress = snapshots[index].MarketPlaceState.MarketPlaceProgramID
			activities[index].AuctionHouseAddress = snapshots[index].MarketPlaceState.MarketPlaceInstanceID
			activities[index].Signature = snapshots[index].MarketPlaceState.Signature
			activities[index].BlockTimestamp = snapshots[index].MarketPlaceState.BlockTimestamp
			activities[index].Seller = snapshots[index].MarketPlaceState.SellerAddress
			activities[index].Buyer = snapshots[index].MarketPlaceState.BuyerAddress
			activities[index].ActivityType = snapshots[index].MarketPlaceState.Type
			activities[index].Price = &price
		}

	}
	return activities
}

func ConvertNFTActivity(states []MarketPlaceState) []types.NFTActivity {
	activities := make([]types.NFTActivity, len(states))

	for index := range states {
		price := GetLamportsFromPointer(states[index].Price)
		activities[index].MarketPlaceProgramAddress = states[index].MarketPlaceProgramID
		activities[index].AuctionHouseAddress = states[index].MarketPlaceInstanceID
		activities[index].Signature = states[index].Signature
		activities[index].BlockTimestamp = states[index].BlockTimestamp
		activities[index].Seller = states[index].SellerAddress
		activities[index].Buyer = states[index].BuyerAddress
		activities[index].ActivityType = states[index].Type
		activities[index].Price = &price
	}
	return activities
}

func ConvertProjectStat(projectStat *ProjectStat, solPrice float64) *types.Collection {
	holders := int64(0)
	if projectStat.TokenHolders != nil {
		holders = int64(*projectStat.TokenHolders)
	}

	marketCap := float64(0)
	if projectStat.MarketCap != nil {
		marketCap = *projectStat.MarketCap
	}

	attributes := []types.AttributeOutput{}
	stat := types.Statistics{
		Volume1D:  GetLamportsFromUSDIntPointer(projectStat.Volume1Day, solPrice),
		Listed1D:  GetFromIntPointer(projectStat.Listed1Day),
		Floor1D:   GetLamportsFromPointer(projectStat.FloorPrice1Day),
		Holders:   holders,
		MarketCap: fmt.Sprintf("%d", int64(math.Round(marketCap/solPrice*LAMPORTS_PER_SOL))),
		Supply:    projectStat.Project.Supply,
	}

	var collection types.Collection
	collection.ID = projectStat.ProjectID
	collection.Description = projectStat.Project.Description
	collection.Image = projectStat.Project.ImgURL
	collection.Name = projectStat.Project.DisplayName
	collection.Statistics = &stat
	collection.Attributes = attributes
	collection.Slug = projectStat.Project.ProjectSlug
	collection.IsVerified = projectStat.Project.IsVerified
	collection.VerifiedCollectionAddress = projectStat.Project.VerifiedCollectionAddress

	return &collection
}

func GetNFTsFromAddresses(addresses []string, pageNumber, pageSize int) (*ProjectSnapshotsRes, error) {
	projectStatParams := StatParams{
		Condition: &Condition{
			TokenAddresses: &addresses,
		},
		PaginationInfo: &PaginationConfig{
			PageSize:   &pageSize,
			PageNumber: &pageNumber,
		},
	}

	payload, err := json.Marshal(projectStatParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-market-place-snapshots", ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var nftRes ProjectSnapshotsRes
	if err := json.Unmarshal(res, &nftRes); err != nil {
		return nil, err
	}

	return &nftRes, nil
}

func GetProjectsFromAddresses(addresses []string, excludeAttr bool, pageNumber, pageSize int) (*ProjectStatRes, error) {
	excludeProjectAttr := excludeAttr
	projectStatParams := StatParams{
		Conditions: &Conditions{
			ProjectIDs:               &addresses,
			ExcludeProjectAttributes: &excludeProjectAttr,
		},
		PaginationInfo: &PaginationConfig{
			PageSize:   &pageSize,
			PageNumber: &pageNumber,
		},
	}

	payload, err := json.Marshal(projectStatParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-project-stats", ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var projectStats ProjectStatRes
	if err := json.Unmarshal(res, &projectStats); err != nil {
		return nil, err
	}

	return &projectStats, nil
}

func ConvertToActivityType(activityType string) string {
	return strings.ToUpper(activityType)
}

func ConvertFromActivityType(activityType string) string {
	return strings.ToLower(activityType)
}

func ChunkAddresses(addresses []string, chunkSize int) [][]string {
	var divided [][]string

	for index := 0; index < len(addresses); index += chunkSize {
		end := index + chunkSize

		if end > len(addresses) {
			end = len(addresses)
		}

		divided = append(divided, addresses[index:end])
	}

	return divided
}
