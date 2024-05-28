package main

import (
	"fmt"
	"net/http"
	"time"
	"todo/modules/api_server"
	"todo/modules/todo_memory_store"
)

func main() {
	todoStore := todo_memory_store.TodoStore{}

	todoStore.Create("Do Laundry")
	todoStore.Create("Walk Dog")
	todoStore.Create("Buy Milk")

	server := api_server.NewApiServer(&todoStore, 10*time.Millisecond)

	fmt.Println("Listeneining on port")
	http.ListenAndServe(":5000", server)
}
