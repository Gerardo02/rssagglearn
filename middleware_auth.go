package main

import (
	"fmt"
	"net/http"

	"github.com/Gerardo02/rssagglearn/internal/auth"
	"github.com/Gerardo02/rssagglearn/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldnt get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
