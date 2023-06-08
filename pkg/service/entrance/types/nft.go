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

type ActionInfo struct {
	User                  string  `json:"userAddress"`
	Price                 string  `json:"price"`
	Signature             string  `json:"signature"`
	BlockTimestamp        int64   `json:"blockTimestamp"`
	MarketPlaceProgramID  string  `json:"auctionHouseProgram"`
	MarketPlaceInstanceID string  `json:"auctionHouseAddress"`
	TradeState            *string `json:"tradeState,omitempty"`
}

type NFT struct {
	ProjectID     string      `json:"projectId"`
	ProjectName   *string     `json:"projectName"`
	MintAddress   string      `json:"mintAddress"`
	Symbol        string      `json:"symbol"`
	URI           string      `json:"uri"`
	Traits        *[]Trait    `json:"traits"`
	Name          *string     `json:"name"`
	Description   string      `json:"description"`
	Royalty       int         `json:"sellerFeeBasisPoints"`
	Image         string      `json:"image"`
	Owner         *string     `json:"owner"`
	TokenStandard string      `json:"tokenStandard"`
	MoonRank      int         `json:"moonrankRank"`
	HighestBid    *ActionInfo `json:"highestBid"`
	LatestListing *ActionInfo `json:"latestListing"`
	LastSale      *ActionInfo `json:"lastSale"`
}

type NFTRes struct {
	NFTs        []NFT `json:"nfts"`
	HasNextPage bool  `json:"hasNextPage"`
}

type BuyParams struct {
	AuctionHouseProgram string `json:"auctionHouseProgram"`
	AuctionHouseAddress string `json:"auctionHouseAddress"`
	Seller              string `json:"seller"`
	Buyer               string `json:"buyer"`
	BuyerBroker         string `json:"buyerBroker"`
	Price               string `json:"price"`
	Mint                string `json:"mint"`
}

type BuyRes struct {
	Data   []byte `json:"data"`
	Buffer []byte `json:"buffer"`
}
