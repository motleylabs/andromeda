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
