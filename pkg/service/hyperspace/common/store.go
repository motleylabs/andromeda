package common

import (
	"andromeda/pkg/service/web3"
	"fmt"
	"time"

	"github.com/gin-contrib/cache/persistence"
)

func FetchSOLPrice(store *persistence.InMemoryStore) {
	solPrice, err := web3.GetSOLPrice()
	if err == nil {
		store.Set("andromeda-sol-price", solPrice, -1)
	}
}

func GetSOLPrice(store *persistence.InMemoryStore) (float64, error) {
	retries := 0
	for {
		if retries == 10 {
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
