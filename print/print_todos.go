package print

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"todo/store"
)

const NoItemsMessage = "No items in list."

func PrintTodo(w io.Writer, item store.Todo) {
	status := "Incomplete"
	if item.Complete {
		status = "Complete  "
	}

	fmt.Fprintf(w, "%s %s %s\n", item.ID, status, item.Description)
}

func PrintTodos(w io.Writer, todos ...store.Todo) error {
	if len(todos) < 1 {
		fmt.Fprint(w, NoItemsMessage)
	}

	for _, item := range todos {
		PrintTodo(w, item)
	}

	return nil
}

func PrintTodosJSON(w io.Writer, todos ...store.Todo) error {
	if len(todos) < 1 {
		return errors.New("no todos provided in argument")
	}

	output, _ := json.Marshal(todos)
	fmt.Fprint(w, string(output))

	return nil
}
