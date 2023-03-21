package collection

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
)

type Project struct {
	ProjectID   string             `json:"project_id"`
	IsVerified  bool               `json:"is_verified"`
	DisplayName string             `json:"display_name"`
	ImgURL      string             `json:"img_url"`
	Description string             `json:"description"`
	Supply      int64              `json:"supply"`
	Attributes  *[]types.Attribute `json:"project_attributes"`
}

type ProjectStat struct {
	ProjectID            string   `json:"project_id"`
	MarketCap            *float64 `json:"market_cap"`
	Volume1Day           *int     `json:"volume_1day"`
	Volume7Day           *int     `json:"volume_7day"`
	Volume1M             *int     `json:"volume_1m"`
	Volume1DayChange     *float64 `json:"volume_1day_change,omitempty"`
	Volume7DayChange     *float64 `json:"volume_7day_change,omitempty"`
	Volume1MChange       *float64 `json:"volume_1m_change,omitempty"`
	FloorPrice1Day       *float64 `json:"floor_price"`
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
	PaginationInfo common.PaginationInfo `json:"pagination_info"`
	ProjectStats   []ProjectStat         `json:"project_stats"`
}

type LastSaleMPA struct {
	UserAddress    string  `json:"user_address"`
	Price          float64 `json:"price"`
	Signature      string  `json:"signature"`
	BlockTimestamp int64   `json:"block_timestamp"`
}

type MarketPlaceSnapshot struct {
	TokenAddress   string                 `json:"token_address"`
	Name           string                 `json:"name"`
	RankeEst       int                    `json:"rank_est"`
	MoonRank       int                    `json:"moonrank"`
	MetadataURI    string                 `json:"meta_data_uri"`
	MetadataImg    string                 `json:"meta_data_img"`
	Attributes     map[string]interface{} `json:"attributes"`
	LastSaleMPA    *LastSaleMPA           `json:"last_sale_mpa"`
	CreatorRoyalty int                    `json:"creator_royalty"`
	NFTStandard    string                 `json:"nft_standard"`
	Owner          string                 `json:"owner"`
}

type ProjectNFTsRes struct {
	PaginationInfo       common.PaginationInfo `json:"pagination_info"`
	MarketPlaceSnapshots []MarketPlaceSnapshot `json:"market_place_snapshots"`
}
