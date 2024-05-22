package store

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Todo struct {
	ID          string `json:"ID"`
	Description string `json:"Description"`
	Complete    bool   `json:"Complete"`
}

type TodoStore struct {
	Items []Todo
}

func (t *TodoStore) Create(description string) (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		errorMessage := fmt.Sprintf("received the following error while creating a uuid for the new item: %v", err)
		return "", errors.New(errorMessage)
	}

	newItem := Todo{
		ID:          id.String(),
		Description: description,
		Complete:    false,
	}

	t.Items = append(t.Items, newItem)

	return id.String(), nil
}

func (t *TodoStore) Read(id string) (Todo, error) {
	_, item, err := t.findItem(id)
	if err != nil {
		return Todo{}, err
	}

	return item, nil
}

func (t *TodoStore) FindItemIndex(id string) (int, error) {
	index, _, err := t.findItem(id)
	if err != nil {
		return -1, err
	}

	return index, nil
}

func (t *TodoStore) UpdateItem(id string, description string) error {
	index, err := t.FindItemIndex(id)
	if err != nil {
		return err
	}

	t.Items[index].Description = description

	return nil
}

func (t *TodoStore) ToggleItemStatus(id string) error {
	index, err := t.FindItemIndex(id)
	if err != nil {
		return err
	}

	t.Items[index].Complete = !t.Items[index].Complete

	return nil
}

func (t *TodoStore) findItem(id string) (index int, item Todo, err error) {
	for index, item := range t.Items {
		if item.ID == id {
			return index, item, nil
		}
	}

	errorMessage := fmt.Sprintf("could not locate item with id '%s'", id)
	return -1, Todo{}, errors.New(errorMessage)
}
