package main

import (
	"github.com/alexedwards/scs/v2"
	config "github.com/powsianik/thinking-in-code/internal/config"
	handlers "github.com/powsianik/thinking-in-code/internal/handlers"
	"github.com/powsianik/thinking-in-code/internal/helpers"
	render "github.com/powsianik/thinking-in-code/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main(){
	err := run()
	if err != nil{
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil{
		log.Fatal(err)
	}
}

func run() error{
	app.IsProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("Error while creating template cache: ", err)
		return err
	}

	app.TemplateCache = templateCache
	render.SetAppConfig(&app)

	var repo = handlers.CreateRepo(&app)
	handlers.SetRepository(repo)

	helpers.NewHelpers(&app)

	return nil
}