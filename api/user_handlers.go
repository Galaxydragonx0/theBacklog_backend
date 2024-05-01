package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"theBacklog/backend/internal/database"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (s *Server) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
	}

	user, err := s.DB.CreateUser(r.Context(), database.CreateUserParams{

		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     params.Email,
	})

	errorString := "no error spotted"

	if err != nil {
		if db_error, ok := err.(*pq.Error); ok {
			if db_error.Code == "23505" {
				errorString = "This user already exists. Please enter a next email."
			}
		}

		respondWithError(w, 400, fmt.Sprintf("%v", errorString))
		return
	}

	err = s.DB.CreateGameList(r.Context(), database.CreateGameListParams{
		ID:     uuid.New(),
		UserID: user.ID,
		List:   []byte("[]"),
	})

	if err != nil {
		// if db_error, ok := err.(*pq.Error); ok {
		// 	if db_error.Code == "23505" {
		// 		errorString = "This user already exists. Please enter a next email."
		// 	}
		// }
		errorString = "Failed to create the user. The server is wack."
		respondWithError(w, 400, fmt.Sprintf("%v", errorString))
		return
	}

	err = s.DB.CreateBookList(r.Context(), database.CreateBookListParams{
		ID:     uuid.New(),
		UserID: user.ID,
		List:   []byte("[]"),
	})

	if err != nil {
		errorString = "Failed to create the user. The server is wack."
		respondWithError(w, 400, fmt.Sprintf("%v", errorString))
		return
	}

	err = s.DB.CreateMovieList(r.Context(), database.CreateMovieListParams{
		ID:     uuid.New(),
		UserID: user.ID,
		List:   []byte("[]"),
	})

	if err != nil {
		errorString = "Failed to create the user. The server is wack."
		respondWithError(w, 400, fmt.Sprintf("%v", errorString))
		return
	}

	err = s.DB.CreateShowList(r.Context(), database.CreateShowListParams{
		ID:     uuid.New(),
		UserID: user.ID,
		List:   []byte("[]"),
	})

	if err != nil {
		errorString = "Failed to create the user. The server is wack."
		respondWithError(w, 400, fmt.Sprintf("%v", errorString))
		return
	}

	err = s.DB.CreateCompletedList(r.Context(), database.CreateCompletedListParams{
		ID:     uuid.New(),
		UserID: user.ID,
		List:   []byte("[]"),
	})

	if err != nil {
		errorString = "Failed to create the user. The server is wack."
		respondWithError(w, 400, fmt.Sprintf("%v", errorString))
		return
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}

func (s *Server) handlerUserLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
	}

	err_uuid, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")
	user, err := s.DB.GetUser(r.Context(), params.Email)

	if err != nil && user.ID == err_uuid {
		respondWithError(w, 404, fmt.Sprint("User doesnt exist"))
	} else {
		respondWithJson(w, 200, databaseUserToUser(user))
	}
}

func (s *Server) handlerGetUserByAPI(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJson(w, 200, databaseUserToUser(user))
}
