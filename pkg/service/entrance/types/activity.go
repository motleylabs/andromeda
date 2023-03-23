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

var activityStringTypes = []string{
	"BID",
	"UPDATEBID",
	"CANCELBID",
	"LISTING",
	"DELISTING",
	"UPDATELISTING",
	"TRANSACTION",
}

func (activityType ActivityType) String() string {
	return activityStringTypes[activityType]
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
	CreatedAt                 string  `json:"createdAt"`
	Signature                 string  `json:"signature"`
}

type NFTActivity struct {
	Seller                    *string `json:"seller"`
	Buyer                     *string `json:"buyer"`
	Price                     *string `json:"price"`
	ActivityType              string  `json:"activityType"`
	MarketPlaceProgramAddress string  `json:"martketplaceProgramAddress"`
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
	Address       string   `json:"address" example:"8xBMPGAj5NzAwRmdfEcksDcZyexr87AAmD6LWwKG7Dqq"`
	Limit         int      `json:"limit" example:"10"`
	Offset        int      `json:"offset" example:"0"`
	ActivityTypes []string `json:"activityTypes,omitempty" example:"LISTING"`
}
