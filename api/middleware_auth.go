/*
This code is used to reduce the repition of logic in authenication
we use throughout the program.

The function allows us to convert the functions into the standard HandlerFunc interface.
*/
package api

import (
	"fmt"
	"net/http"
	"theBacklog/backend/internal/database"
	"theBacklog/backend/internal/database/auth"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (s *Server) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error %v", err))
			return
		}

		user, err := s.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldnt get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
