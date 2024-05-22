package print_test

import (
	"bytes"
	"testing"
	"todo/print"
	"todo/store"
	"todo/tests/helpers"
)

func TestPrintTodos(t *testing.T) {
	t.Run("Print todos with statuses", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []store.Todo{
			{ID: 0, Item: "Item 1", Complete: false},
			{ID: 1, Item: "Item 2", Complete: true},
		}

		err := print.PrintTodos(buffer, todos...)
		helpers.AssertNoError(t, err)

		actual := buffer.String()
		expected := "Item 1 - Incomplete\nItem 2 - Complete\n"

		helpers.AssertEqual(t, actual, expected)
	})

	t.Run("Return an error if no todos provided", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []store.Todo{}

		err := print.PrintTodos(buffer, todos...)

		helpers.AssertError(t, err)
	})
}
