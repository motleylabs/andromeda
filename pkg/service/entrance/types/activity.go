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
	Symbol                    string       `json:"symbol"`
	Mint                      string       `json:"mint"`
	Name                      string       `json:"name"`
	Image                     string       `json:"image"`
	Owner                     string       `json:"owner"`
	MarketPlace               string       `json:"marketplace"`
	Price                     *string      `json:"price"`
	ActivityType              ActivityType `json:"activityType"`
	MarketPlaceProgramAddress string       `json:"martketplaceProgramAddress"`
	CreatedAt                 string       `json:"createdAt"`
	Signature                 string       `json:"signature"`
}

type ActivityResult struct {
	Activities []Activity `json:"activities"`
	LastToken  *string    `json:"lastToken"`
}

type ActivityParams struct {
	Address       *string
	Limit         int
	Offset        *int
	Next          *string
	ActivityTypes []ActivityType
}
