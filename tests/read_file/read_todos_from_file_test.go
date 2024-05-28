package todos_test

import (
	"testing"
	"todo/modules/read_file"
	"todo/modules/todo_store"
	"todo/tests/helpers"
)

func TestReadTodosFromFile(t *testing.T) {
	t.Run("Reads json from file and returns todos", func(t *testing.T) {
		contents := "[{\"ID\":\"0\",\"Description\":\"Item 1\",\"Complete\":false},{\"ID\":\"1\",\"Description\":\"Item 2\",\"Complete\":true}]"

		tempFile, closeFile := helpers.CreateTempFile(t, "todos", contents)
		defer closeFile()

		actual, err := read_file.ReadTodosFromFile(tempFile.Name())
		helpers.AssertNoError(t, err)

		expected := []todo_store.Todo{
			{ID: "0", Description: "Item 1", Complete: false},
			{ID: "1", Description: "Item 2", Complete: true},
		}

		helpers.AssertSliceEqual(t, actual, expected)
	})

	t.Run("Returns error if bad file path supplied", func(t *testing.T) {
		_, err := read_file.ReadTodosFromFile("bad_file")

		helpers.AssertError(t, err)
	})

	t.Run("Returns error if invalid JSON in file", func(t *testing.T) {
		contents := "Not a JSON file..."

		tempFile, closeFile := helpers.CreateTempFile(t, "bad_file", contents)
		defer closeFile()

		_, err := read_file.ReadTodosFromFile(tempFile.Name())

		helpers.AssertError(t, err)
	})
}
