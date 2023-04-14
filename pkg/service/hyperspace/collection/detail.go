package collection

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"fmt"

	"github.com/gin-contrib/cache/persistence"
)

func GetDetail(address string, store *persistence.InMemoryStore) (*types.Collection, error) {
	go common.FetchSOLPrice(store)

	projectIDs := []string{
		address,
	}
	projectStats, err := common.GetProjectsFromAddresses(projectIDs, 1, 10)
	if err != nil {
		return nil, err
	}

	if len(projectStats.ProjectStats) == 0 {
		return nil, fmt.Errorf("invalid project id")
	}

	solPrice, err := common.GetSOLPrice(store)
	if err != nil {
		return nil, err
	}

	return common.ConvertProjectStat(&projectStats.ProjectStats[0], solPrice), nil
}
