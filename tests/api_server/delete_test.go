package api_server_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"time"
	"todo/modules/api_server"
	"todo/modules/todo_memory_store"
	"todo/tests/helpers"
)

func TestApiServerDELETE(t *testing.T) {
	timeout := 10 * time.Millisecond

	t.Run("DELETE /api/ deleted an item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")
		id := todoStore.GetItems()[0].ID

		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", id)))
		request, response := helpers.NewRequestResponse(t, http.MethodDelete, "/api", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)

		helpers.AssertEqual(t, len(todoStore.GetItems()), 1)
		helpers.AssertEqual(t, todoStore.GetItems()[0].Description, "Item 2")
	})

	t.Run("DELETE /api returns 404 if item does not exist", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")

		server := api_server.NewApiServer(&todoStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte("{\"id\":\"xyz\"}"))
		request, response := helpers.NewRequestResponse(t, http.MethodDelete, "/api", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusNotFound)
	})

	t.Run("DELETE /api times out", func(t *testing.T) {
		spyStore := helpers.SpyStore{}
		server := api_server.NewApiServer(&spyStore, timeout)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte("{\"id\":\"xyz\"}"))
		request, response := helpers.NewRequestResponse(t, http.MethodDelete, "/api", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusRequestTimeout)
	})

	t.Run("DELETE /api returns 400 if no body is provided", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		server := api_server.NewApiServer(&todoStore, timeout)

		request, response := helpers.NewRequestResponse(t, http.MethodDelete, "/api", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusBadRequest)
	})
}
