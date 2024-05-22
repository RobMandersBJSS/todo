package print_test

import (
	"bytes"
	"testing"
	"todo/print"
	"todo/store"
	"todo/tests/helpers"
)

func TestPrintTodo(t *testing.T) {
	t.Run("Print single todo item", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		item := store.Todo{ID: "0", Description: "Item 1", Complete: false}

		print.PrintTodo(buffer, item)

		actual := buffer.String()
		expected := "0 Incomplete Item 1\n"

		helpers.AssertEqual(t, actual, expected)
	})
}

func TestPrintTodos(t *testing.T) {
	t.Run("Print todos with statuses", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []store.Todo{
			{ID: "0", Description: "Item 1", Complete: false},
			{ID: "1", Description: "Item 2", Complete: true},
		}

		err := print.PrintTodos(buffer, todos...)
		helpers.AssertNoError(t, err)

		actual := buffer.String()
		expected := "0 Incomplete Item 1\n1 Complete   Item 2\n"

		helpers.AssertEqual(t, actual, expected)
	})

	t.Run("Print a message if there are no items to display", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []store.Todo{}

		err := print.PrintTodos(buffer, todos...)
		helpers.AssertNoError(t, err)

		actual := buffer.String()
		expected := print.NoItemsMessage

		helpers.AssertEqual(t, actual, expected)
	})
}
