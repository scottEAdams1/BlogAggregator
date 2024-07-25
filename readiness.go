package main

import "net/http"

func readiness(w http.ResponseWriter, r *http.Request) {
	type Status struct {
		Status string `json:"status"`
	}
	respondWithJSON(w, 200, Status{
		Status: "ok",
	})
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}
