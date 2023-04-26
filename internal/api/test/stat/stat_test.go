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

	// initialize router
	router = gin.Default()
	api := router.Group("/api")
	routers.StatRouter(api)
}

func processGet(url string) (int, []byte) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/%s", url), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	return w.Code, responseData
}

func TestOverall(t *testing.T) {
	statusCode, resBytes := processGet("stat/overall")
	assert.Equal(t, 200, statusCode)

	var nftRes types.StatRes
	err := json.Unmarshal(resBytes, &nftRes)
	assert.Nil(t, err)
}

func TestSearch(t *testing.T) {
	testURLs := testCase["search"]

	for _, testURL := range testURLs {
		statusCode, resBytes := processGet(testURL)
		assert.Equal(t, 200, statusCode)

		var searchRes types.SearchRes
		err := json.Unmarshal(resBytes, &searchRes)
		assert.Nil(t, err)
	}

}
