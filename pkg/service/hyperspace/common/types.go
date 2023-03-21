package common

type OrderConfig struct {
	FieldName string `json:"field_name"`
	SortOrder string `json:"sort_order"`
}

type PaginationConfig struct {
	PageNumber *int `json:"page_number,omitempty"`
	PageSize   *int `json:"page_size,omitempty"`
}

type ProjectIDItem struct {
	ProjectID string `json:"project_id"`
}

type Condition struct {
	ProjectIDs               *[]ProjectIDItem `json:"project_ids,omitempty"`
	ExcludeProjectAttributes *bool            `json:"exclude_project_attributes,omitempty"`
	IsVerified               *bool            `json:"is_verified,omitempty"`
}

type PaginationInfo struct {
	CurrentPageNumber int  `json:"current_page_number"`
	CurrentPageSize   int  `json:"current_page_size"`
	HasNextPage       bool `json:"has_next_page"`
	TotalPageNumber   int  `json:"total_page_number"`
}

type StatParams struct {
	Condition      *Condition        `json:"condition,omitempty"`
	OrderBy        OrderConfig       `json:"order_by"`
	PaginationInfo *PaginationConfig `json:"pagination_info,omitempty"`
}
