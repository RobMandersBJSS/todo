package print_test

import (
	"bytes"
	"testing"
	"todo/modules/print"
	"todo/modules/store"
	"todo/tests/helpers"
)

func TestPrintTodosJSON(t *testing.T) {
	t.Run("Print todos in JSON format", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []store.Todo{
			{ID: "0", Description: "Item 1", Complete: false},
			{ID: "1", Description: "Item 2", Complete: true},
		}

		err := print.PrintTodosJSON(buffer, todos...)
		helpers.AssertNoError(t, err)

		actual := buffer.String()
		expected := "[{\"ID\":\"0\",\"Description\":\"Item 1\",\"Complete\":false},{\"ID\":\"1\",\"Description\":\"Item 2\",\"Complete\":true}]"

		helpers.AssertEqual(t, actual, expected)
	})

	t.Run("Return an error if no todos provided", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		todos := []store.Todo{}

		err := print.PrintTodosJSON(buffer, todos...)

		helpers.AssertError(t, err)
	})
}
