package types

type ReportRes struct {
	Volume   float64 `json:"volume"`
	TPS      uint64  `json:"tps"`
	SOLPrice float64 `json:"solPrice"`
}
