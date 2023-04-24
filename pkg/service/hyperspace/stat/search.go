package stat

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-contrib/cache/persistence"
)

func GetHits(params *types.SearchParams, store *persistence.InMemoryStore) (*types.SearchRes, error) {
	go common.FetchSOLPrice(store)

	client := search.NewClient(common.ALGOLIA_APP_ID, common.ALGOLIA_API_KEY)

	indexName := "volume_1d_desc"
	if params.Mode == "user" {
		indexName = "users_index"
	}
	index := client.InitIndex(indexName)

	res, err := index.Search(params.Keyword, opt.Offset(params.Offset), opt.Length(params.Limit))
	if err != nil {
		return nil, err
	}

	solPrice, err := common.GetSOLPrice(store)
	if err != nil {
		return nil, err
	}

	var records []types.FoundObj
	if err := res.UnmarshalHits(&records); err != nil {
		spew.Dump(records)
		return nil, err
	}

	searchRes := types.SearchRes{
		Results:     convertRecords(records, solPrice, store),
		HasNextPage: len(records) == params.Limit,
	}

	return &searchRes, nil
}

func convertRecords(records []types.FoundObj, solPrice float64, store *persistence.InMemoryStore) []types.ObjInfo {
	objects := make([]types.ObjInfo, len(records))

	for index := range records {
		var volume1D *string
		if records[index].Volume != nil {
			volume := common.GetLamportsFromUSDIntPointer(records[index].Volume, solPrice)
			volume1D = &volume
		}

		curSlug := records[index].ProjectSlug
		curID := records[index].ProjectID

		if curSlug == nil || curID == nil {
			continue
		}

		if err := store.Set(*curSlug, *curID, -1); err != nil {
			continue
		}

		objects[index] = types.ObjInfo{
			Twitter:     records[index].Twitter,
			Address:     records[index].Address,
			ProjectSlug: records[index].ProjectSlug,
			IsVerified:  records[index].IsVerified,
			ProjectName: records[index].ProjectName,
			ImgURL:      records[index].ImgURL,
			Volume1D:    volume1D,
		}

	}

	return objects
}
