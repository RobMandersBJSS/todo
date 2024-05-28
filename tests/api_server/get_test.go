package api_server_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"
	"todo/modules/api_server"
	"todo/modules/todo_memory_store"
	"todo/modules/todo_store"
	"todo/tests/helpers"
)

func TestApiServerGET(t *testing.T) {
	timeout := 10 * time.Millisecond

	t.Run("GET /api returns a list of todos as JSON", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")

		server := api_server.NewApiServer(&todoStore, timeout)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/api", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)

		actual := helpers.UnmarshalBody[[]todo_store.Todo](t, response.Body.Bytes())

		helpers.AssertEqual(t, len(actual), 2)
		helpers.AssertEqual(t, actual[0].Description, "Item 1")
		helpers.AssertEqual(t, actual[1].Description, "Item 2")
	})

	t.Run("GET /api returns 404 if there are no items", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		server := api_server.NewApiServer(&todoStore, timeout)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/api", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusNotFound)
	})

	t.Run("GET /api/{id} returns the specified item", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")
		id := todoStore.GetItems()[0].ID

		server := api_server.NewApiServer(&todoStore, timeout)
		url := fmt.Sprintf("/api/%s", id)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, url, nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)

		actual := helpers.UnmarshalBody[todo_store.Todo](t, response.Body.Bytes())

		helpers.AssertEqual(t, actual.Description, "Item 1")
	})

	t.Run("GET /api/{id} returns 404 if item does not exist", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")

		server := api_server.NewApiServer(&todoStore, timeout)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/api/xyz", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusNotFound)
	})

	t.Run("GET request times out after set time", func(t *testing.T) {
		spyStore := helpers.SpyStore{}
		server := api_server.NewApiServer(&spyStore, timeout)

		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/api", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusRequestTimeout)
	})
}
