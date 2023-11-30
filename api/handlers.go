package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"theBacklog/backend/internal/database"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
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

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create a user: %v", err))
	}

	respondWithJson(w, 200, databaseUserToUser(user))
}

func (s *Server) handlerGetMovieList(w http.ResponseWriter, r *http.Request, user database.User) {
	movieList, err := s.DB.GetMovieListByUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
	}

	respondWithJson(w, 200, movieList)
}

// get the list by user id
// take the list and parse from string to json
// add the new selected json info to the array
// stringify the json array and store to db

func (s *Server) handlerUpdateMovieList(w http.ResponseWriter, r *http.Request, user database.User) {

	respBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error occured in reading the request body:", err)
	}

	err = s.DB.UpdateMovieList(r.Context(),
		database.UpdateMovieListParams{
			UserID: user.ID,
			List:   respBytes,
		})

	if err != nil {
		log.Fatal("Error occured, could not update the movie list:", err)
	}
	respondWithJson(w, 200, struct{}{})
}

func (s *Server) movieSearchQuery(w http.ResponseWriter, r *http.Request) {

	movieName := chi.URLParam(r, "mName")

	fmt.Println("This is the name:", movieName)

	//tmdb_api_key := os.Getenv("tmdb_api_key")

	movieName = strings.ReplaceAll(movieName, " ", "%20")
	// bookName = strings.ReplaceAll(bookName, " ", "+")
	endpoint := fmt.Sprintf(`https://api.themoviedb.org/3/search/movie?query=%s&include_adult=false&language=en-US&page=1`, movieName)

	fmt.Println("This is the endpoint:", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", os.Getenv("tmdb_auth_token"))

	//response, err := http.Get(endpoint)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	respondToSearch(w, 200, body)

}

func respondToSearch(w http.ResponseWriter, code int, resp []byte) {

	// this uses the response endpoint and simply returns it to the user as a request
	w.Header().Add("Content-Type", "applicaiton/json")
	w.WriteHeader(code)
	w.Write(resp)

}
