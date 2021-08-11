package handlers

import (
	config "github.com/powsianik/thinking-in-code/internal/config"
	"github.com/powsianik/thinking-in-code/internal/forms"
	models "github.com/powsianik/thinking-in-code/internal/models"
	render "github.com/powsianik/thinking-in-code/internal/render"
	"log"
	"net/http"
	"time"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// CreateRepo create a new repository
func CreateRepo(app *config.AppConfig) *Repository {
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

	render.RenderTemplate(w, r,"home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello! From StringMap ;)"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, r,"about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

// Post is the post page handler
func (m *Repository) Post(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	post := models.PostData{
		ImageUrl: "static/img/about.jpg",
		Title: "Test post",
		Content: "Test post's content",
		CreatorName: "Przemysław Owsianik",
		CreatedAt: time.Now(),
	}

	render.RenderTemplate(w, r,"post.page.tmpl",
		&models.TemplateData{StringMap: stringMap, Post: post})
}

// Posts is the posts page handler
func (m *Repository) Posts(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	post := models.PostData{
		ImageUrl: "static/img/about.jpg",
		Title: "Test post",
		Content: "Test post's content",
		CreatorName: "Przemysław Owsianik",
		CreatedAt: time.Now(),
	}

	render.RenderTemplate(w, r,"posts.page.tmpl",
		&models.TemplateData{StringMap: stringMap, Post: post})
}

// CreatePost is the page handler for render page for creating new blog post
func (m *Repository) CreatePost(w http.ResponseWriter, r *http.Request){
	var emptyPost models.PostData
	data := make(map[string]interface{})
	data["SavePost"] = emptyPost

	render.RenderTemplate(w, r,"createPost.page.tmpl", &models.TemplateData{ Form: forms.New(nil), Data: data})
}

func (m*Repository) SavePost(w http.ResponseWriter, r *http.Request){
	stringMap := make(map[string]string)

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	err := r.ParseForm()
	if err != nil{
		log.Println(err)
		return
	}

	timeLayout := "2006-01-02T15:04:05.000Z"
	createPostTime, err := time.Parse(timeLayout, r.Form.Get("image"))
	if err == nil{
		log.Fatal()
		return
	}

	post := models.PostData{
		ImageUrl: r.Form.Get("image"),
		Title: r.Form.Get("title"),
		Content: r.Form.Get("content"),
		CreatorName: r.Form.Get("creatorName"),
		Description: r.Form.Get("description"),
		CreatedAt: createPostTime,
	}

	form := forms.New(r.PostForm)

	form.Required("creatorName", "title", "description", "content")
	if !form.Valid(){
		data := make(map[string]interface{})
		data["SavePost"] = post

		render.RenderTemplate(w, r,"createPost.page.tmpl", &models.TemplateData{StringMap: stringMap, Form: form, Data: data})
		return
	}

	render.RenderTemplate(w, r,"post.page.tmpl", &models.TemplateData{StringMap: stringMap, Post: post})
}