package handlers

import (
	config2 "github.com/powsianik/thinking-in-code/internal/config"
	models2 "github.com/powsianik/thinking-in-code/internal/models"
	render2 "github.com/powsianik/thinking-in-code/internal/render"
	"log"
	"net/http"
	"time"
)

type Repository struct {
	App *config2.AppConfig
}

var Repo *Repository

// CreateRepo create a new repository
func CreateRepo(app *config2.AppConfig) *Repository {
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

	render2.RenderTemplate(w, r,"home.page.tmpl", &models2.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello! From StringMap ;)"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render2.RenderTemplate(w, r,"about.page.tmpl", &models2.TemplateData{StringMap: stringMap})
}

// Post is the post page handler
func (m *Repository) Post(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	post := models2.PostData{
		ImageUrl: "static/img/about.jpg",
		Title: "Test post",
		Content: "Test post's content",
		CreatorName: "Przemysław Owsianik",
		CreatedAt: time.Now(),
	}

	render2.RenderTemplate(w, r,"post.page.tmpl",
		&models2.TemplateData{StringMap: stringMap, Post: post})
}

// Posts is the posts page handler
func (m *Repository) Posts(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	post := models2.PostData{
		ImageUrl: "static/img/about.jpg",
		Title: "Test post",
		Content: "Test post's content",
		CreatorName: "Przemysław Owsianik",
		CreatedAt: time.Now(),
	}

	render2.RenderTemplate(w, r,"posts.page.tmpl",
		&models2.TemplateData{StringMap: stringMap, Post: post})
}

// CreatePost is the page handler for creating new blog post
func (m *Repository) CreatePost(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render2.RenderTemplate(w, r,"createPost.page.tmpl", &models2.TemplateData{StringMap: stringMap})
}

func (m*Repository) SavePost(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	timeLayout := "2006-01-02T15:04:05.000Z"
	createPostTime, err := time.Parse(timeLayout, r.Form.Get("image"))
	if err == nil{
		log.Fatal()
		return
	}

	post := models2.PostData{
		ImageUrl: r.Form.Get("image"),
		Title: r.Form.Get("title"),
		Content: r.Form.Get("content"),
		CreatorName: r.Form.Get("creatorName"),
		CreatedAt: createPostTime,
	}

	render2.RenderTemplate(w, r,"post.page.tmpl", &models2.TemplateData{StringMap: stringMap, Post: post})
}