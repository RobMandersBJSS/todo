package todo_memory_store

import (
	"errors"
	"fmt"
	"sync"
	"todo/modules/todo_store"

	"github.com/google/uuid"
)

type TodoStore struct {
	mutex sync.Mutex
	items []todo_store.Todo
}

func (t *TodoStore) GetItems() []todo_store.Todo {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	return t.items
}

func (t *TodoStore) Create(description string) (string, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	id, err := uuid.NewV7()
	if err != nil {
		errorMessage := fmt.Sprintf("received the following error while creating a uuid for the new item: %v", err)
		return "", errors.New(errorMessage)
	}

	newItem := todo_store.Todo{
		ID:          id.String(),
		Description: description,
		Complete:    false,
	}

	t.items = append(t.items, newItem)

	return id.String(), nil
}

func (t *TodoStore) Read(id string) (todo_store.Todo, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	_, item, err := t.findItem(id)
	if err != nil {
		return todo_store.Todo{}, err
	}

	return item, nil
}

func (t *TodoStore) UpdateItem(id, description string) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	index, err := t.findItemIndex(id)
	if err != nil {
		return err
	}

	t.items[index].Description = description

	return nil
}

func (t *TodoStore) ToggleItemStatus(id string) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	index, err := t.findItemIndex(id)
	if err != nil {
		return err
	}

	t.items[index].Complete = !t.items[index].Complete

	return nil
}

func (t *TodoStore) Delete(id string) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	index, _, err := t.findItem(id)
	if err != nil {
		return err
	}

	t.items = append(t.items[:index], t.items[index+1:]...)

	return nil
}

func (t *TodoStore) findItem(id string) (index int, item todo_store.Todo, err error) {
	for index, item := range t.items {
		if item.ID == id {
			return index, item, nil
		}
	}

	errorMessage := fmt.Sprintf("could not locate item with id '%s'", id)
	return -1, todo_store.Todo{}, errors.New(errorMessage)
}

func (t *TodoStore) findItemIndex(id string) (int, error) {
	index, _, err := t.findItem(id)
	if err != nil {
		return -1, err
	}

	return index, nil
}
