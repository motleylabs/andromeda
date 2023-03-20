package utils

import (
	"andromeda/pkg/service/entrance"
	"andromeda/pkg/service/hyperspace"
	"os"
	"strings"
)

func GetProvider() entrance.DataProvider {
	provider := os.Getenv("PROVIDER")
	var dataProvider entrance.DataProvider

	// get trends data
	if strings.ToLower(provider) == "solsniper" {
		dataProvider = hyperspace.Hyperspace{}
	}

	return dataProvider
}
