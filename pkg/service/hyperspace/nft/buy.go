package nft

import (
	"andromeda/pkg/request"
	"andromeda/pkg/service/entrance/types"
	"andromeda/pkg/service/hyperspace/common"
	"encoding/json"
	"fmt"
	"strconv"
)

type HSBuyParams struct {
	Buyer        string  `json:"buyer_address"`
	Price        float64 `json:"price"`
	TokenAddress string  `json:"token_address"`
	BuyerBroker  string  `json:"buyer_broker"`
}

type HSBuyError struct {
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
}

type HSBuyResponse struct {
	Data      []byte      `json:"data"`
	StdBuffer []byte      `json:"stdBuffer"`
	Error     *HSBuyError `json:"error"`
}

func GetBuyNowTx(params *types.BuyParams) (*types.BuyRes, error) {
	price, err := strconv.ParseInt(params.Price, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("price param is not valid")
	}

	hsParams := HSBuyParams{
		Buyer:        params.Buyer,
		BuyerBroker:  params.BuyerBroker,
		Price:        float64(price) / common.LAMPORTS_PER_SOL,
		TokenAddress: params.Mint,
	}
	payload, err := json.Marshal(hsParams)
	if err != nil {
		return nil, err
	}

	res, err := request.ProcessPost(fmt.Sprintf("%s/create-buy-tx", common.ENDPOINT), payload)
	if err != nil {
		return nil, err
	}

	var buyRes HSBuyResponse
	if err := json.Unmarshal(res, &buyRes); err != nil {
		return nil, err
	}

	if buyRes.Error != nil {
		return nil, fmt.Errorf(buyRes.Error.Message)
	}

	if len(buyRes.StdBuffer) > 0 {
		return &types.BuyRes{
			Buffer: buyRes.StdBuffer,
		}, nil
	}

	return nil, fmt.Errorf("no data returned")
}
