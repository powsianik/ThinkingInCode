package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	config "github.com/powsianik/thinking-in-code/internal/config"
	handlers "github.com/powsianik/thinking-in-code/internal/handlers"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/posts", handlers.Repo.Posts)
	mux.Get("/post/{id}", handlers.Repo.Post)
	mux.Get("/createPost", handlers.Repo.CreatePost)
	mux.Post("/createPost", handlers.Repo.SavePost)
	mux.Get("/editPost/{id}", handlers.Repo.EditPostGet)
	mux.Post("/editPost", handlers.Repo.EditPost)

	workDir, _ := os.Getwd()
	FileServer(mux, "/static", http.Dir(filepath.Join(workDir, "static")))
	return mux
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}