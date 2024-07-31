package main

import "net/http"

func (cfg *apiConfig) getFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	convertedFeeds := make([]Feed, len(feeds))
	for i, feed := range feeds {
		convertedFeeds[i] = databaseFeedToFeed(feed)
	}
	respondWithJSON(w, http.StatusOK, convertedFeeds)
}
