package handlers

import (
	"errors"
	"github.com/powsianik/thinking-in-code/pkg/config"
	"github.com/powsianik/thinking-in-code/pkg/models"
	"github.com/powsianik/thinking-in-code/pkg/render"
	"net/http"
	"time"
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

// Post is the post page handler
func (m *Repository) Post(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "post.page.tmpl",
		&models.TemplateData{StringMap: stringMap,
								ImageUrl: "static/img/about.jpg",
								Title: "Test post",
								Content: "Test post's content",
								CreatorName: "Przemysław Owsianik",
								CreatedAt: time.Now() })
}

// Posts is the posts page handler
func (m *Repository) Posts(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "post.page.tmpl",
		&models.TemplateData{StringMap: stringMap,
			ImageUrl: "static/img/about.jpg",
			Title: "Test post",
			Content: "Test post's content",
			CreatorName: "Przemysław Owsianik",
			CreatedAt: time.Now() })
}

// CreatePost is the page handler for creating new blog post
func (m *Repository) CreatePost(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "createPost.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func divide(x,y float64)(float64, error){
	if y <= 0{
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x/y
	return result, nil
}