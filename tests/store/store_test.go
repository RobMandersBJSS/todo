package store_test

import (
	"testing"
	"todo/store"
	"todo/tests/helpers"
)

func TestCreateItem(t *testing.T) {
	t.Run("Create new todo item", func(t *testing.T) {
		item := store.Todo{ID: 0, Item: "New Item", Complete: false}
		todoStore := store.TodoStore{Items: []store.Todo{}}

		todoStore.Create(item)

		expected := store.TodoStore{Items: []store.Todo{item}}

		helpers.AssertStructEqual(t, todoStore, expected)
	})
}

func TestReadItem(t *testing.T) {
	t.Run("Read todo item", func(t *testing.T) {
		todos := []store.Todo{
			{ID: 0, Item: "Item 1", Complete: false},
			{ID: 1, Item: "Item 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		actual, _ := todoStore.Read(0)
		expected := store.Todo{ID: 0, Item: "Item 1", Complete: false}

		helpers.AssertStructEqual(t, actual, expected)
	})

	t.Run("Return an error if the ID cannot be found", func(t *testing.T) {
		todos := []store.Todo{
			{ID: 0, Item: "Item 1", Complete: false},
			{ID: 1, Item: "Item 2", Complete: true},
		}
		todoStore := store.TodoStore{Items: todos}

		_, err := todoStore.Read(10)
		helpers.AssertError(t, err)
	})
}
