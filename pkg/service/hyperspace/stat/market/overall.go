package market

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
)

type StatRes struct {
	MarketCap  uint64 `json:"market_cap"`
	Volume     uint64 `json:"volume"`
	Volume1Day uint64 `json:"volume_1day"`
}

func GetOverall() (*types.StatRes, error) {
	res, err := request.ProcessGet(fmt.Sprintf("%s/get-overall-project-stats", common.ENDPOINT))
	if err != nil {
		return nil, err
	}

	var statRes StatRes
	if err := json.Unmarshal(res, &statRes); err != nil {
		return nil, err
	}

	return &types.StatRes{
		MarketCap: statRes.MarketCap,
		Volume:    statRes.Volume,
		Volume1D:  statRes.Volume1Day,
	}, nil
}
