package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"andromeda/internal/api/routers"
	"andromeda/internal/api/state"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
)

func main() {
	var err error

	logf, err := rotateLogs.New("./logs/%Y-%m-%d.logs")
	if err != nil {
		log.Printf("Main >> RotateLogs.New; %s", err.Error())
	}
	log.SetOutput(logf)

	if err := godotenv.Load(); err != nil {
		log.Printf("Dotenv Load; %s", err.Error())
	}

	// set mode
	if os.Getenv("ENV") == "PRODUCTION" {
		log.Printf("now production mode")
		gin.SetMode(gin.ReleaseMode)
	}

	if !state.Ensure() {
		log.Fatal("Main >> state.Ensure; failed to ensure state")
	}

	go state.Runloop()

	// router initialization
	r := routers.Initialize()

	// define server
	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// listen and serve
	err = s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
