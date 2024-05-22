package main

import (
	"embed"
	"fmt"
	"net/http"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

func main() {
	// todos := store.TodoStore{Items: []store.Todo{}}
	// print.PrintTodos(os.Stdout, todos.Items...)

	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(handler))

	fmt.Println("Listening on port 5000...")
	http.ListenAndServe(":5000", router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test")
}
