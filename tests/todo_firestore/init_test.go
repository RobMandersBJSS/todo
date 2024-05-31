package todo_firestore_test

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
	"todo/modules/todo_firestore"
	"todo/tests/helpers"
)

func TestInitialiseFirebase(t *testing.T) {
	ctx := context.Background()
	_, basePath, _, _ := runtime.Caller(0)
	rootFilepath := filepath.Join(filepath.Dir(basePath), "../..")
	serviceAccountKeyPath := fmt.Sprintf("%s/serviceAccountKey.json", rootFilepath)

	t.Run("Create a new Firestore client", func(t *testing.T) {
		client, err := todo_firestore.CreateFirestoreClient(ctx, serviceAccountKeyPath)
		helpers.AssertNoError(t, err)
		defer client.Close()

		if client == nil {
			t.Error("Expected a Firestore client, but got nil.")
		}
	})
}
