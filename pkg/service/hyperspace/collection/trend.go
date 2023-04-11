package collection

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"andromeda/pkg/service/web3"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-contrib/cache/persistence"
)

func fetchSOLPrice(store *persistence.InMemoryStore) {
	solPrice, err := web3.GetSOLprice()
	if err == nil {
		store.Set("andromeda-sol-price", solPrice, -1)
	}
}

func getSOLPrice(store *persistence.InMemoryStore) (float64, error) {
	retries := 0
	for {
		if retries == 5 {
			return 0, fmt.Errorf("failed to get SOL price")
		}

		var priceStr interface{}
		if err := store.Get("andromeda-sol-price", &priceStr); err != nil {
			time.Sleep(200 * time.Millisecond)
			retries += 1
			continue
		}

		price, ok := priceStr.(float64)
		if !ok {
			return 0, fmt.Errorf("SOL price is invalid")
		}

		return price, nil
	}
}

func GetTrends(params *types.TrendParams, store *persistence.InMemoryStore) (*types.TrendRes, error) {
	go fetchSOLPrice(store)

	if params == nil {
		return nil, fmt.Errorf("no trend params")
	}
	projectStatParams := getProjectStatParams(params)
	payload, err := json.Marshal(projectStatParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-project-stats", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var projectStats common.ProjectStatRes
	if err := json.Unmarshal(res, &projectStats); err != nil {
		return nil, err
	}

	solPrice, err := getSOLPrice(store)
	if err != nil {
		return nil, err
	}

	trendRes := types.TrendRes{
		HasNextPage: projectStats.PaginationInfo.HasNextPage,
		Trends:      convertStatistics(projectStats.ProjectStats, solPrice),
	}

	// cache project id and slugs
	if len(trendRes.Trends) > 0 {
		for index := range trendRes.Trends {
			tempTrend := trendRes.Trends[index]
			curID := tempTrend.Collection.ID
			curSlug := tempTrend.Collection.Slug

			if err := store.Set(curSlug, curID, -1); err != nil {
				return nil, err
			}
		}
	}

	return &trendRes, nil
}

func convertStatistics(stats []common.ProjectStat, solPrice float64) []types.Trend {
	trends := make([]types.Trend, len(stats))

	for index := range stats {
		// volume
		trends[index].Volume1D = common.GetLamportsFromUSDIntPointer(stats[index].Volume1Day, solPrice)
		trends[index].Volume7D = common.GetLamportsFromUSDIntPointer(stats[index].Volume7Day, solPrice)
		trends[index].Volume30D = common.GetLamportsFromUSDIntPointer(stats[index].Volume1M, solPrice)
		trends[index].ChangeVolume1D = common.GetPercentFromPointer(stats[index].Volume1DayChange)

		// floor price
		trends[index].Floor1D = common.GetLamportsFromPointer(stats[index].FloorPrice1Day)
		trends[index].ChangeFloor1D = common.GetPercentFromPointer(stats[index].FloorPrice1DayChange)

		// listing
		trends[index].Listed1D = common.GetFromIntPointer(stats[index].Listed1Day)

		// collection
		trends[index].Collection = types.Collection{
			ID:          stats[index].ProjectID,
			Name:        stats[index].Project.DisplayName,
			Description: stats[index].Project.Description,
			Image:       stats[index].Project.ImgURL,
			Slug:        stats[index].Project.ProjectSlug,
		}
	}

	return trends
}

func getProjectStatParams(input *types.TrendParams) *common.StatParams {
	excludeProjectAttr := true
	return &common.StatParams{
		Conditions: &common.Conditions{
			ExcludeProjectAttributes: &excludeProjectAttr,
		},
		OrderBy:        getOrderField(input),
		PaginationInfo: getPaginationInfo(input),
	}
}

func getOrderField(input *types.TrendParams) *common.OrderConfig {
	orderFieldName := "floor_price"

	switch input.SortBy {
	case "volume":
		orderFieldName = "volume"
	case "listed":
		orderFieldName = "listed"
	}

	periodSuffix := "1day"
	switch input.Period {
	case "7d":
		periodSuffix = "7day"
	case "1m":
		periodSuffix = "1m"
	}

	return &common.OrderConfig{
		FieldName: fmt.Sprintf("%s_%s", orderFieldName, periodSuffix),
		SortOrder: strings.ToUpper(input.Order),
	}
}

func getPaginationInfo(input *types.TrendParams) *common.PaginationConfig {
	pageNumber := input.Offset/input.Limit + 1
	pageSize := input.Limit

	paginationInfo := common.PaginationConfig{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}

	return &paginationInfo
}
