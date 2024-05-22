package store_test

import (
	"testing"
	"todo/store"
	"todo/tests/helpers"
)

func TestCreateItem(t *testing.T) {
	t.Run("Create new todo item", func(t *testing.T) {
		// item := store.Todo{ID: "", Description: "New Description", Complete: false}
		todoStore := store.TodoStore{Items: []store.Todo{}}

		_, err := todoStore.Create("New Description")
		helpers.AssertNoError(t, err)

		helpers.AssertEqual(t, len(todoStore.Items), 1)
		helpers.AssertEqual(t, todoStore.Items[0].Description, "New Description")
		helpers.AssertEqual(t, todoStore.Items[0].Complete, false)
	})

	t.Run("Create todo returns the ID", func(t *testing.T) {
		todoStore := store.TodoStore{Items: []store.Todo{}}

		id, err := todoStore.Create("New Description")
		helpers.AssertNoError(t, err)

		item, err := todoStore.Read(id)
		helpers.AssertNoError(t, err)

		helpers.AssertEqual(t, id, item.ID)
	})
}

func TestReadItem(t *testing.T) {
	t.Run("Read todo item", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "0", Description: "Description 1", Complete: false},
			{ID: "1", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		actual, err := todoStore.Read("0")
		helpers.AssertNoError(t, err)

		expected := store.Todo{ID: "0", Description: "Description 1", Complete: false}

		helpers.AssertStructEqual(t, actual, expected)
	})

	t.Run("Return an error if the ID cannot be found", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "0", Description: "Description 1", Complete: false},
			{ID: "1", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		_, err := todoStore.Read("10")
		helpers.AssertError(t, err)
	})
}

func TestFindItemIndex(t *testing.T) {
	t.Run("FInd index of item from ID", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "10", Description: "Description 1", Complete: false},
			{ID: "20", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		actual, _ := todoStore.FindItemIndex("10")
		expected := 0

		helpers.AssertEqual(t, actual, expected)
	})

	t.Run("Return an error if the ID cannot be found", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "10", Description: "Description 1", Complete: false},
			{ID: "20", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		_, err := todoStore.FindItemIndex("0")
		helpers.AssertError(t, err)
	})
}

func TestUpdateItem(t *testing.T) {
	t.Run("Update an existing todo item", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "0", Description: "Description 1", Complete: false},
			{ID: "1", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		err := todoStore.UpdateItem("0", "Updated Description")
		helpers.AssertNoError(t, err)

		actual, err := todoStore.Read("0")
		helpers.AssertNoError(t, err)

		expected := store.Todo{ID: "0", Description: "Updated Description", Complete: false}

		helpers.AssertStructEqual(t, actual, expected)
	})

	t.Run("Return an error if the ID cannot be found", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "0", Description: "Description 1", Complete: false},
			{ID: "1", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		err := todoStore.UpdateItem("10", "Updated Description")
		helpers.AssertError(t, err)
	})
}

func TestToggleItemStatus(t *testing.T) {
	t.Run("Toggle item completed status", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "0", Description: "Description 1", Complete: false},
			{ID: "1", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		err1 := todoStore.ToggleItemStatus("0")
		helpers.AssertNoError(t, err1)

		err2 := todoStore.ToggleItemStatus("1")
		helpers.AssertNoError(t, err2)

		actual := todoStore.Items

		expected := []store.Todo{
			{ID: "0", Description: "Description 1", Complete: true},
			{ID: "1", Description: "Description 2", Complete: false},
		}

		helpers.AssertStructEqual(t, actual, expected)
	})

	t.Run("Return an error if ID cannot be found", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "0", Description: "Description 1", Complete: false},
			{ID: "1", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		err := todoStore.ToggleItemStatus("10")
		helpers.AssertError(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete item", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "0", Description: "Description 1", Complete: false},
			{ID: "1", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		err := todoStore.Delete("0")
		helpers.AssertNoError(t, err)

		expected := []store.Todo{
			{ID: "1", Description: "Description 2", Complete: true},
		}

		helpers.AssertStructEqual(t, todoStore.Items, expected)
	})

	t.Run("Return an error if ID cannot be found", func(t *testing.T) {
		todos := []store.Todo{
			{ID: "0", Description: "Description 1", Complete: false},
			{ID: "1", Description: "Description 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		err := todoStore.Delete("10")
		helpers.AssertError(t, err)
	})
}
