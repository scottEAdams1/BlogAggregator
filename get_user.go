package main

import (
	"net/http"

	"github.com/scottEAdams1/BlogAggregator/internal/auth"
)

func (cfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
