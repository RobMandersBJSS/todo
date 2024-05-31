package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"todo/modules/todo_firestore"
	"todo/modules/web_app"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

func main() {
	template := loadTemplates()
	// todoStore := todo_memory_store.TodoStore{}

	// todoStore.Create("Do Laundry")
	// todoStore.Create("Walk Dog")
	// todoStore.Create("Buy Milk")

	ctx := context.Background()
	client, err := todo_firestore.CreateFirestoreClient(ctx, "serviceAccountKey.json")
	if err != nil {
		log.Fatal(err)
	}

	todoStore := todo_firestore.NewTodoFirestore(ctx, client, "items")

	server := web_app.NewServer(template, &todoStore)

	fmt.Println("Web app started: http://localhost:5000")
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
