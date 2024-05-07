package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"theBacklog/backend/internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Server struct {
	addr    string
	handler *chi.Mux
	DB      *database.Queries
}

func NewServer() *Server {

	godotenv.Load(".env")
	servePort := os.Getenv("PORT")

	//Router
	router := chi.NewRouter()

	//Setting up CORS
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	dbURL := os.Getenv("DB_CONN")
	if dbURL == "" {
		log.Fatal("Database Connection String was not found in the environment file!")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cant connect to the database:", err)
	}

	db := database.New(conn)

	return &Server{
		addr:    ":" + servePort,
		handler: router,
		DB:      db,
	}
}

func (S *Server) Start() error {
	server := &http.Server{
		Handler: S.handler,
		Addr:    S.addr,
	}

	router := S.handler

	router.Get("/new",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			resp := `{"name": "Drew"}`
			w.Write([]byte(resp))
		})

	router.Post("/register", S.handlerCreateUser)

	router.Post("/login", S.handlerUserLogin)

	router.Get("/getUser", S.middlewareAuth(S.handlerGetUserByAPI))

	// MOVIE LIST
	router.Get("/getMovies", S.middlewareAuth(S.handlerGetMovieList))

	router.Post("/movies", S.middlewareAuth(S.handlerUpdateMovieList))

	// GAMES LIST
	router.Get("/getGames", S.middlewareAuth(S.handlerGetGameList))

	router.Post("/games", S.middlewareAuth(S.handlerUpdateGameList))

	// COMPLETED TITLES
	router.Post("/completed", S.middlewareAuth(S.handlerUpdateCompletedList))

	router.Get("/getCompleted", S.middlewareAuth(S.handlerGetCompletedList))

	// SEARCH QUERY
	router.Get("/movieSearch/{mName}/{page}", S.movieSearchQuery)

	router.Get("/gameSearch/{gName}/{page}", S.gameSearchQuery)

	fmt.Printf("Server is serving on: http://localhost%s \n", S.addr)

	return server.ListenAndServe()

}
