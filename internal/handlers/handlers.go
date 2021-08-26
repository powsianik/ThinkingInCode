package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	config "github.com/powsianik/thinking-in-code/internal/config"
	"github.com/powsianik/thinking-in-code/internal/dbAccess"
	"github.com/powsianik/thinking-in-code/internal/editorjs"
	"github.com/powsianik/thinking-in-code/internal/forms"
	"github.com/powsianik/thinking-in-code/internal/helpers"
	models "github.com/powsianik/thinking-in-code/internal/models"
	render "github.com/powsianik/thinking-in-code/internal/render"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	render.RenderTemplate(w, r,"home.page.tmpl", &models.TemplateData{})
}

// Post is the post page handler
func (m *Repository) Post(w http.ResponseWriter, r *http.Request){
	postId := chi.URLParam(r, "id")
	objID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		panic(err)
	}

	post := dbAccess.Read("_id", objID)
	post.Content = editorjs.HTML(string(post.Content))
	render.RenderTemplate(w, r,"post.page.tmpl",
		&models.TemplateData{Post: post})
}

// Posts is the posts page handler
func (m *Repository) Posts(w http.ResponseWriter, r *http.Request){
	posts := dbAccess.ReadAll()

	render.RenderTemplate(w, r,"posts.page.tmpl",
		&models.TemplateData{Posts: posts})
}

// CreatePost is the page handler for render page for creating new blog post
func (m *Repository) CreatePost(w http.ResponseWriter, r *http.Request){
	var emptyPost models.PostData
	data := make(map[string]interface{})
	data["SavePost"] = emptyPost

	render.RenderTemplate(w, r,"createPost.page.tmpl", &models.TemplateData{ Form: forms.New(nil), Data: data})
}

func (m*Repository) SavePost(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil{
		helpers.ServerError(w, err)
		return
	}

	post := models.PostData{
		ImageUrl: r.Form.Get("image"),
		Title: r.Form.Get("title"),
		Content: r.Form.Get("content"),
		CreatorName: r.Form.Get("creatorName"),
		Description: r.Form.Get("description"),
		CreatedAt: time.Now().Format("2006-01-02"),
	}

	fmt.Println(post.Content)

	form := forms.New(r.PostForm)

	form.Required("creatorName", "title", "description", "content")
	if !form.Valid(){
		data := make(map[string]interface{})
		data["SavePost"] = post

		render.RenderTemplate(w, r,"createPost.page.tmpl", &models.TemplateData{Form: form, Data: data})
		return
	}

	dbAccess.Write(post)
	post.Content = editorjs.HTML(string(post.Content))
	render.RenderTemplate(w, r,"post.page.tmpl", &models.TemplateData{Post: post})
}

