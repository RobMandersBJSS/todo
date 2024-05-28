package api_server_test

import (
	"net/http"
	"testing"
	"todo/modules/api_server"
	"todo/modules/store"
	"todo/tests/helpers"
)

func TestApiServerGET(t *testing.T) {
	t.Run("GET /api/ returns a 200 response", func(t *testing.T) {
		todoStore := store.TodoStore{}
		server := api_server.NewApiServer(&todoStore)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/api/", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)
	})

	t.Run("GET /api/ returns a list of todos as JSON", func(t *testing.T) {
		todoStore := store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")

		server := api_server.NewApiServer(&todoStore)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/api/", nil)

		server.ServeHTTP(response, request)

		actual := helpers.UnmarshalBody[[]store.Todo](t, response.Body.Bytes())

		helpers.AssertEqual(t, len(actual), 2)
		helpers.AssertEqual(t, actual[0].Description, "Item 1")
		helpers.AssertEqual(t, actual[1].Description, "Item 2")
	})
}
