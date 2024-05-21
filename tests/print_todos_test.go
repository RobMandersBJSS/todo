package todos_test

import (
	"bytes"
	"testing"
	"todo"
	"todo/tests/helpers"
)

func TestPrintTodos(t *testing.T) {
	t.Run("Print todos with statuses", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []todo.Todo{
			{Item: "Item 1", Status: false},
			{Item: "Item 2", Status: true},
		}

		err := todo.PrintTodos(buffer, todos...)
		helpers.AssertNoError(t, err)

		actual := buffer.String()
		expected := "Item 1 - Incomplete\nItem 2 - Complete\n"

		helpers.AssertEqual(t, actual, expected)
	})

	t.Run("Return an error if no todos provided", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []todo.Todo{}

		err := todo.PrintTodos(buffer, todos...)

		helpers.AssertError(t, err)
	})
}

func TestPrintTodosJSON(t *testing.T) {
	t.Run("Print todos in JSON format", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []todo.Todo{
			{Item: "Item 1", Status: false},
			{Item: "Item 2", Status: true},
		}

		err := todo.PrintTodosJSON(buffer, todos...)
		helpers.AssertNoError(t, err)

		actual := buffer.String()
		expected := "[{\"Item\":\"Item 1\",\"Status\":false},{\"Item\":\"Item 2\",\"Status\":true}]"

		helpers.AssertEqual(t, actual, expected)
	})

	t.Run("Return an error if no todos provided", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []todo.Todo{}

		err := todo.PrintTodosJSON(buffer, todos...)

		helpers.AssertError(t, err)
	})
}
