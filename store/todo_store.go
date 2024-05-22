package store

import (
	"errors"
	"fmt"
)

type Todo struct {
	ID       int    `json:"ID"`
	Item     string `json:"Item"`
	Complete bool   `json:"Complete"`
}

type TodoStore struct {
	Items []Todo
}

func (t *TodoStore) Create(item Todo) {
	t.Items = append(t.Items, item)
}

func (t *TodoStore) Read(id int) (Todo, error) {
	for _, item := range t.Items {
		if item.ID == id {
			return item, nil
		}
	}

	errorMessage := fmt.Sprintf("could not locate item with id '%d'", id)
	return Todo{}, errors.New(errorMessage)
}
