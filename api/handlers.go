package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"theBacklog/backend/internal/database"

	"github.com/go-chi/chi"
)

func (s *Server) handlerGetBookList(w http.ResponseWriter, r *http.Request, user database.User) {
	bookList, err := s.DB.GetBookListByUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get games: %v", err))
	}

	respondWithJson(w, 200, bookList)
}

func (s *Server) handlerUpdateBookList(w http.ResponseWriter, r *http.Request, user database.User) {
	respBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error occured in reading the request body:", err)
	}

	err = s.DB.UpdateBookList(r.Context(),
		database.UpdateBookListParams{
			UserID: user.ID,
			List:   respBytes,
		})

	if err != nil {
		log.Fatal("Error occured, could not update the movie list:", err)
	}
	respondWithJson(w, 200, struct{}{})
}

func (s *Server) handlerGetGameList(w http.ResponseWriter, r *http.Request, user database.User) {
	gameList, err := s.DB.GetGameListByUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get games: %v", err))
	}

	respondWithJson(w, 200, gameList)
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

func (s *Server) handlerGetShowList(w http.ResponseWriter, r *http.Request, user database.User) {
	movieList, err := s.DB.GetShowListByUser(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get movies: %v", err))
	}

	respondWithJson(w, 200, movieList)
}

func (s *Server) handlerUpdateShowList(w http.ResponseWriter, r *http.Request, user database.User) {

	respBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error occured in reading the request body:", err)
	}

	err = s.DB.UpdateShowList(r.Context(),
		database.UpdateShowListParams{
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
			UserID:  user.ID,
			Column1: respBytes,
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

func (s *Server) showSearchQuery(w http.ResponseWriter, r *http.Request) {

	showName := chi.URLParam(r, "sName")
	pageNum := chi.URLParam(r, "page")
	fmt.Println("This is the name:", showName)

	showName = strings.ReplaceAll(showName, " ", "%20")

	endpoint := fmt.Sprintf(`https://api.themoviedb.org/3/search/tv?query=%s&include_adult=false&language=en-US&page=%s`, showName, pageNum)

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
	pageNum := chi.URLParam(r, "page")
	searchOffset := 0
	searchPageNum, _ := strconv.Atoi(pageNum)

	fmt.Println("This is the name:", gameName)

	if searchPageNum > 1 {
		searchOffset = (searchPageNum - 1) * 15
	}

	gameName = strings.ReplaceAll(gameName, " ", "%20")

	endpoint := fmt.Sprintf(`http://www.giantbomb.com/api/games/?api_key=%s&limit=15&offset=%d&filter=name:%s&format=json&field_list=original_release_date,id,genres,api_detail_url,name,image,deck&page=%s,`, os.Getenv("game_api_key"), searchOffset, gameName, pageNum)

	fmt.Println("This is the endpoint:", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("accept", "application/json")
	// req.Header.Add("Authorization", os.Getenv("tmdb_auth_token"))

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

func (s *Server) bookSearchQuery(w http.ResponseWriter, r *http.Request) {
	bookName := chi.URLParam(r, "bName")
	pageNum := chi.URLParam(r, "page")
	searchOffset := 0
	endpoint := ""
	searchPageNum, _ := strconv.Atoi(pageNum)
	bookName = strings.ReplaceAll(bookName, " ", "%20")

	fmt.Println("This is the name:", bookName)

	if searchPageNum > 1 {
		searchOffset = (searchPageNum - 1) * 15
		endpoint = fmt.Sprintf(`https://www.googleapis.com/books/v1/volumes?q=%s&startIndex=%d&maxResults=%s&key=%s`, bookName, searchOffset, "14", os.Getenv("book_api"))
	}

	if searchPageNum == 1 {
		endpoint = fmt.Sprintf(`https://www.googleapis.com/books/v1/volumes?q=%s&startIndex=%s&maxResults=%s&key=%s`, bookName, "0", "14", os.Getenv("book_api"))
	}

	fmt.Println("This is the endpoint:", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("accept", "application/json")

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
