package test

import (
	"andromeda/internal/api/controllers"
	"andromeda/internal/api/routers"
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

func init() {
	// load env
	curDir, _ := os.Getwd()
	envDir := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(curDir)))), ".env")

	if err := godotenv.Load(envDir); err != nil {
		log.Printf("Dotenv Load; %s", err.Error())
	}

	// initialize router
	router = gin.Default()
	api := router.Group("/api")
	routers.RPCRouter(api)
}

func processGet(url string) (int, []byte) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/%s", url), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	return w.Code, responseData
}

func TestReport(t *testing.T) {
	statusCode, resBytes := processGet("rpc/report")
	assert.Equal(t, 200, statusCode)

	var nftRes controllers.ReportRes
	err := json.Unmarshal(resBytes, &nftRes)
	assert.Nil(t, err)
}
