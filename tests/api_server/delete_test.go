package api_server_test

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
	"todo/modules/api_server"
	"todo/modules/store"
	"todo/tests/helpers"
)

func TestApiServerDELETE(t *testing.T) {
	t.Run("DELETE /api/ deleted an item", func(t *testing.T) {
		todoStore := store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")
		id := todoStore.GetItems()[0].ID

		server := api_server.NewApiServer(&todoStore)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte(fmt.Sprintf("{\"id\":\"%s\"}", id)))
		request, response := helpers.NewRequestResponse(t, http.MethodDelete, "/api", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)

		helpers.AssertEqual(t, len(todoStore.GetItems()), 1)
		helpers.AssertEqual(t, todoStore.GetItems()[0].Description, "Item 2")
	})

	t.Run("DELETE /api returns 404 if item does not exist", func(t *testing.T) {
		todoStore := store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")

		server := api_server.NewApiServer(&todoStore)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte("{\"id\":\"xyz\"}"))
		request, response := helpers.NewRequestResponse(t, http.MethodDelete, "/api", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusNotFound)
	})
}
