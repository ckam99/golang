package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/ckam225/golang/webapp/internal/models"

	"github.com/ckam225/golang/webapp/config"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]models.Bag)
	data["message"] = "Welcome golang"
	renderTemplate(w, "home", &models.TemplateData{
		Data: data,
	})
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about", &models.TemplateData{})
}

var appConfif *config.Config

func CreateTemplates(app *config.Config) {
	appConfif = app
}

func renderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	templateCache := appConfif.TemplateCache
	tmpl, ok := templateCache[templateName+".page.tmpl"]
	if !ok {
		http.Error(w, fmt.Sprintf("Temple %s.page.html does not exist", templateName), http.StatusInternalServerError)
		return
	}
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, data)
	buffer.WriteTo(w)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))
		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}
		cache[name] = tmpl
	}
	return cache, nil
}
