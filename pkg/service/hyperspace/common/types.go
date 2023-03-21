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
	ProjectIDs *[]ProjectIDItem `json:"project_ids,omitempty"`
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
