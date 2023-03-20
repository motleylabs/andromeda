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
	switch strings.ToLower(provider) {
	case "hyperspace":
		dataProvider = hyperspace.Hyperspace{}
	default:
		dataProvider = hyperspace.Hyperspace{}
	}

	return dataProvider
}
