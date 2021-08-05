package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	config2 "github.com/powsianik/thinking-in-code/internal/config"
	handlers2 "github.com/powsianik/thinking-in-code/internal/handlers"
	"net/http"
)

func routes(app *config2.AppConfig) http.Handler{
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers2.Repo.Home)
	mux.Get("/about", handlers2.Repo.About)
	mux.Get("/posts", handlers2.Repo.Posts)
	mux.Get("/post", handlers2.Repo.Post)
	mux.Get("/createPost", handlers2.Repo.CreatePost)
	mux.Post("/createPost", handlers2.Repo.SavePost)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}