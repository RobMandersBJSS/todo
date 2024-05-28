package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"todo/modules/store"
	"todo/modules/web_app"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

func main() {
	template := loadTemplates()
	todoStore := store.TodoStore{}

	todoStore.Create("Do Laundry")
	todoStore.Create("Walk Dog")
	todoStore.Create("Buy Milk")

	server := web_app.NewServer(template, &todoStore)

	http.ListenAndServe(":5000", server)
}

func loadTemplates() *template.Template {
	templates, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		fmt.Printf("Experienced the following error parsing the template: %v", err)
		os.Exit(1)
	}

	return templates
}
