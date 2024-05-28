package api_server_test

import (
	"fmt"
	"net/http"
	"testing"
	"todo/modules/api_server"
	"todo/modules/store"
	"todo/tests/helpers"
)

func TestApiServerGET(t *testing.T) {
	t.Run("GET /api returns a list of todos as JSON", func(t *testing.T) {
		todoStore := store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")

		server := api_server.NewApiServer(&todoStore)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/api", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)

		actual := helpers.UnmarshalBody[[]store.Todo](t, response.Body.Bytes())

		helpers.AssertEqual(t, len(actual), 2)
		helpers.AssertEqual(t, actual[0].Description, "Item 1")
		helpers.AssertEqual(t, actual[1].Description, "Item 2")
	})

	t.Run("GET /api returns 404 if there are no items", func(t *testing.T) {
		todoStore := store.TodoStore{}
		server := api_server.NewApiServer(&todoStore)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/api", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusNotFound)
	})

	t.Run("GET /api/{id} returns the specified item", func(t *testing.T) {
		todoStore := store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")
		id := todoStore.GetItems()[0].ID

		server := api_server.NewApiServer(&todoStore)
		url := fmt.Sprintf("/api/%s", id)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, url, nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)

		actual := helpers.UnmarshalBody[store.Todo](t, response.Body.Bytes())

		helpers.AssertEqual(t, actual.Description, "Item 1")
	})

	t.Run("GET /api/{id} returns 404 if item does not exist", func(t *testing.T) {
		todoStore := store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")

		server := api_server.NewApiServer(&todoStore)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/api/xyz", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusNotFound)
	})
}
