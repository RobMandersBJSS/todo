package api_server_test

import (
	"bytes"
	"net/http"
	"testing"
	"time"
	"todo/modules/api_server"
	"todo/modules/todo_memory_store"
	"todo/modules/todo_store"
	"todo/tests/helpers"
)

func TestApiServerPOST(t *testing.T) {
	timeout := 10 * time.Millisecond

	t.Run("POST /api/ returns 201 and creates a new item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte("{\"description\":\"New Item\"}"))
		request, response := helpers.NewRequestResponse(t, http.MethodPost, "/api", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusCreated)
		helpers.AssertEqual(t, len(todoStore.GetItems()), 1)
		helpers.AssertEqual(t, todoStore.GetItems()[0].Description, "New Item")
	})

	t.Run("POST /api/ returns the new item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte("{\"description\":\"New Item\"}"))
		request, response := helpers.NewRequestResponse(t, http.MethodPost, "/api", &requestBody)

		server.ServeHTTP(response, request)

		actual := helpers.UnmarshalBody[todo_store.Todo](t, response.Body.Bytes())

		helpers.AssertEqual(t, actual.Description, "New Item")
	})
}
