package todo_memory_store_test

import (
	"testing"
	"todo/modules/todo_memory_store"
	"todo/modules/todo_store"
	"todo/tests/helpers"
)

func TestCreateItem(t *testing.T) {
	t.Run("Create new todo item", func(t *testing.T) {
		// item := todo_memory_store.Todo{ID: "", Description: "New Description", Complete: false}
		todoStore := todo_memory_store.TodoStore{}

		_, err := todoStore.Create("New Description")
		helpers.AssertNoError(t, err)

		items := todoStore.GetItems()

		helpers.AssertEqual(t, len(items), 1)
		helpers.AssertEqual(t, items[0].Description, "New Description")
		helpers.AssertEqual(t, items[0].Complete, false)
	})

	t.Run("Create todo returns the ID", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}

		id, err := todoStore.Create("New Description")
		helpers.AssertNoError(t, err)

		item, err := todoStore.Read(id)
		helpers.AssertNoError(t, err)

		helpers.AssertEqual(t, id, item.ID)
	})
}

func TestReadItem(t *testing.T) {
	t.Run("Read todo item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}

		todoStore.Create("Description 1")
		todoStore.Create("Description 2")

		id := todoStore.GetItems()[0].ID

		actual, err := todoStore.Read(id)
		helpers.AssertNoError(t, err)

		expected := "Description 1"

		helpers.AssertStructEqual(t, actual.Description, expected)
	})

	t.Run("Return an error if the ID cannot be found", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}

		todoStore.Create("Description 1")
		todoStore.Create("Description 2")

		_, err := todoStore.Read("XYZ")
		helpers.AssertError(t, err)
	})
}

func TestUpdateItem(t *testing.T) {
	t.Run("Update an existing todo item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}

		todoStore.Create("Description 1")
		todoStore.Create("Description 2")

		id := todoStore.GetItems()[0].ID

		err := todoStore.UpdateItem(id, "Updated Description")
		helpers.AssertNoError(t, err)

		actual, err := todoStore.Read(id)
		helpers.AssertNoError(t, err)

		expected := "Updated Description"

		helpers.AssertStructEqual(t, actual.Description, expected)
	})

	t.Run("Return an error if the ID cannot be found", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}

		todoStore.Create("Description 1")
		todoStore.Create("Description 2")

		err := todoStore.UpdateItem("10", "Updated Description")
		helpers.AssertError(t, err)
	})
}

func TestToggleItemStatus(t *testing.T) {
	t.Run("Toggle item completed status", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}

		todoStore.Create("Description 1")
		todoStore.Create("Description 2")

		id0 := todoStore.GetItems()[0].ID

		err := todoStore.ToggleItemStatus(id0)
		helpers.AssertNoError(t, err)

		actual := todoStore.GetItems()

		helpers.AssertStructEqual(t, actual[0].Complete, true)
	})

	t.Run("Return an error if ID cannot be found", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}

		todoStore.Create("Description 1")
		todoStore.Create("Description 2")

		err := todoStore.ToggleItemStatus("XYZ")
		helpers.AssertError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}

		todoStore.Create("Description 1")
		todoStore.Create("Description 2")

		id := todoStore.GetItems()[0].ID

		err := todoStore.Delete(id)
		helpers.AssertNoError(t, err)

		expected := []todo_store.Todo{
			{ID: "1", Description: "Description 2", Complete: true},
		}

		helpers.AssertStructEqual(t, len(todoStore.GetItems()), len(expected))
	})

	t.Run("Return an error if ID cannot be found", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}

		todoStore.Create("Description 1")
		todoStore.Create("Description 2")

		err := todoStore.Delete("10")
		helpers.AssertError(t, err)
	})
}
