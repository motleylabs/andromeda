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
	Owner                     string  `json:"owner"`
	MarketPlace               string  `json:"marketplace"`
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

type ActivityParams struct {
	Address       string   `json:"collection"`
	Limit         int      `json:"limit"`
	Offset        int      `json:"offset"`
	ActivityTypes []string `json:"activityTypes"`
}
