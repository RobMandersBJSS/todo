package main

import (
	"log"
	"os"
	"todo/print"
	"todo/read_file"
)

func main() {
	todos, err := read_file.ReadTodosFromFile("todos.json")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	print.PrintTodos(os.Stdout, todos...)
}
