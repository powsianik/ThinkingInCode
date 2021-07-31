package handlers

import (
	"errors"
	"fmt"
	"github.com/powsianik/thinking-in-code/pkg/config"
	"github.com/powsianik/thinking-in-code/pkg/models"
	"github.com/powsianik/thinking-in-code/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// CreateRepo create a new repository
func CreateRepo(app *config.AppConfig) *Repository{
	return &Repository{
		App: app,
	}
}

// SetRepository sets repository for handlers
func SetRepository(r *Repository){
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request){
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello! From StringMap ;)"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func Divide(w http.ResponseWriter, r *http.Request){
	x := 100.0
	y := 0.0
	f, err := divide(x, y)
	if err != nil{
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("Result of %f divided by %f is %f", x, y, f))
}

func divide(x,y float64)(float64, error){
	if y <= 0{
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x/y
	return result, nil
}