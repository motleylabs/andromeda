package types

type Statistics struct {
	Floor1D      string `json:"floor1d"`
	Volume1D     string `json:"volume1d"`
	Listed1D     string `json:"listed1d"`
	MarketCap    string `json:"marketCap"`
	MarketCapSol string `json:"marketCapSol"`
	Holders      int64  `json:"holders"`
	Supply       int64  `json:"supply"`
}

type AttributeInput struct {
	Name   string   `json:"name" example:"Background"`
	Type   string   `json:"type" example:"CATEGORY"`
	Values []string `json:"values" example:"Vivid-tangerine"`
}

type AttributeStat struct {
	Value      string  `json:"value"`
	Counts     int     `json:"counts"`
	FloorPrice *string `json:"floorPrice"`
	Listed     int     `json:"listed"`
}

type AttributeOutput struct {
	Name   string          `json:"name" example:"Background"`
	Type   string          `json:"type" example:"CATEGORY"`
	Values []AttributeStat `json:"values"`
}

type Collection struct {
	ID                        string            `json:"id"`
	Name                      string            `json:"name"`
	IsVerified                bool              `json:"isVerified"`
	VerifiedCollectionAddress string            `json:"verifiedCollectionAddress"`
	Description               string            `json:"description"`
	Image                     string            `json:"image"`
	Symbol                    string            `json:"symbol"`
	Slug                      string            `json:"slug"`
	Twitter                   *string           `json:"twitter"`
	Discord                   *string           `json:"discord"`
	Website                   *string           `json:"website"`
	Statistics                *Statistics       `json:"statistics,omitempty"`
	Attributes                []AttributeOutput `json:"attributes"`
}

type Trend struct {
	Floor1D         string     `json:"floor1d"`
	ChangeFloor1D   *float64   `json:"changeFloor1d"`
	Volume1H        string     `json:"volume1h"`
	ChangeVolume1H  *float64   `json:"changeVolume1h,omitempty"`
	Volume1D        string     `json:"volume1d"`
	ChangeVolume1D  *float64   `json:"changeVolume1d"`
	Volume7D        string     `json:"volume7d"`
	ChangeVolume7D  *float64   `json:"changeVolume7d,omitempty"`
	Volume30D       string     `json:"volume30d"`
	ChangeVolume30D *float64   `json:"changeVolume30d,omitempty"`
	Listed1D        string     `json:"listed1d"`
	ChangeListed1D  *float64   `json:"changeListed1d,omitempty"`
	Collection      Collection `json:"collection"`
}

type TrendRes struct {
	Trends      []Trend `json:"trends"`
	HasNextPage bool    `json:"hasNextPage"`
}

type TimeSeries struct {
	Timestamp  string `json:"timestamp"`
	FloorPrice string `json:"floorPrice"`
	Listed     int    `json:"listed"`
	Holders    int    `json:"holders"`
	Volume     int    `json:"volume"`
}

type TimeSeriesRes struct {
	Series      []TimeSeries `json:"series"`
	HasNextPage bool         `json:"hasNextPage"`
}

type TrendParams struct {
	Period string `json:"period" example:"1d|7d|1m"`
	SortBy string `json:"sortBy" example:"volume"`
	Order  string `json:"order" example:"asc|desc"`
	Limit  int    `json:"limit" example:"10"`
	Offset int    `json:"offset" example:"0"`
}

type TimeSeriesParams struct {
	Address     string `json:"collection" example:"8xBMPGAj5NzAwRmdfEcksDcZyexr87AAmD6LWwKG7Dqq"`
	FromTime    int    `json:"startTimestamp" example:"1671128400"`
	ToTime      int    `json:"endTimestamp" example:"1679410436"`
	Granularity string `json:"timeGranularity" example:"per_hour|per_day"`
	Limit       int    `json:"limit" example:"10"`
	Offset      int    `json:"offset" example:"0"`
}

type NFTParams struct {
	Address      string           `json:"collection" example:"8xBMPGAj5NzAwRmdfEcksDcZyexr87AAmD6LWwKG7Dqq"`
	Attributes   []AttributeInput `json:"attributes,omitempty"`
	SortBy       string           `json:"sortBy" example:"timestamp|price"`
	Order        string           `json:"order" example:"asc|desc"`
	Limit        int              `json:"limit" example:"10"`
	Offset       int              `json:"offset" example:"0"`
	ListingOnly  bool             `json:"listingOnly,omitempty"`
	Program      *string          `json:"program,omitempty" example:"RwDDvPp7ta9qqUwxbBfShsNreBaSsKvFcHzMxfBC3Ki"`
	AuctionHouse *string          `json:"auctionHouse,omitempty" example:"6hW2rVdPUD5qn1amEvN3K9zkvgsCA34LqCvTPcpamQHc"`
	Name         *string          `json:"name"`
	PriceMin     *float64         `json:"priceMin"`
	PriceMax     *float64         `json:"priceMax"`
}

type WebsocketParams struct {
	CollectionID string `json:"collection_id" example:"target collectionSlug"`
}

type AblyWSData struct {
	TokenAddress string `json:"token_address"`
	ActionType   string `json:"action_type"`
	Item         struct {
		TokenAddress      string      `json:"token_address"`
		ProjectID         string      `json:"project_id"`
		FirstCreator      string      `json:"first_creator"`
		MccID             *string     `json:"mcc_id"`
		ProjectSlug       string      `json:"project_slug"`
		CandyMachineID    *string     `json:"candy_machine_id"`
		Name              string      `json:"name"`
		RankEst           int         `json:"rank_est"`
		MetaDataImg       string      `json:"meta_data_img"`
		MetaDataURI       string      `json:"meta_data_uri"`
		Attributes        interface{} `json:"attributes"`
		ProjectName       string      `json:"project_name"`
		ProjectImage      string      `json:"project_image"`
		IsProjectVerified bool        `json:"is_project_verified"`
		CreatorRoyalty    int         `json:"creator_royalty"`
		NftStandard       string      `json:"nft_standard"`
		MarketPlaceState  struct {
			Price *float64 `json:"price"`
		} `json:"market_place_state"`
	} `json:"item"`
}

type LiveDataResponse struct {
	MintAddress string   `json:"mintAddress"`
	ActionType  string   `json:"actionType"`
	Price       *float64 `json:"price"`
}
