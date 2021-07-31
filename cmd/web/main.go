package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/powsianik/thinking-in-code/pkg/config"
	"github.com/powsianik/thinking-in-code/pkg/handlers"
	"github.com/powsianik/thinking-in-code/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main(){
	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("Error while creating template cache: ", err)
		return
	}

	app.TemplateCache = templateCache
	render.SetAppConfig(&app)

	var repo = handlers.CreateRepo(&app)
	handlers.SetRepository(repo)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}
}

