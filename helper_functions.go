package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Respond to request with an error
func respondWithError(w http.ResponseWriter, code int, msg string) {
	//Server error
	if code > 499 {
		log.Printf("Error decoding parameters: %s", msg)
	}

	//Return an error struct with response
	type errorStruct struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorStruct{
		Error: msg,
	})
}

// Respond in JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//Set response to JSON
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	//Convert struct into JSON
	jsonResponse, jsonErr := json.Marshal(payload)
	if jsonErr != nil {
		log.Printf("Error marshalling JSON: %s", jsonErr)
		w.WriteHeader(500)
		return
	}

	//Add code(e.g. 200, 500) to header
	w.WriteHeader(code)

	//Add JSON to the body of the response
	w.Write(jsonResponse)
}
