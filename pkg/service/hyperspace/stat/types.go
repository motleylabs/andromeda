package stat

import "andromeda/pkg/service/entrance/types"

type AlgoliaResult struct {
	Hits []types.FoundObj `json:"hits"`
}

type AlgoliaRes struct {
	Results []AlgoliaResult `json:"results"`
}

type ErrorRes struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
