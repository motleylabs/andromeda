package types

type StatRes struct {
	MarketCap uint64 `json:"marketCap"`
	Volume    uint64 `json:"volume"`
	Volume1D  uint64 `json:"volume1d"`
}
