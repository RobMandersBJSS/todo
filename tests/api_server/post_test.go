package api_server_test

import (
	"bytes"
	"net/http"
	"testing"
	"todo/modules/api_server"
	"todo/modules/store"
	"todo/tests/helpers"
)

func TestApiServerPOST(t *testing.T) {
	t.Run("POST /api/ returns 201 and creates a new item", func(t *testing.T) {
		todoStore := store.TodoStore{}
		server := api_server.NewApiServer(&todoStore)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte("{\"description\":\"New Item\"}"))
		request, response := helpers.NewRequestResponse(t, http.MethodPost, "/api/", &requestBody)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusCreated)
		helpers.AssertEqual(t, len(todoStore.GetItems()), 1)
		helpers.AssertEqual(t, todoStore.GetItems()[0].Description, "New Item")
	})

	t.Run("POST /api/ returns the new item", func(t *testing.T) {
		todoStore := store.TodoStore{}
		server := api_server.NewApiServer(&todoStore)

		requestBody := bytes.Buffer{}
		requestBody.Write([]byte("{\"description\":\"New Item\"}"))
		request, response := helpers.NewRequestResponse(t, http.MethodPost, "/api/", &requestBody)

		server.ServeHTTP(response, request)

		actual := helpers.UnmarshalBody[store.Todo](t, response.Body.Bytes())

		helpers.AssertEqual(t, actual.Description, "New Item")
	})
}
