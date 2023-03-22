package common

import "andromeda/pkg/service/entrance/types"

type OrderConfig struct {
	FieldName string `json:"field_name"`
	SortOrder string `json:"sort_order"`
}

type PaginationConfig struct {
	PageNumber *int `json:"page_number,omitempty"`
	PageSize   *int `json:"page_size,omitempty"`
}

type ProjectIDItem struct {
	ProjectID  string             `json:"project_id"`
	Attributes *[]types.Attribute `json:"attributes,omitempty"`
}

type Conditions struct {
	ProjectIDs               *[]string `json:"project_ids,omitempty"`
	ExcludeProjectAttributes *bool     `json:"exclude_project_attributes,omitempty"`
	IsVerified               *bool     `json:"is_verified,omitempty"`
	StartTimestamp           *int      `json:"start_timestamp,omitempty"`
	EndTimestamp             *int      `json:"end_timestamp,omitempty"`
	TimeGranularity          *string   `json:"time_granularity,omitempty"`
}

type Condition struct {
	ProjectIDs     *[]ProjectIDItem `json:"project_ids,omitempty"`
	TokenAddresses *[]string        `json:"token_addresses,omitempty"`
}

type ActivityCondition struct {
	Projects   []ProjectIDItem `json:"projects"`
	ByMPATypes []string        `json:"by_mpa_types"`
}

type PaginationInfo struct {
	CurrentPageNumber int  `json:"current_page_number"`
	CurrentPageSize   int  `json:"current_page_size"`
	HasNextPage       bool `json:"has_next_page"`
	TotalPageNumber   int  `json:"total_page_number"`
}

type StatParams struct {
	Conditions     *Conditions       `json:"conditions,omitempty"`
	Condition      *Condition        `json:"condition,omitempty"`
	OrderBy        *OrderConfig      `json:"order_by,omitempty"`
	PaginationInfo *PaginationConfig `json:"pagination_info,omitempty"`
}

type ActivityParams struct {
	ActivityCondition ActivityCondition `json:"condition"`
	PaginationInfo    *PaginationConfig `json:"pagination_info,omitempty"`
}

type LastSaleMPA struct {
	UserAddress    string  `json:"user_address"`
	Price          float64 `json:"price"`
	Signature      string  `json:"signature"`
	BlockTimestamp int64   `json:"block_timestamp"`
}

type MarketPlaceState struct {
	BlockTimestamp       int      `json:"block_timestamp"`
	Signature            string   `json:"signature"`
	SellerAddress        *string  `json:"seller_address"`
	BuyerAddress         string   `json:"buyer_address"`
	Type                 string   `json:"type"`
	MarketPlaceProgramID string   `json:"marketplace_program_id"`
	Price                *float64 `json:"price"`
	CreatedAt            string   `json:"created_at"`
}

type MarketPlaceSnapshot struct {
	TokenAddress     string                 `json:"token_address"`
	Name             string                 `json:"name"`
	RankeEst         int                    `json:"rank_est"`
	MoonRank         int                    `json:"moonrank"`
	MetadataURI      string                 `json:"meta_data_uri"`
	MetadataImg      string                 `json:"meta_data_img"`
	Attributes       map[string]interface{} `json:"attributes"`
	LastSaleMPA      *LastSaleMPA           `json:"last_sale_mpa"`
	CreatorRoyalty   int                    `json:"creator_royalty"`
	NFTStandard      string                 `json:"nft_standard"`
	Owner            string                 `json:"owner"`
	MarketPlaceState *MarketPlaceState      `json:"market_place_state"`
}

type ProjectSnapshotsRes struct {
	PaginationInfo       PaginationInfo        `json:"pagination_info"`
	MarketPlaceSnapshots []MarketPlaceSnapshot `json:"market_place_snapshots"`
}
