package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo/modules/print"
	"todo/modules/store"
)

func main() {
	running := true
	todos := store.TodoStore{Items: []store.Todo{}}
	input := bufio.NewScanner(os.Stdin)

	fmt.Println("Todo List App")
	fmt.Println("Enter a command or type 'help' for more information.")

	for running {
		inputString := readInput(input)
		command := strings.Split(inputString, " ")[0]

		switch command {
		case "help":
			printHelpPage()
		case "list":
			print.PrintTodos(os.Stdout, todos.Items...)
		case "new":
			newItem(&todos, inputString)
		case "print":
			printItem(&todos, inputString)
		case "delete":
			deleteItem(&todos, inputString)
		case "update":
			updateItem(&todos, inputString)
		case "toggle":
			updateStatus(&todos, inputString)
		case "exit":
			running = false
		default:
			fmt.Println("Invalid command. Type 'help' for the command list.")
		}
	}
}

func newItem(todos *store.TodoStore, inputString string) {
	description := inputString[3:]

	id, err := todos.Create(description)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("New item created:")
		printItemWithID(todos, id)
	}
}

func printItem(todos *store.TodoStore, inputString string) {
	id := inputString[6:]
	printItemWithID(todos, id)
}

func deleteItem(todos *store.TodoStore, inputString string) {
	id := inputString[7:]
	err := todos.Delete(id)
	if err != nil {
		fmt.Printf("Could not find an item with id '%s'\n", id)
	} else {
		fmt.Println("Item deleted.")
	}
}

func updateItem(todos *store.TodoStore, inputString string) {
	id := strings.Split(inputString, " ")[1]
	description := inputString[7+len(id):]

	err := todos.UpdateItem(id, description)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("Item updated:")
		printItemWithID(todos, id)
	}
}

func updateStatus(todos *store.TodoStore, inputString string) {
	id := inputString[7:]
	err := todos.ToggleItemStatus(id)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println("Status updated:")
		printItemWithID(todos, id)
	}
}

func printItemWithID(todos *store.TodoStore, id string) {
	item, err := todos.Read(id)
	if err != nil {
		fmt.Println(err)
	} else {
		print.PrintTodo(os.Stdout, item)
	}
}

func printHelpPage() {
	fmt.Println("")
	fmt.Println("list                       List all todos.")
	fmt.Println("new [description]          Create a new todo item.")
	fmt.Println("print [id]                 Print a single item.")
	fmt.Println("delete [id]                Delete todo item with a given ID.")
	fmt.Println("update [id] [description]  Update todo of a given ID with a new description.")
	fmt.Println("toggle [id]                Toggle the complete status of the item.")
	fmt.Println("help                       Print command list.")
	fmt.Println("exit                       Exit the application.")
	fmt.Println("")
}

func readInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
