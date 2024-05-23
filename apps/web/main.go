package main

import (
	"embed"
	"fmt"
	"html/template"
	"os"
	"todo/modules/web_server"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

func main() {
	// todos := store.TodoStore{Items: []store.Todo{}}
	// print.PrintTodos(os.Stdout, todos.Items...)

	template := loadTemplates()
	server := web_server.NewServer(template)

	server.Start()
}

func loadTemplates() *template.Template {
	templates, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		fmt.Printf("Experienced the following error parsing the template: %v", err)
		os.Exit(1)
	}

	return templates
}
