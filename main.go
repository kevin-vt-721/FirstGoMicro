package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kmlowe/FirstGoMicroservice/homepage"
	"github.com/kmlowe/FirstGoMicroservice/server"
)

var (
//If we have these we can do https connection (required)
//CertFile    = os.Getenv("CERT_FILE")
//KeyFile     = os.Getenv("KEY_FILE")
//ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "first ", log.LstdFlags|log.Lshortfile)
	h := homepage.NewHandlers(logger)
	mux := http.NewServeMux() //mux is something that takes a request
	h.SetupRoutes(mux)

	srv := server.New(mux, ":8080")
	logger.Println("Starting server...")
	err := srv.ListenAndServe()
	//https
	//err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
