package render

import (
	"bytes"
	"fmt"
	"github.com/powsianik/thinking-in-code/pkg/config"
	"github.com/powsianik/thinking-in-code/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

// SetAppConfig sets app config for render packages
func SetAppConfig(a *config.AppConfig){
	app = a
}

// RenderTemplate render a template with given name
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData){
	tc := app.TemplateCache

	t, isExists := tc[tmpl]
	if !isExists{
		log.Fatal("Template don't exists")
	}

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

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil{
		return nil, err
	}

	for _, page := range pages{
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil{
			return nil, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil{
			return nil, err
		}

		if len(matches) > 0{
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil{
				return nil, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

