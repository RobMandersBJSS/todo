package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"todo/modules/store"
	"todo/modules/web_server"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

func main() {
	template := loadTemplates()
	todoStore := store.TodoStore{Items: []store.Todo{
		{ID: "0", Description: "Item 1", Complete: false},
		{ID: "1", Description: "Item 2", Complete: false},
	}}

	server := web_server.NewServer(template, &todoStore)

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
