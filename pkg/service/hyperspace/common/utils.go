package common

import (
	"andromeda/pkg/service/entrance/types"
	"fmt"
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
	var lastSold *string
	if snapshot.LastSaleMPA != nil {
		lastSoldStr := GetLamports(snapshot.LastSaleMPA.Price)
		lastSold = &lastSoldStr
	}
	traits := GetTraits(&snapshot.Attributes)

	nft := types.NFT{
		Name:          &snapshot.Name,
		Image:         snapshot.MetadataImg,
		LastSold:      lastSold,
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
		activities[index].Owner = snapshots[index].Owner

		if snapshots[index].MarketPlaceState != nil {
			price := GetLamportsFromPointer(snapshots[index].MarketPlaceState.Price)
			activities[index].MarketPlaceProgramAddress = snapshots[index].MarketPlaceState.MarketPlaceProgramID
			activities[index].Signature = snapshots[index].MarketPlaceState.Signature
			activities[index].CreatedAt = snapshots[index].MarketPlaceState.CreatedAt
			activities[index].ActivityType = snapshots[index].MarketPlaceState.Type
			activities[index].Price = &price
		}

	}
	return activities
}
