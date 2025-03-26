package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abhishekkumar11724/scratch/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error in creating new feed: %v", err))
	}

	respondWithJson(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {

	feed, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("error in getting all feeds: %v", err))
	}

	respondWithJson(w, 201, databaseFeedsToFeeds(feed))
}
