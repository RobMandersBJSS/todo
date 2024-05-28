package helpers

import (
	"time"
	"todo/modules/todo_store"
)

type SpyStore struct{}

func (s *SpyStore) GetItems() []todo_store.Todo {
	time.Sleep(15 * time.Millisecond)

	return []todo_store.Todo{}
}

func (s *SpyStore) Create(description string) (string, error) {
	return "", nil
}

func (s *SpyStore) Read(id string) (todo_store.Todo, error) {
	return todo_store.Todo{}, nil
}

func (s *SpyStore) UpdateItem(id, description string) error {
	return nil
}

func (s *SpyStore) ToggleItemStatus(id string) error {
	return nil
}

func (s *SpyStore) Delete(id string) error {
	return nil
}
