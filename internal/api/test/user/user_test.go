package test

import (
	"andromeda/internal/api/routers"
	"andromeda/pkg/service/entrance/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	assert "github.com/stretchr/testify/require"
)

var router *gin.Engine
var wallets []string

func init() {
	// load env
	curDir, _ := os.Getwd()
	envDir := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(curDir)))), ".env")

	if err := godotenv.Load(envDir); err != nil {
		log.Printf("Dotenv Load; %s", err.Error())
	}

	// load test cases
	data, err := os.ReadFile("test_cases.json")
	if err != nil {
		log.Fatalf("Test case file not found")
	}
	if err := json.Unmarshal(data, &wallets); err != nil {
		log.Fatalf("Test case file format is wrong")
	}

	// initialize router
	router = gin.Default()
	api := router.Group("/api")
	routers.UserRouter(api)
}

func processGet(url string) (int, []byte) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/%s", url), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	return w.Code, responseData
}

func TestNFTs(t *testing.T) {
	for _, wallet := range wallets {
		statusCode, resBytes := processGet(fmt.Sprintf("users/nfts?address=%s", wallet))
		assert.Equal(t, 200, statusCode)

		var nftRes types.UserNFT
		err := json.Unmarshal(resBytes, &nftRes)
		assert.Nil(t, err)
	}
}

func TestActivity(t *testing.T) {
	for _, wallet := range wallets {
		statusCode, resBytes := processGet(fmt.Sprintf(
			"users/activities?address=%s&limit=10&offset=0&activity_types=[\"transaction\", \"listing\"]", wallet,
		))
		assert.Equal(t, 200, statusCode)

		var activityRes types.ActivityRes
		err := json.Unmarshal(resBytes, &activityRes)
		assert.Nil(t, err)
	}
}

func TestOffer(t *testing.T) {
	for _, wallet := range wallets {
		statusCode, resBytes := processGet(fmt.Sprintf(
			"users/offers?address=%s&limit=10&offset=0", wallet,
		))
		assert.Equal(t, 200, statusCode)

		var offerRes types.ActivityRes
		err := json.Unmarshal(resBytes, &offerRes)
		assert.Nil(t, err)
	}
}
