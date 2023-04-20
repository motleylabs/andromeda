package collection

import (
	"andromeda/pkg/service/hyperspace/common"
)

type ProjectStatHistEntry struct {
	Timestamp         int      `json:"timestamp"`
	Volume            int      `json:"volume"`
	FloorPrice        *float64 `json:"floor_price"`
	NumOfSales        *int     `json:"num_of_sales"`
	NumOfTokenListed  *int     `json:"num_of_token_listed"`
	NumOfTokenHolders *int     `json:"num_of_token_holders"`
}

type ProjectStatHistRes struct {
	PaginationInfo common.PaginationInfo  `json:"pagination_info"`
	HistEntries    []ProjectStatHistEntry `json:"project_stat_hist_entries"`
}

type ProjectAttributeStats struct {
	ProjectID   string              `json:"project_id"`
	Name        string              `json:"name"`
	Type        string              `json:"type"`
	Counts      map[string]int      `json:"counts"`
	FloorPrices map[string]*float64 `json:"floor_prices"`
	NumListed   map[string]int      `json:"num_listed"`
}

type ProjectAttributeRes struct {
	PaginationInfo         common.PaginationInfo   `json:"pagination_info"`
	ProjectAttributesStats []ProjectAttributeStats `json:"project_attribute_stats"`
}
