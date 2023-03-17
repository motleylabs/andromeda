package types

type Collection struct {
	ID                        string `json:"id"`
	Name                      string `json:"name"`
	VerifiedCollectionAddress string `json:"verifiedCollectionAddress"`
	Description               string `json:"description"`
	Image                     string `json:"image"`
	Symbol                    string `json:"symbol"`
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

type Statistics struct {
	Floor1D   string  `json:"floor1d"`
	Volume30D string  `json:"volume30d"`
	Listed1D  string  `json:"listed1d"`
	Holders   int64   `json:"holders"`
	MarketCap float64 `json:"marketCap"`
	Supply    int64   `json:"supply"`
}

type TimeSeries struct {
	TimeStamp string `json:"timestamp"`
	Value     string `json:"value"`
}

type TrendParams struct {
	Period  string
	SortBy  string
	OrderBy string
	Limit   int32
	Offset  int32
}

type TimeSeriesParams struct {
	Collection  string
	Mode        string
	FromTime    int64
	ToTime      int64
	Granularity string
}
