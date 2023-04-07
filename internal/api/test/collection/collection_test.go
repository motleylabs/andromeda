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

	"github.com/joho/godotenv"
	assert "github.com/stretchr/testify/require"
)

var router = routers.Initialize()
var testCase map[string][]string

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
	if err := json.Unmarshal(data, &testCase); err != nil {
		log.Fatalf("Test case file format is wrong")
	}

	callPreliminaryAPI()
}

func processGet(url string) (int, []byte) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/%s", url), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	return w.Code, responseData
}

func callPreliminaryAPI() {
	preliminaryURL := testCase["preliminary"]
	if len(preliminaryURL) > 0 {
		processGet(preliminaryURL[0])
	} else {
		log.Fatalf("Preliminary URL is missing")
	}
}

func TestTrends(t *testing.T) {
	testURLs := testCase["trends"]

	for _, testURL := range testURLs {
		statusCode, resBytes := processGet(testURL)
		assert.Equal(t, 200, statusCode)

		var trendRes types.TrendRes
		err := json.Unmarshal(resBytes, &trendRes)
		assert.Nil(t, err)
	}
}

func TestDetail(t *testing.T) {
	testURLs := testCase["detail"]

	for _, testURL := range testURLs {
		statusCode, resBytes := processGet(testURL)
		assert.Equal(t, 200, statusCode)

		var detailRes types.Collection
		err := json.Unmarshal(resBytes, &detailRes)
		assert.Nil(t, err)
	}
}

func TestTimeseries(t *testing.T) {
	testURLs := testCase["timeseries"]

	for _, testURL := range testURLs {
		statusCode, resBytes := processGet(testURL)
		assert.Equal(t, 200, statusCode)

		var seriesRes types.TimeSeriesRes
		err := json.Unmarshal(resBytes, &seriesRes)
		assert.Nil(t, err)
	}
}

func TestNFTs(t *testing.T) {
	testURLs := testCase["nfts"]

	for _, testURL := range testURLs {
		statusCode, resBytes := processGet(testURL)
		assert.Equal(t, 200, statusCode)

		var nftRes types.NFTRes
		err := json.Unmarshal(resBytes, &nftRes)
		assert.Nil(t, err)
	}
}

func TestActivities(t *testing.T) {
	testURLs := testCase["activities"]

	for _, testURL := range testURLs {
		statusCode, resBytes := processGet(testURL)
		assert.Equal(t, 200, statusCode)

		var activityRes types.ActivityRes
		err := json.Unmarshal(resBytes, &activityRes)
		assert.Nil(t, err)
	}
}
