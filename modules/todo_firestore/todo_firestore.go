package todo_firestore

import (
	"context"
	"errors"
	"fmt"
	"todo/modules/todo_store"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type TodoFirestore struct {
	ctx        context.Context
	client     *firestore.Client
	collection string
}

func NewTodoFirestore(ctx context.Context, client *firestore.Client, collection string) TodoFirestore {
	todoFirestore := new(TodoFirestore)
	todoFirestore.ctx = ctx
	todoFirestore.client = client
	todoFirestore.collection = collection

	return *todoFirestore
}

func (f *TodoFirestore) GetItems() []todo_store.Todo {
	docs := f.client.Collection(f.collection).Documents(f.ctx)

	var items []todo_store.Todo

	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		data := doc.Data()
		items = append(items, todo_store.Todo{
			ID:          doc.Ref.ID,
			Description: data["description"].(string),
			Complete:    data["complete"].(bool),
		})
	}

	return items
}

func (f *TodoFirestore) Create(description string) (string, error) {
	ref, _, err := f.client.Collection(f.collection).Add(f.ctx, map[string]interface{}{
		"description": description,
		"complete":    false,
	})
	if err != nil {
		return "", err
	}

	return ref.ID, nil
}

func (f *TodoFirestore) Read(id string) (todo_store.Todo, error) {
	ref := f.client.Collection(f.collection).Doc(id)
	doc, err := ref.Get(f.ctx)
	if err != nil {
		return todo_store.Todo{}, err
	}

	data := doc.Data()
	description := fmt.Sprint(data["description"])
	var complete = data["complete"].(bool)

	return todo_store.Todo{
		ID:          id,
		Description: description,
		Complete:    complete,
	}, nil
}

func (f *TodoFirestore) UpdateItem(id, description string) error {
	_, err := f.client.Collection(f.collection).Doc(id).Update(f.ctx, []firestore.Update{
		{
			Path:  "description",
			Value: description,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (f *TodoFirestore) ToggleItemStatus(id string) error {
	item, readError := f.Read(id)
	if readError != nil {
		return errors.New("item does not exist")
	}

	_, updateError := f.client.Collection(f.collection).Doc(id).Update(f.ctx, []firestore.Update{
		{
			Path:  "complete",
			Value: !item.Complete,
		},
	})
	if updateError != nil {
		return updateError
	}

	return nil
}

func (f *TodoFirestore) Delete(id string) error {
	_, readError := f.Read(id)
	if readError != nil {
		return errors.New("item does not exist")
	}

	_, deleteError := f.client.Collection(f.collection).Doc(id).Delete(f.ctx)
	if deleteError != nil {
		return deleteError
	}

	return nil
}
