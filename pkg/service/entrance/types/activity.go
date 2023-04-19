package types

type ActivityType int

const (
	BID ActivityType = iota
	UPDATEBID
	CANCELBID
	LISTING
	DELISTING
	UPDATELISTING
	TRANSACTION
)

var ActivityStringTypes = []string{
	"bid",
	"update_bid",
	"cancel_bid",
	"listing",
	"delisting",
	"update_listing",
	"transaction",
}

func (activityType ActivityType) String() string {
	return ActivityStringTypes[activityType]
}

type Activity struct {
	Symbol                    string  `json:"symbol"`
	Mint                      string  `json:"mint"`
	Name                      string  `json:"name"`
	Image                     string  `json:"image"`
	Seller                    *string `json:"seller"`
	Buyer                     *string `json:"buyer"`
	Price                     *string `json:"price"`
	ActivityType              string  `json:"activityType"`
	MarketPlaceProgramAddress string  `json:"martketplaceProgramAddress"`
	AuctionHouseAddress       string  `json:"auctionHouseAddress"`
	CreatedAt                 string  `json:"createdAt"`
	Signature                 string  `json:"signature"`
}

type NFTActivity struct {
	Seller                    *string `json:"seller"`
	Buyer                     *string `json:"buyer"`
	Price                     *string `json:"price"`
	ActivityType              string  `json:"activityType"`
	MarketPlaceProgramAddress string  `json:"martketplaceProgramAddress"`
	AuctionHouseAddress       string  `json:"auctionHouseAddress"`
	CreatedAt                 string  `json:"createdAt"`
	Signature                 string  `json:"signature"`
}

type ActivityRes struct {
	Activities  []Activity `json:"activities"`
	HasNextPage bool       `json:"hasNextPage"`
}

type NFTActivityRes struct {
	Activities  []NFTActivity `json:"activities"`
	HasNextPage bool          `json:"hasNextPage"`
}

type ActivityParams struct {
	Address       string   `json:"address" example:"target address"`
	Limit         int      `json:"limit" example:"10"`
	Offset        int      `json:"offset" example:"0"`
	ActivityTypes []string `json:"activityTypes,omitempty" example:"listing"`
}
