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
	Statistics                *Statistics `json:"statistics,omitempty"`
	Attributes                []Attribute `json:"attributes"`
}

type Trend struct {
	Floor1D         string     `json:"floor1d"`
	Floor7D         string     `json:"floor7d"`
	Floor30D        string     `json:"floor30d"`
	Volume1D        string     `json:"volume1d"`
	Volume7D        string     `json:"volume7d"`
	Volume30D       string     `json:"volume30d"`
	Listed1D        string     `json:"listed1d"`
	Listed7D        string     `json:"listed7d"`
	Listed30D       string     `json:"listed30d"`
	ChangeFloor1D   int32      `json:"changeFloor1d"`
	ChangeFloor7D   int32      `json:"changeFloor7d"`
	ChangeFloor30D  int32      `json:"changeFloor30d"`
	ChangeVolume1D  int32      `json:"changeVolume1d"`
	ChangeVolume7D  int32      `json:"changeVolume7d"`
	ChangeVolume30D int32      `json:"changeVolume30d"`
	ChangeListed1D  int32      `json:"changeListed1d"`
	ChangeListed7D  int32      `json:"changeListed7d"`
	ChangeListed30D int32      `json:"changeListed30d"`
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
	Order  string `json:"order" example:"ASC|DESC"`
	Limit  int    `json:"limit" example:"10"`
	Offset int    `json:"offset" example:"0"`
}

type TimeSeriesParams struct {
	Address     string `json:"collection" example:"8xBMPGAj5NzAwRmdfEcksDcZyexr87AAmD6LWwKG7Dqq"`
	FromTime    int    `json:"start_timestamp" example:"1671128400"`
	ToTime      int    `json:"end_timestamp" example:"1679410436"`
	Granularity string `json:"time_granularity" example:"PER_HOUR|PER_DAY"`
	Limit       int    `json:"limit" example:"10"`
	Offset      int    `json:"offset" example:"0"`
}

type NFTParams struct {
	Address    string      `json:"collection" example:"8xBMPGAj5NzAwRmdfEcksDcZyexr87AAmD6LWwKG7Dqq"`
	Attributes []Attribute `json:"attributes,omitempty"`
	SortBy     string      `json:"sort_by" example:"lowest_listing_block_timestamp"`
	Order      string      `json:"order" example:"ASC|DESC"`
	Limit      int         `json:"limit" example:"10"`
	Offset     int         `json:"offset" example:"0"`
}
