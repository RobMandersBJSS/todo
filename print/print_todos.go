package print

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"todo/todo"
)

func PrintTodos(w io.Writer, todos ...todo.Todo) error {
	if len(todos) < 1 {
		return errors.New("no todos provided in argument")
	}

	for _, todo := range todos {
		status := "Incomplete"
		if todo.Status {
			status = "Complete"
		}

		fmt.Fprintf(w, "%s - %s\n", todo.Item, status)
	}

	return nil
}

func PrintTodosJSON(w io.Writer, todos ...todo.Todo) error {
	if len(todos) < 1 {
		return errors.New("no todos provided in argument")
	}

	output, _ := json.Marshal(todos)
	fmt.Fprint(w, string(output))

	return nil
}
