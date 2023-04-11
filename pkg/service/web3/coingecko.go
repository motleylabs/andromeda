package web3

import (
	"andromeda/pkg/request"
	"encoding/json"
)

type USDCost struct {
	USD float64 `json:"usd"`
}

type SOLPrice struct {
	Solana USDCost `json:"solana"`
}

func GetSOLprice() (float64, error) {
	price := float64(0)

	URL := "https://api.coingecko.com/api/v3/simple/price?ids=solana&vs_currencies=usd"
	res, err := request.ProcessGet(URL)
	if err != nil {
		return price, err
	}

	var priceRes SOLPrice
	if err := json.Unmarshal(res, &priceRes); err != nil {
		return price, err
	}

	return priceRes.Solana.USD, nil
}
