package collection

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

func GetDetail(address string) (*types.Collection, error) {
	projectIDs := []string{
		address,
	}
	excludeProjectAttr := true
	projectStatParams := common.StatParams{
		Conditions: &common.Conditions{
			ProjectIDs:               &projectIDs,
			ExcludeProjectAttributes: &excludeProjectAttr,
		},
	}

	payload, err := json.Marshal(projectStatParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-project-stats", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var projectStats ProjectStatRes
	if err := json.Unmarshal(res, &projectStats); err != nil {
		return nil, err
	}

	if len(projectStats.ProjectStats) == 0 {
		return nil, fmt.Errorf("invalid project id")
	}

	fmt.Println(projectStats.ProjectStats[0].ProjectID)

	return ConvertProjectStat(&projectStats.ProjectStats[0]), nil
}

func ConvertProjectStat(projectStat *ProjectStat) *types.Collection {
	holders := int64(0)
	if projectStat.TokenHolders != nil {
		holders = int64(*projectStat.TokenHolders)
	}

	marketCap := float64(0)
	if projectStat.MarketCap != nil {
		marketCap = *projectStat.MarketCap
	}

	stat := types.Statistics{
		Volume30D: common.GetFromIntPointer(projectStat.Volume1M),
		Listed1D:  common.GetFromIntPointer(projectStat.Listed1Day),
		Floor1D:   common.GetLamportsFromPointer(projectStat.FloorPrice1Day),
		Holders:   holders,
		MarketCap: marketCap,
		Supply:    projectStat.Project.Supply,
	}

	var collection types.Collection
	collection.ID = projectStat.ProjectID
	collection.Description = projectStat.Project.Description
	collection.Image = projectStat.Project.ImgURL
	collection.Name = projectStat.Project.DisplayName
	collection.Statistics = &stat

	return &collection
}
