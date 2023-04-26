package types

type StatRes struct {
	MarketCap uint64 `json:"marketCap"`
	Volume    uint64 `json:"volume"`
	Volume1D  uint64 `json:"volume1d"`
}

type SearchParams struct {
	Mode    string `json:"mode"`
	Keyword string `json:"keyword"`
	Offset  int    `json:"offset"`
	Limit   int    `json:"limit"`
}

type FoundObj struct {
	Twitter     *string `json:"twitter,omitempty"`
	Address     *string `json:"address,omitempty"`
	ProjectSlug *string `json:"project_slug,omitempty"`
	ProjectID   *string `json:"project_id,omitempty"`
	IsVerified  *bool   `json:"is_verified,omitempty"`
	ProjectName *string `json:"display_name,omitempty"`
	ImgURL      *string `json:"img_url,omitempty"`
	Volume      *int    `json:"volume,omitempty"`
}

type ObjInfo struct {
	Twitter     *string `json:"twitter,omitempty"`
	Address     *string `json:"address,omitempty"`
	ProjectSlug *string `json:"slug,omitempty"`
	IsVerified  *bool   `json:"isVerified,omitempty"`
	ProjectName *string `json:"name,omitempty"`
	ImgURL      *string `json:"imgURL,omitempty"`
	Volume1D    *string `json:"volume1d,omitempty"`
}

type SearchRes struct {
	Results     []ObjInfo `json:"results"`
	HasNextPage bool      `json:"hasNextPage"`
}
