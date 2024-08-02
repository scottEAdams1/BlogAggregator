package main

import (
	"net/http"
	"strconv"

	"github.com/scottEAdams1/BlogAggregator/internal/database"
)

func (cfg *apiConfig) getPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	limitString := r.URL.Query().Get("limit")
	limit := 10
	if limitGiven, err := strconv.Atoi(limitString); err == nil {
		limit = limitGiven
	}
	posts, err := cfg.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	convertedPosts := make([]Post, len(posts))
	for i, post := range posts {
		convertedPosts[i] = databasePostToPost(post)
	}
	respondWithJSON(w, http.StatusOK, convertedPosts)
}
