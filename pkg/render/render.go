package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/utkuuci/go-helloworld/pkg/config"
	"github.com/utkuuci/go-helloworld/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// creates new templates
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {

		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Couldn't get template from template cache")
	}

	td = AddDefaultData(td)
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(w)

	if err != nil {
		fmt.Println(err)
	}
	// parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	// err = parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./template/*.page.html")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./template/*.layout.html")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./template/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
