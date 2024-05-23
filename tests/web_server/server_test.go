package server_test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"todo/modules/store"
	"todo/modules/web_server"
	"todo/tests/helpers"

	approvals "github.com/approvals/go-approval-tests"
	websocket "github.com/gorilla/websocket"
)

func TestServer(t *testing.T) {
	var dummyStore = store.TodoStore{}

	t.Run("GET / returns 200", func(t *testing.T) {
		server := web_server.NewServer(createTemplate(t), &dummyStore)

		request, err := http.NewRequest(http.MethodGet, "/", nil)
		helpers.AssertNoError(t, err)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)
	})

	t.Run("GET / returns a body from a given template", func(t *testing.T) {
		server := web_server.NewServer(createTemplate(t), &dummyStore)

		request, err := http.NewRequest(http.MethodGet, "/", nil)
		helpers.AssertNoError(t, err)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		approvals.VerifyString(t, response.Body.String())
	})

	t.Run("Add new item via websocket request", func(t *testing.T) {
		todoStore := store.TodoStore{}
		server := createTestServer(t, &todoStore)

		ws := createWebSocket(t, server)
		defer ws.Close()

		message := "{\"action\":\"new\",\"description\":\"Todo Item\"}"
		writeWebSocketMessage(t, ws, message)

		time.Sleep(10 * time.Millisecond)

		actual := todoStore.Items[0].Description
		expected := "Todo Item"

		helpers.AssertEqual(t, actual, expected)
	})

	t.Run("Update item via websocket request", func(t *testing.T) {
		todoStore := store.TodoStore{
			Items: []store.Todo{
				{ID: "0", Description: "Todo Item", Complete: false},
			},
		}
		server := createTestServer(t, &todoStore)

		ws := createWebSocket(t, server)
		defer ws.Close()

		message := "{\"action\":\"update\",\"id\":\"0\",\"description\":\"Updated Item\"}"
		writeWebSocketMessage(t, ws, message)

		time.Sleep(10 * time.Millisecond)

		actual, err := todoStore.Read("0")
		helpers.AssertNoError(t, err)

		expected := store.Todo{ID: "0", Description: "Updated Item", Complete: false}

		helpers.AssertStructEqual(t, actual, expected)
	})

	t.Run("Toggle item status via websocket request", func(t *testing.T) {
		todoStore := store.TodoStore{
			Items: []store.Todo{
				{ID: "0", Description: "Todo Item", Complete: false},
			},
		}
		server := createTestServer(t, &todoStore)

		ws := createWebSocket(t, server)
		defer ws.Close()

		message := "{\"action\":\"toggle\",\"id\":\"0\"}"
		writeWebSocketMessage(t, ws, message)

		time.Sleep(10 * time.Millisecond)

		actual, err := todoStore.Read("0")
		helpers.AssertNoError(t, err)

		expected := store.Todo{ID: "0", Description: "Todo Item", Complete: true}

		helpers.AssertStructEqual(t, actual, expected)
	})

	t.Run("Delete item via websocket request", func(t *testing.T) {
		todoStore := store.TodoStore{
			Items: []store.Todo{
				{ID: "0", Description: "Item 1", Complete: false},
				{ID: "1", Description: "Item 2", Complete: false},
			},
		}
		server := createTestServer(t, &todoStore)

		ws := createWebSocket(t, server)
		defer ws.Close()

		message := "{\"action\":\"delete\",\"id\":\"0\"}"
		writeWebSocketMessage(t, ws, message)

		time.Sleep(10 * time.Millisecond)

		actual := todoStore.Items
		expected := []store.Todo{{ID: "1", Description: "Item 2", Complete: false}}

		helpers.AssertStructEqual(t, actual, expected)
	})
}

func writeWebSocketMessage(t testing.TB, ws *websocket.Conn, message string) {
	t.Helper()

	err := ws.WriteMessage(websocket.TextMessage, []byte(message))
	helpers.AssertNoError(t, err)
}

func createWebSocket(t *testing.T, server *httptest.Server) *websocket.Conn {
	t.Helper()

	webSocketURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

	ws, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	helpers.AssertNoError(t, err)

	return ws
}

func createTestServer(t testing.TB, todoStore *store.TodoStore) *httptest.Server {
	t.Helper()

	template := createTemplate(t)

	return httptest.NewServer(web_server.NewServer(template, todoStore))
}

func createTemplate(t testing.TB) *template.Template {
	t.Helper()

	test_template, err := template.New("main.gohtml").Parse("<h1>Test Template</h1>")
	helpers.AssertNoError(t, err)

	return test_template
}
