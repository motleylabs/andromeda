package types

type Trait struct {
	Value     string  `json:"value"`
	TraitType string  `json:"traitType"`
	Rarity    float64 `json:"rarity,omitempty"`
}

var TokenStandard = []string{
	"FUNGIBLE",
	"FUNGIBLE_ASSET",
	"NON_FUNGIBLE",
	"NON_FUNGIBLE_EDITION",
	"PROGRAMMABLE_NON_FUNGIBLE",
}

type NFT struct {
	ProjectID     string   `json:"projectId"`
	MintAddress   string   `json:"mintAddress"`
	Symbol        string   `json:"symbol"`
	URI           string   `json:"uri"`
	Traits        *[]Trait `json:"traits"`
	Name          *string  `json:"name"`
	Description   string   `json:"description"`
	Royalty       int      `json:"sellerFeeBasisPoints"`
	ListingType   string   `json:"listingType"`
	ListingPrice  *string  `json:"listingPrice"`
	Image         string   `json:"image"`
	Owner         string   `json:"owner"`
	LastSold      *string  `json:"lastSold"`
	TokenStandard string   `json:"tokenStandard"`
	MoonRank      int      `json:"moonrankRank"`
}

type NFTRes struct {
	NFTs        []NFT `json:"nfts"`
	HasNextPage bool  `json:"hasNextPage"`
}
