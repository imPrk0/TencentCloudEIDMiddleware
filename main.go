package main

import (
	"TencentCloudEIDMiddleware/bootstrap"
	"TencentCloudEIDMiddleware/routers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bootstrap.Init()
}

func main() {
	api := routers.InitRouter()
	server := &http.Server{
		Handler: api,
	}

	server.Addr = os.Getenv("SERVER_ADDR")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to listen to %q: %s", os.Getenv("SERVER_ADDR"), err)
	}
}
