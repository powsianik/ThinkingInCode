package main

import (
	"github.com/alexedwards/scs/v2"
	config2 "github.com/powsianik/thinking-in-code/internal/config"
	handlers2 "github.com/powsianik/thinking-in-code/internal/handlers"
	render2 "github.com/powsianik/thinking-in-code/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config2.AppConfig
var session *scs.SessionManager

func main(){
	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	templateCache, err := render2.CreateTemplateCache()
	if err != nil{
		log.Fatal("Error while creating template cache: ", err)
		return
	}

	app.TemplateCache = templateCache
	render2.SetAppConfig(&app)

	var repo = handlers2.CreateRepo(&app)
	handlers2.SetRepository(repo)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}
}

