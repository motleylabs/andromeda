package test

import (
	"andromeda/internal/api/routers"
	"andromeda/pkg/service/entrance/types"
	"encoding/json"
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
}

func ProcessGet(url string) (int, []byte) {
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	return w.Code, responseData
}

func TestTrends(t *testing.T) {
	testURLs := testCase["trends"]

	for _, testURL := range testURLs {
		statusCode, resBytes := ProcessGet(testURL)
		assert.Equal(t, 200, statusCode)

		var trendRes types.TrendRes
		err := json.Unmarshal(resBytes, &trendRes)
		assert.Nil(t, err)
	}
}
