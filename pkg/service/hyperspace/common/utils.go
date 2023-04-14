package common

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"encoding/json"
	"fmt"
	"math"
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

	return res
}

func ConvertNFTSnapshot(snapshot *MarketPlaceSnapshot) *types.NFT {
	// set last sold
	var lastSold *string
	if snapshot.LastSaleMPA != nil {
		lastSoldStr := GetLamports(snapshot.LastSaleMPA.Price)
		lastSold = &lastSoldStr
	}

	// set listing price
	var listingPrice *string
	if snapshot.LowestListingMPA != nil {
		listingPriceStr := GetLamports(snapshot.LowestListingMPA.Price)
		listingPrice = &listingPriceStr
	}

	traits := GetTraits(&snapshot.Attributes)

	nft := types.NFT{
		ProjectID:     snapshot.ProjectID,
		Name:          &snapshot.Name,
		Symbol:        snapshot.ProjectSlug,
		Image:         snapshot.MetadataImg,
		LastSold:      lastSold,
		ListingPrice:  listingPrice,
		MintAddress:   snapshot.TokenAddress,
		MoonRank:      snapshot.MoonRank,
		Royalty:       snapshot.CreatorRoyalty,
		Owner:         snapshot.Owner,
		TokenStandard: snapshot.NFTStandard,
		Traits:        &traits,
		URI:           snapshot.MetadataURI,
	}
	return &nft
}

func ConvertActivitySnapshots(snapshots []MarketPlaceSnapshot) []types.Activity {
	activities := make([]types.Activity, len(snapshots))

	for index := range snapshots {
		activities[index].Name = snapshots[index].Name
		activities[index].Image = snapshots[index].MetadataImg
		activities[index].Mint = snapshots[index].TokenAddress

		if snapshots[index].MarketPlaceState != nil {
			price := GetLamportsFromPointer(snapshots[index].MarketPlaceState.Price)
			activities[index].MarketPlaceProgramAddress = snapshots[index].MarketPlaceState.MarketPlaceProgramID
			activities[index].Signature = snapshots[index].MarketPlaceState.Signature
			activities[index].CreatedAt = snapshots[index].MarketPlaceState.CreatedAt
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
		activities[index].Signature = states[index].Signature
		activities[index].CreatedAt = states[index].CreatedAt
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

	attributes := []types.Attribute{}
	if projectStat.Project.Attributes != nil {
		attributes = *projectStat.Project.Attributes
	}

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

func GetProjectsFromAddresses(addresses []string, pageNumber, pageSize int) (*ProjectStatRes, error) {
	excludeProjectAttr := false
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
