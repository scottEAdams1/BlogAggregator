package main

import (
	"net/http"

	"github.com/scottEAdams1/BlogAggregator/internal/database"
)

func (cfg *apiConfig) getUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
