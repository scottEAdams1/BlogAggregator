package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}
	port := os.Getenv("PORT")

	//Handlers
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", readiness)
	mux.HandleFunc("GET /v1/err", errorHandler)

	//Create server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	//Run server
	server.ListenAndServe()
}
