package types

type Statistics struct {
	Floor1D   string  `json:"floor1d"`
	Volume30D string  `json:"volume30d"`
	Listed1D  string  `json:"listed1d"`
	Holders   int64   `json:"holders"`
	MarketCap float64 `json:"marketCap"`
	Supply    int64   `json:"supply"`
}

type Attribute struct {
	Name   string   `json:"name" example:"Background"`
	Type   string   `json:"type" example:"CATEGORY"`
	Values []string `json:"values" example:"Vivid-tangerine"`
}

type Collection struct {
	ID                        string      `json:"id"`
	Name                      string      `json:"name"`
	VerifiedCollectionAddress string      `json:"verifiedCollectionAddress"`
	Description               string      `json:"description"`
	Image                     string      `json:"image"`
	Symbol                    string      `json:"symbol"`
	Slug                      string      `json:"slug"`
	Statistics                *Statistics `json:"statistics,omitempty"`
	Attributes                []Attribute `json:"attributes"`
}

type Trend struct {
	Floor1D         string     `json:"floor1d"`
	Floor7D         string     `json:"floor7d,omitempty"`
	Floor30D        string     `json:"floor30d,omitempty"`
	Volume1D        string     `json:"volume1d"`
	Volume7D        string     `json:"volume7d"`
	Volume30D       string     `json:"volume30d"`
	Listed1D        string     `json:"listed1d"`
	Listed7D        string     `json:"listed7d,omitempty"`
	Listed30D       string     `json:"listed30d,omitempty"`
	ChangeFloor1D   *float64   `json:"changeFloor1d"`
	ChangeFloor7D   *float64   `json:"changeFloor7d,omitempty"`
	ChangeFloor30D  *float64   `json:"changeFloor30d,omitempty"`
	ChangeVolume1D  *float64   `json:"changeVolume1d"`
	ChangeVolume7D  *float64   `json:"changeVolume7d,omitempty"`
	ChangeVolume30D *float64   `json:"changeVolume30d,omitempty"`
	ChangeListed1D  *float64   `json:"changeListed1d,omitempty"`
	ChangeListed7D  *float64   `json:"changeListed7d,omitempty"`
	ChangeListed30D *float64   `json:"changeListed30d,omitempty"`
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
	Granularity string `json:"timeGranularity" example:"PER_HOUR|PER_DAY"`
	Limit       int    `json:"limit" example:"10"`
	Offset      int    `json:"offset" example:"0"`
}

type NFTParams struct {
	Address      string      `json:"collection" example:"8xBMPGAj5NzAwRmdfEcksDcZyexr87AAmD6LWwKG7Dqq"`
	Attributes   []Attribute `json:"attributes,omitempty"`
	SortBy       string      `json:"sortBy" example:"lowest_listing_block_timestamp"`
	Order        string      `json:"order" example:"asc|desc"`
	Limit        int         `json:"limit" example:"10"`
	Offset       int         `json:"offset" example:"0"`
	ListingOnly  bool        `json:"listingOnly,omitempty"`
	Program      *string     `json:"program,omitempty" example:"RwDDvPp7ta9qqUwxbBfShsNreBaSsKvFcHzMxfBC3Ki"`
	AuctionHouse *string     `json:"auctionHouse,omitempty" example:"6hW2rVdPUD5qn1amEvN3K9zkvgsCA34LqCvTPcpamQHc"`
}
