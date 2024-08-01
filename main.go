package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/scottEAdams1/BlogAggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	//Load .env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}
	port := os.Getenv("PORT")
	dbURL := os.Getenv("dbURL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	go startScraping(dbQueries, 10, time.Minute)

	//Handlers
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", readiness)
	mux.HandleFunc("GET /v1/err", errorHandler)

	mux.HandleFunc("POST /v1/users", apiCfg.createUsers)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.getUser))

	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.createFeeds))
	mux.HandleFunc("GET /v1/feeds", apiCfg.getFeeds)

	mux.HandleFunc("POST /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.feedFollow))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.deleteFeedFollows))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.middlewareAuth(apiCfg.getFeedFollows))

	//Create server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	//Run server
	log.Fatal(server.ListenAndServe())
}
