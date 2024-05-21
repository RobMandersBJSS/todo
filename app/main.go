package main

import (
	"log"
	"os"
	"todo"
)

func main() {
	todos, err := todo.ReadTodosFromFile("todos.json")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	todo.PrintTodos(os.Stdout, todos...)
}
