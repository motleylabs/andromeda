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
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Values []string `json:"values"`
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
	HasNextPage bool    `json:"has_next_page"`
}

type TimeSeries struct {
	Timestamp  string `json:"timestamp"`
	FloorPrice string `json:"floorPrice"`
	Listed     int    `json:"listed"`
}

type TimeSeriesRes struct {
	Series      []TimeSeries `json:"series"`
	HasNextPage bool         `json:"has_next_page"`
}

type TrendParams struct {
	Period string `json:"period"`
	SortBy string `json:"sort_by"`
	Order  string `json:"order"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type TimeSeriesParams struct {
	Address     string `json:"collection"`
	FromTime    int    `json:"start_timestamp"`
	ToTime      int    `json:"end_timestamp"`
	Granularity string `json:"time_granularity"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}

type NFTParams struct {
	Address    string      `json:"collection"`
	Attributes []Attribute `json:"attributes,omitempty"`
	SortBy     string      `json:"sort_by"`
	Order      string      `json:"order"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
}
