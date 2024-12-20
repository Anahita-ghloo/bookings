package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Aghloo/bookings/config"
	"github.com/Aghloo/bookings/pkg/models"
)

var app *config.AppConfig
var functions = template.FuncMap{}

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
func RenderTemplates(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCahe {
		//Get the template caache from the app config

		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplatecache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Could not get template from templatecache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplatecache() (map[string]*template.Template, error) {
	//mycache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files names *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// range through the files
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil

}
