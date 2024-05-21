package todos_test

import (
	"testing"
	"todo"
	"todo/tests/helpers"
)

func TestReadTodosFromFile(t *testing.T) {
	t.Run("Reads json from file and returns todos", func(t *testing.T) {
		contents := "[{\"Item\":\"Item 1\",\"Status\":false},{\"Item\":\"Item 2\",\"Status\":true}]"

		tempFile, closeFile := helpers.CreateTempFile(t, "todos", contents)
		defer closeFile()

		actual, err := todo.ReadTodosFromFile(tempFile.Name())
		helpers.AssertNoError(t, err)

		expected := []todo.Todo{
			{Item: "Item 1", Status: false},
			{Item: "Item 2", Status: true},
		}

		helpers.AssertSliceEqual(t, actual, expected)
	})

	t.Run("Returns error if bad file path supplied", func(t *testing.T) {
		_, err := todo.ReadTodosFromFile("bad_file")

		helpers.AssertError(t, err)
	})

	t.Run("Returns error if invalid JSON in file", func(t *testing.T) {
		contents := "Not a JSON file..."

		tempFile, closeFile := helpers.CreateTempFile(t, "bad_file", contents)
		defer closeFile()

		_, err := todo.ReadTodosFromFile(tempFile.Name())

		helpers.AssertError(t, err)
	})
}
