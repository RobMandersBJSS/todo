package api_server_test

import (
	"net/http"
	"testing"
	"time"
	"todo/modules/api_server"
	"todo/modules/todo_memory_store"
	"todo/tests/helpers"
)

func TestApiServer(t *testing.T) {
	timeout := 10 * time.Millisecond

	t.Run("GET /status returns 200 status", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		server := api_server.NewApiServer(&todoStore, timeout)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/status", nil)

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)
	})

	t.Run("GET /status returns JSON message", func(t *testing.T) {
		todoStore := todo_memory_store.TodoStore{}
		server := api_server.NewApiServer(&todoStore, timeout)
		request, response := helpers.NewRequestResponse(t, http.MethodGet, "/status", nil)

		server.ServeHTTP(response, request)

		actual := helpers.UnmarshalBody[api_server.StatusMessage](t, response.Body.Bytes()).Message
		expected := api_server.OnlineMessage

		helpers.AssertEqual(t, response.Code, http.StatusOK)
		helpers.AssertEqual(t, actual, expected)
	})
}
