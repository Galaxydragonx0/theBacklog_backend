package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"theBacklog/backend/internal/database"
	"time"
	"unsafe"

	"github.com/dimuska139/rawg-sdk-go/v3"
	"github.com/go-chi/chi"
)

func (s *Server) handlerGetGameList(w http.ResponseWriter, r *http.Request, user database.User) {
	gameList, err := s.DB.GetGameListByUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get games: %v", err))
	}

	respondWithJson(w, 300, gameList)
}

func (s *Server) handlerUpdateGameList(w http.ResponseWriter, r *http.Request, user database.User) {
	respBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error occured in reading the request body:", err)
	}

	err = s.DB.UpdateGameList(r.Context(),
		database.UpdateGameListParams{
			UserID: user.ID,
			List:   respBytes,
		})

	if err != nil {
		log.Fatal("Error occured, could not update the movie list:", err)
	}
	respondWithJson(w, 200, struct{}{})
}

func (s *Server) handlerGetMovieList(w http.ResponseWriter, r *http.Request, user database.User) {
	movieList, err := s.DB.GetMovieListByUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get movies: %v", err))
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

func (s *Server) handlerGetCompletedList(w http.ResponseWriter, r *http.Request, user database.User) {
	completedList, err := s.DB.GetCompletedListByUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get titles: %v", err))
		return
	}

	respondWithJson(w, 200, completedList)
}

func (s *Server) handlerUpdateCompletedList(w http.ResponseWriter, r *http.Request, user database.User) {

	respBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error occured in reading the request body:", err)
	}

	err = s.DB.UpdateCompletedList(r.Context(),
		database.UpdateCompletedListParams{
			UserID: user.ID,
			List:   respBytes,
		})

	if err != nil {
		log.Fatal("Error occured, could not update the completed list:", err)
	}
	respondWithJson(w, 200, struct{}{})
}

func (s *Server) movieSearchQuery(w http.ResponseWriter, r *http.Request) {

	movieName := chi.URLParam(r, "mName")
	pageNum := chi.URLParam(r, "page")
	fmt.Println("This is the name:", movieName)

	movieName = strings.ReplaceAll(movieName, " ", "%20")

	endpoint := fmt.Sprintf(`https://api.themoviedb.org/3/search/movie?query=%s&include_adult=false&language=en-US&page=%s`, movieName, pageNum)

	fmt.Println("This is the endpoint:", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", os.Getenv("tmdb_auth_token"))

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

func (s *Server) gameSearchQuery(w http.ResponseWriter, r *http.Request) {

	gameName := chi.URLParam(r, "gName")
	pageNum, _ := strconv.Atoi(chi.URLParam(r, "page"))
	config := rawg.Config{
		ApiKey:   os.Getenv("rawg_game_api"), // Your personal API key (see https://rawg.io/apidocs)
		Language: "en",
		Rps:      5,
	}

	client := rawg.NewClient(http.DefaultClient, &config)
	filter := rawg.NewGamesFilter().
		SetSearch(gameName).
		SetPage(pageNum).
		SetPageSize(10).
		ExcludeCollection(1).
		WithoutParents()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*500))
	defer cancel()
	data, total, err := client.GetGames(ctx, filter)

	fmt.Println("this is the data from search", data, "this is the amount", total)
	if err != nil {
		log.Fatal(err)
	}

	//convert data into bytes using whatever this is
	size := unsafe.Sizeof(data)
	gameResponseBytes := unsafe.Slice((*byte)(unsafe.Pointer(&data)), int(size))

	//use data to send back to user
	respondToSearch(w, 200, gameResponseBytes)
}

func respondToSearch(w http.ResponseWriter, code int, resp []byte) {

	// this uses the response endpoint and simply returns it to the user as a request
	w.Header().Add("Content-Type", "applicaiton/json")
	w.WriteHeader(code)
	w.Write(resp)

}
