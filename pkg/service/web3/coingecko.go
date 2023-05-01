package web3

import (
	"andromeda/pkg/request"
	"encoding/json"
	"sync"
)

type USDCost struct {
	USD float64 `json:"usd"`
}

type SOLPrice struct {
	Solana USDCost `json:"solana"`
}

var LastSOLPrice float64
var PriceLock sync.Mutex

func GetSOLPrice() (float64, error) {
	PriceLock.Lock()
	defer PriceLock.Unlock()
	price := float64(0)

	URL := "https://api.coingecko.com/api/v3/simple/price?ids=solana&vs_currencies=usd"
	res, err := request.ProcessGet(URL)
	if err != nil {
		if LastSOLPrice != 0 {
			return LastSOLPrice, nil
		}

		return price, err
	}

	var priceRes SOLPrice
	if err := json.Unmarshal(res, &priceRes); err != nil {
		if LastSOLPrice != 0 {
			return LastSOLPrice, nil
		}

		return price, err
	}

	LastSOLPrice = priceRes.Solana.USD
	return priceRes.Solana.USD, nil
}
