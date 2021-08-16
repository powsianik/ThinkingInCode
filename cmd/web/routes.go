package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	config "github.com/powsianik/thinking-in-code/internal/config"
	handlers "github.com/powsianik/thinking-in-code/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/posts", handlers.Repo.Posts)
	mux.Get("/post", handlers.Repo.Post)
	mux.Get("/createPost", handlers.Repo.CreatePost)
	mux.Post("/createPost", handlers.Repo.SavePost)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}