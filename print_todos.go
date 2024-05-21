package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Todo struct {
	Item   string
	Status bool
}

func PrintTodos(w io.Writer, todos ...Todo) error {
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

func PrintTodosJSON(w io.Writer, todos ...Todo) {
	output, _ := json.Marshal(todos)

	fmt.Fprint(w, string(output))
}
