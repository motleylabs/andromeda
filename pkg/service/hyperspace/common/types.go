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
	ProjectID  string                  `json:"project_id"`
	Attributes *[]types.AttributeInput `json:"attributes,omitempty"`
}

type Conditions struct {
	ProjectIDs               *[]string `json:"project_ids,omitempty"`
	ExcludeProjectAttributes *bool     `json:"exclude_project_attributes,omitempty"`
	IsVerified               *bool     `json:"is_verified,omitempty"`
	StartTimestamp           *int      `json:"start_timestamp,omitempty"`
	EndTimestamp             *int      `json:"end_timestamp,omitempty"`
	TimeGranularity          *string   `json:"time_granularity,omitempty"`
}

type MarketPlaceProgram struct {
	MarketPlaceProgramID  string `json:"marketplace_program_id"`
	MarketPlaceInstanceID string `json:"marketplace_instance_id"`
}

type MarketPlaceProgramCondition struct {
	MarketPlacePrograms []MarketPlaceProgram `json:"marketplace_programs"`
}

type Condition struct {
	ProjectID                   *string                      `json:"project_id,omitempty"`
	ProjectIDs                  *[]ProjectIDItem             `json:"project_ids,omitempty"`
	TokenAddresses              *[]string                    `json:"token_addresses,omitempty"`
	ActionType                  *string                      `json:"action_type,omitempty"`
	MarketPlaceProgramCondition *MarketPlaceProgramCondition `json:"marketplace_program_condition,omitempty"`
	ListingType                 *string                      `json:"listing_type,omitempty"`
}

type ActivityCondition struct {
	Projects      *[]ProjectIDItem `json:"projects,omitempty"`
	SellerAddress *string          `json:"seller_address,omitempty"`
	BuyerAddress  *string          `json:"buyer_address,omitempty"`
	ByMPATypes    []string         `json:"by_mpa_types"`
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

type Metadata struct {
	Seller     *string `json:"seller_address,omitempty"`
	TradeState *string `json:"seller_trade_state,omitempty"`
}

type MPAInfo struct {
	UserAddress           string   `json:"user_address"`
	Price                 float64  `json:"price"`
	Signature             string   `json:"signature"`
	BlockTimestamp        int64    `json:"block_timestamp"`
	MarketPlaceProgramID  string   `json:"marketplace_program_id"`
	MarketPlaceInstanceID string   `json:"marketplace_instance_id"`
	Metadata              Metadata `json:"metadata"`
}

type MarketPlaceState struct {
	BlockTimestamp        int      `json:"block_timestamp"`
	Signature             string   `json:"signature"`
	SellerAddress         *string  `json:"seller_address"`
	BuyerAddress          *string  `json:"buyer_address"`
	Type                  string   `json:"type"`
	MarketPlaceProgramID  string   `json:"marketplace_program_id"`
	MarketPlaceInstanceID string   `json:"marketplace_instance_id"`
	Price                 *float64 `json:"price"`
	CreatedAt             string   `json:"created_at"`
}

type MarketPlaceSnapshot struct {
	ProjectID        string                 `json:"project_id"`
	TokenAddress     string                 `json:"token_address"`
	Name             string                 `json:"name"`
	RankeEst         int                    `json:"rank_est"`
	MoonRank         int                    `json:"moonrank"`
	MetadataURI      string                 `json:"meta_data_uri"`
	MetadataImg      string                 `json:"meta_data_img"`
	Attributes       map[string]interface{} `json:"attributes"`
	LastSaleMPA      *MPAInfo               `json:"last_sale_mpa"`
	LowestListingMPA *MPAInfo               `json:"lowest_listing_mpa"`
	HighestBidMPA    *MPAInfo               `json:"highest_bid_mpa"`
	CreatorRoyalty   int                    `json:"creator_royalty"`
	NFTStandard      string                 `json:"nft_standard"`
	Owner            *string                `json:"owner"`
	MarketPlaceState *MarketPlaceState      `json:"market_place_state"`
	ProjectSlug      string                 `json:"project_slug"`
}

type Attribute struct {
	Name   string         `json:"name"`
	Type   string         `json:"type"`
	Values []string       `json:"values"`
	Counts map[string]int `json:"counts"`
}

type Project struct {
	ProjectID                 string       `json:"project_id"`
	IsVerified                bool         `json:"is_verified"`
	DisplayName               string       `json:"display_name"`
	ImgURL                    string       `json:"img_url"`
	Description               string       `json:"description"`
	Supply                    int64        `json:"supply"`
	Attributes                *[]Attribute `json:"project_attributes"`
	ProjectSlug               string       `json:"project_slug"`
	VerifiedCollectionAddress string       `json:"mcc_id"`
}

type ProjectSnapshotsRes struct {
	PaginationInfo       PaginationInfo        `json:"pagination_info"`
	MarketPlaceSnapshots []MarketPlaceSnapshot `json:"market_place_snapshots"`
}

type ProjectStat struct {
	ProjectID            string   `json:"project_id"`
	MarketCap            *float64 `json:"market_cap"`
	Volume1Hr            *int     `json:"volume_1hr"`
	Volume1Day           *int     `json:"volume_1day"`
	Volume7Day           *int     `json:"volume_7day"`
	Volume1HrChange      *float64 `json:"volume_1r_change,omitempty"`
	Volume1DayChange     *float64 `json:"volume_1day_change,omitempty"`
	Volume7DayChange     *float64 `json:"volume_7day_change,omitempty"`
	FloorPrice1Day       *float64 `json:"feeless_floor_price"`
	FloorPrice7Day       *float64 `json:"floor_price_7day,omitempty"`
	FloorPrice1M         *float64 `json:"floor_price_1m,omitempty"`
	FloorPrice1DayChange *float64 `json:"floor_price_1day_change,omitempty"`
	FloorPrice7DayChange *float64 `json:"floor_price_7day_change,omitempty"`
	FloorPrice1MChange   *float64 `json:"floor_price_1m_change,omitempty"`
	Listed1Day           *int     `json:"num_of_token_listed"`
	TokenHolders         *int     `json:"num_of_token_holders"`
	Project              Project  `json:"project"`
}

type ProjectStatRes struct {
	PaginationInfo PaginationInfo `json:"pagination_info"`
	ProjectStats   []ProjectStat  `json:"project_stats"`
}
