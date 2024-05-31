package todo_firestore_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"todo/modules/todo_firestore"
	"todo/modules/todo_store"
	"todo/tests/helpers"

	"cloud.google.com/go/firestore"
)

// TODO: Figure out a way to mock a Firestore database.

const collection = "test_collection"

var serviceAccountKeyPath = fmt.Sprintf("%s/serviceAccountKey.json", helpers.GetRootDir())

func TestNewTodoFirestore(t *testing.T) {
	ctx := context.Background()

	t.Run("Create a new a new Firestore object", func(t *testing.T) {
		client, _ := todo_firestore.CreateFirestoreClient(ctx, serviceAccountKeyPath)
		todoStore := todo_firestore.NewTodoFirestore(ctx, client, collection)

		helpers.AssertEqual(t, reflect.TypeOf(todoStore).String(), "todo_firestore.TodoFirestore")
	})
}

func TestGetItems(t *testing.T) {
	ctx := context.Background()
	client, _ := todo_firestore.CreateFirestoreClient(ctx, serviceAccountKeyPath)
	todoStore := todo_firestore.NewTodoFirestore(ctx, client, collection)

	_, deleteCollection := addItemsToFirestore(t, ctx, client, "Item 1", "Item 2")
	defer deleteCollection()

	t.Run("Returns all items in database", func(t *testing.T) {
		items := todoStore.GetItems()
		helpers.AssertEqual(t, len(items), 2)
	})
}

func TestCreateItem(t *testing.T) {
	ctx := context.Background()

	t.Run("Creates a new item in the database", func(t *testing.T) {
		client, _ := todo_firestore.CreateFirestoreClient(ctx, serviceAccountKeyPath)
		todoStore := todo_firestore.NewTodoFirestore(ctx, client, collection)

		description := "New Item"
		id, err := todoStore.Create(description)
		helpers.AssertNoError(t, err)
		defer client.Collection(collection).Doc(id).Delete(ctx)

		doc := getDoc(t, ctx, client, id)

		actual := fmt.Sprintf("%v", doc["description"])

		helpers.AssertEqual(t, actual, description)
	})
}

func TestReadItem(t *testing.T) {
	ctx := context.Background()
	client, _ := todo_firestore.CreateFirestoreClient(ctx, serviceAccountKeyPath)
	todoStore := todo_firestore.NewTodoFirestore(ctx, client, collection)

	ids, deleteCollection := addItemsToFirestore(t, ctx, client, "Item 1")
	defer deleteCollection()

	t.Run("Reads an item from the database", func(t *testing.T) {
		actual, err := todoStore.Read(ids[0])
		helpers.AssertNoError(t, err)

		expected := todo_store.Todo{
			ID:          ids[0],
			Description: "Item 1",
			Complete:    false,
		}

		helpers.AssertStructEqual(t, actual, expected)
	})

	t.Run("Returns an error if the item does not exist", func(t *testing.T) {
		_, err := todoStore.Read("xyz")
		helpers.AssertError(t, err)
	})
}

func TestUpdateItemDescription(t *testing.T) {
	ctx := context.Background()
	client, _ := todo_firestore.CreateFirestoreClient(ctx, serviceAccountKeyPath)
	todoStore := todo_firestore.NewTodoFirestore(ctx, client, collection)

	ids, deleteCollection := addItemsToFirestore(t, ctx, client, "Item 1")
	defer deleteCollection()

	t.Run("Updates an item description in the database", func(t *testing.T) {
		description := "Updated Item"

		err := todoStore.UpdateItem(ids[0], description)
		helpers.AssertNoError(t, err)

		doc := getDoc(t, ctx, client, ids[0])

		actual := fmt.Sprintf("%v", doc["description"])

		helpers.AssertEqual(t, actual, description)
	})

	t.Run("Returns an error if the item does not exist", func(t *testing.T) {
		err := todoStore.UpdateItem("xyz", "Updated Item")
		helpers.AssertError(t, err)
	})
}

func TestToggleItemStatus(t *testing.T) {
	ctx := context.Background()
	client, _ := todo_firestore.CreateFirestoreClient(ctx, serviceAccountKeyPath)
	todoStore := todo_firestore.NewTodoFirestore(ctx, client, collection)

	ids, deleteCollection := addItemsToFirestore(t, ctx, client, "Item 1")
	defer deleteCollection()

	t.Run("Toggles the status of an item in the database", func(t *testing.T) {
		err := todoStore.ToggleItemStatus(ids[0])
		helpers.AssertNoError(t, err)

		doc := getDoc(t, ctx, client, ids[0])

		helpers.AssertEqual(t, doc["complete"].(bool), true)
	})

	t.Run("Returns an error if the item does not exist", func(t *testing.T) {
		err := todoStore.ToggleItemStatus("xyz")
		helpers.AssertError(t, err)
	})
}

func TestDeleteItem(t *testing.T) {
	ctx := context.Background()
	client, _ := todo_firestore.CreateFirestoreClient(ctx, serviceAccountKeyPath)
	todoStore := todo_firestore.NewTodoFirestore(ctx, client, collection)

	ids, deleteCollection := addItemsToFirestore(t, ctx, client, "Item 1", "Item 2")
	defer deleteCollection()

	t.Run("Deletes an item from the database", func(t *testing.T) {
		err := todoStore.Delete(ids[0])
		helpers.AssertNoError(t, err)

		docs, _ := client.Collection(collection).Documents(ctx).GetAll()

		helpers.AssertEqual(t, len(docs), 1)
	})

	t.Run("Returns an error if the item does not exist", func(t *testing.T) {
		err := todoStore.Delete("xyz")
		helpers.AssertError(t, err)
	})
}

func addItemsToFirestore(t testing.TB, ctx context.Context, client *firestore.Client, items ...string) ([]string, func()) {
	t.Helper()

	var ids []string

	for _, item := range items {
		ref, _, err := client.Collection(collection).Add(ctx, map[string]interface{}{
			"description": item,
			"complete":    false,
		})
		helpers.AssertNoError(t, err)

		ids = append(ids, ref.ID)
	}

	return ids, func() {
		for _, id := range ids {
			client.Collection(collection).Doc(id).Delete(ctx)
		}
	}
}

func getDoc(t testing.TB, ctx context.Context, client *firestore.Client, id string) map[string]interface{} {
	t.Helper()

	ref := client.Collection(collection).Doc(id)
	doc, err := ref.Get(ctx)
	helpers.AssertNoError(t, err)

	return doc.Data()
}
