package render

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	config "github.com/powsianik/thinking-in-code/internal/config"
	models "github.com/powsianik/thinking-in-code/internal/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},

	"dec": func(i int) int {
		return i - 1
	},
}

var app *config.AppConfig
var pathToTemplates = "./templates"

// SetAppConfig sets app config for render packages
func SetAppConfig(a *config.AppConfig){
	app = a
}

func AddDefaultData (td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate render a template with given name
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, data *models.TemplateData){
	tc := app.TemplateCache

	t, isExists := tc[tmpl]
	if !isExists{
		log.Fatal("Template don't exists")
	}

	data = AddDefaultData(data, r)

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil{
		fmt.Println("Error writing template to browser", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
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

