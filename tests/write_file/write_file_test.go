package write_file_test

import (
	"fmt"
	"os"
	"testing"
	"todo/modules/read_file"
	"todo/modules/todo_memory_store"
	"todo/modules/todo_store"
	"todo/modules/write_file"
	"todo/tests/helpers"
)

func TestWriteFile(t *testing.T) {
	t.Run("Creates a new file", func(t *testing.T) {
		dir, removeDir := helpers.CreateTempDirectory(t)
		defer removeDir()

		filepath := fmt.Sprintf("%s/test.json", dir)

		items := []todo_store.Todo{}
		write_file.WriteFile(filepath, items...)

		_, readError := os.ReadFile(filepath)
		if readError != nil {
			t.Errorf("Expected a file at %q, but none present.", filepath)
		}
	})

	t.Run("Writes todos to a new file", func(t *testing.T) {
		dir, removeDir := helpers.CreateTempDirectory(t)
		defer removeDir()

		filepath := fmt.Sprintf("%s/test.json", dir)

		todoStore := todo_memory_store.TodoStore{}
		for i := 0; i < 10; i++ {
			todoStore.Create(fmt.Sprintf("Item %d", i))
		}
		items := todoStore.GetItems()

		write_file.WriteFile(filepath, items...)

		contents, err := read_file.ReadTodosFromFile(filepath)
		helpers.AssertNoError(t, err)

		helpers.AssertEqual(t, len(contents), 10)
	})
}
