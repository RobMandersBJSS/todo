package api_server_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"time"
	"todo/modules/api_server"
	"todo/modules/todo_memory_store"
	"todo/modules/todo_store"
	"todo/tests/helpers"
)

func TestApiServerPATCH(t *testing.T) {
	timeout := 10 * time.Millisecond

	t.Run("PATCH /api/update updates an existing item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		id := todoStore.GetItems()[0].ID

		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte(fmt.Sprintf("{\"id\":\"%s\",\"description\":\"Updated Item\"}", id)))
		request, response := helpers.NewRequestResponse(t, http.MethodPatch, "/api/update", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)
		helpers.AssertEqual(t, len(todoStore.GetItems()), 1)
		helpers.AssertEqual(t, todoStore.GetItems()[0].Description, "Updated Item")
	})

	t.Run("PATCH /api/update returns 404 if item does not exist", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")

		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte("{\"id\":\"xyz\",\"description\":\"Updated Item\"}"))
		request, response := helpers.NewRequestResponse(t, http.MethodPatch, "/api/update", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusNotFound)
	})

	t.Run("PATCH /api/update returns the updated item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		id := todoStore.GetItems()[0].ID

		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte(fmt.Sprintf("{\"id\":\"%s\",\"description\":\"Updated Item\"}", id)))
		request, response := helpers.NewRequestResponse(t, http.MethodPatch, "/api/update", &requestBody)

		server.ServeHTTP(response, request)

		actual := helpers.UnmarshalBody[todo_store.Todo](t, response.Body.Bytes())

		helpers.AssertEqual(t, actual.Description, "Updated Item")
	})

	t.Run("PATCH /api/toggle changes the item completion status", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		id := todoStore.GetItems()[0].ID

		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", id)))
		request, response := helpers.NewRequestResponse(t, http.MethodPatch, "/api/toggle", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)
		helpers.AssertEqual(t, len(todoStore.GetItems()), 1)
		helpers.AssertEqual(t, todoStore.GetItems()[0].Complete, true)
	})

	t.Run("PATCH /api/toggle returns the updated item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		id := todoStore.GetItems()[0].ID

		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", id)))
		request, response := helpers.NewRequestResponse(t, http.MethodPatch, "/api/toggle", &requestBody)

		server.ServeHTTP(response, request)

		actual := helpers.UnmarshalBody[todo_store.Todo](t, response.Body.Bytes())

		helpers.AssertEqual(t, actual.Description, "Item 1")
		helpers.AssertEqual(t, actual.Complete, true)
	})

	t.Run("PATCH /api/toggle returns 404 if item does not exist", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")

		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte("{\"id\":\"xyz\"}"))
		request, response := helpers.NewRequestResponse(t, http.MethodPatch, "/api/toggle", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusNotFound)
	})
}
