package handlers

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/powsianik/thinking-in-code/internal/config"
	"github.com/powsianik/thinking-in-code/internal/render"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"
var functions = template.FuncMap{}

func getRoutes() http.Handler{
	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	templateCache, err := CreateTestTemplateCache()
	if err != nil{
		log.Fatal("Error while creating template cache: ", err)
	}

	app.TemplateCache = templateCache
	app.UseCache = true
	render.SetAppConfig(&app)

	var repo = CreateRepo(&app)
	SetRepository(repo)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/posts", Repo.Posts)
	mux.Get("/post", Repo.Post)
	mux.Get("/createPost", Repo.CreatePost)
	mux.Post("/createPost", Repo.SavePost)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}

// CreateTemplateCache creates a template cache as a map
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil{
		return nil, err
	}

	for _, page := range pages{
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return nil, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil{
			return nil, err
		}

		if len(matches) > 0{
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil{
				return nil, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}