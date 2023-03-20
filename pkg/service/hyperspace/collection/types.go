package collection

import (
	"andromeda/pkg/service/hyperspace/common"
)

type ProjectStatParam struct {
	OrderBy        common.OrderConfig       `json:"order_by"`
	PaginationInfo *common.PaginationConfig `json:"pagination_info,omitempty"`
}

type Project struct {
	ProjectID   string `json:"project_id"`
	IsVerified  bool   `json:"is_verified"`
	DisplayName string `json:"display_name"`
	ImgURL      string `json:"img_url"`
	Description string `json:"description"`
	Supply      int    `json:"supply"`
}

type ProjectStat struct {
	ProjectID            string   `json:"project_id"`
	MarketCap            *int     `json:"market_cap"`
	Volume1Day           *int     `json:"volume_1day"`
	Volume7Day           *int     `json:"volume_7day"`
	Volume1M             *int     `json:"volume_1m"`
	Volume1DayChange     *int     `json:"volume_1day_change,omitempty"`
	Volume7DayChange     *int     `json:"volume_1day_change,omitempty"`
	Volume1MChange       *int     `json:"volume_1m_change,omitempty"`
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
