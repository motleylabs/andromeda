package common

type OrderConfig struct {
	FieldName string `json:"field_name"`
	SortOrder string `json:"sort_order"`
}

type PaginationConfig struct {
	PageNumber *int `json:"page_number,omitempty"`
	PageSize   *int `json:"page_size,omitempty"`
}

type PaginationInfo struct {
	CurrentPageNumber int  `json:"current_page_number"`
	CurrentPageSize   int  `json:"current_page_size"`
	HasNextPage       bool `json:"has_next_page"`
	TotalPageNumber   int  `json:"total_page_number"`
}
