package collection

import (
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"fmt"
)

func GetDetail(address string) (*types.Collection, error) {
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

	return common.ConvertProjectStat(&projectStats.ProjectStats[0]), nil
}
