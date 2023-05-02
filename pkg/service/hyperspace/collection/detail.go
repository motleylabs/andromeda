package collection

import (
	"andromeda/internal/api/state"
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/gin-contrib/cache/persistence"
)

func GetDetail(address string, store *persistence.InMemoryStore) (*types.Collection, error) {
	projectIDs := []string{
		address,
	}

	wg := sync.WaitGroup{}
	wg.Add(2)

	// project detail data
	var projectStats *common.ProjectStatRes
	var projectResError error

	go func() {
		defer wg.Done()
		projectStats, projectResError = common.GetProjectsFromAddresses(projectIDs, true, 1, 10)
	}()

	// project attribute data
	var projectAttributeRes *ProjectAttributeRes
	var projectAttributeError error

	go func() {
		defer wg.Done()
		projectAttributeRes, projectAttributeError = getProjectAttributesFromAddress(address)
	}()

	wg.Wait()

	solPrice := state.GetSOLPrice()

	// error handler
	if projectResError != nil {
		return nil, projectResError
	}

	if projectAttributeError != nil {
		return nil, projectAttributeError
	}

	// get project stat data
	if len(projectStats.ProjectStats) == 0 {
		return nil, fmt.Errorf("invalid project id")
	}
	collection := common.ConvertProjectStat(&projectStats.ProjectStats[0], solPrice)
	store.Set(collection.Slug, collection.ID, -1)

	// get attribute data
	if len(projectAttributeRes.ProjectAttributesStats) > 0 {
		attributes := make([]types.AttributeOutput, len(projectAttributeRes.ProjectAttributesStats))

		for index := range projectAttributeRes.ProjectAttributesStats {
			curStats := projectAttributeRes.ProjectAttributesStats[index]
			curAttribute := types.AttributeOutput{
				Name:   curStats.Name,
				Type:   curStats.Type,
				Values: []types.AttributeStat{},
			}

			for count_k, count_v := range curStats.Counts {

				if count_v == 0 {
					continue
				}

				// get floor price
				var floorPrice *string
				floorSolPrice, ok := curStats.FloorPrices[count_k]
				if ok {
					floorPriceLamports := common.GetLamportsFromPointer(floorSolPrice)
					if floorPriceLamports != "0" {
						floorPrice = &floorPriceLamports
					}
				}

				// get listed count
				listed := curStats.NumListed[count_k]

				curAttribute.Values = append(curAttribute.Values, types.AttributeStat{
					Value:      count_k,
					Counts:     count_v,
					FloorPrice: floorPrice,
					Listed:     listed,
				})
			}

			sort.Slice(curAttribute.Values, func(i, j int) bool {
				if curAttribute.Values[i].Value == "None" {
					return true
				}

				if curAttribute.Values[j].Value == "None" {
					return false
				}

				return curAttribute.Values[i].Value < curAttribute.Values[j].Value
			})

			attributes[index] = curAttribute
		}

		sort.Slice(attributes, func(i, j int) bool {
			if attributes[i].Name == "Attributes Count" {
				return true
			}

			if attributes[j].Name == "Attributes Count" {
				return false
			}

			return attributes[i].Name < attributes[j].Name
		})

		collection.Attributes = attributes
	}

	return collection, nil
}

func getProjectAttributesFromAddress(address string) (*ProjectAttributeRes, error) {
	projectAttributeParam := common.StatParams{
		Condition: &common.Condition{
			ProjectID: &address,
		},
	}

	payload, err := json.Marshal(projectAttributeParam)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/get-project-attribute-stats", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var projectAttributeRes ProjectAttributeRes
	if err := json.Unmarshal(res, &projectAttributeRes); err != nil {
		return nil, err
	}

	return &projectAttributeRes, nil
}
