package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/utkuuci/go-helloworld/pkg/config"
	"github.com/utkuuci/go-helloworld/pkg/handlers"
	"github.com/utkuuci/go-helloworld/pkg/render"
)

const portNumber = ":3000"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)
	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
