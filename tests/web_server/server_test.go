package server_test

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
	"todo/modules/store"
	"todo/modules/web_app"
	"todo/tests/helpers"

	approvals "github.com/approvals/go-approval-tests"
	websocket "github.com/gorilla/websocket"
)

func TestServer(t *testing.T) {
	var dummyStore = store.TodoStore{}

	t.Run("GET / returns 200", func(t *testing.T) {
		server := web_app.NewServer(createTemplate(t), &dummyStore)

		request, err := http.NewRequest(http.MethodGet, "/", nil)
		helpers.AssertNoError(t, err)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)
	})

	t.Run("GET / returns a HTML document with listed todo tasks", func(t *testing.T) {
		server := web_app.NewServer(createTemplate(t), &dummyStore)

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

		actual := todoStore.GetItems()[0].Description
		expected := "Todo Item"

		helpers.AssertEqual(t, actual, expected)
	})

	t.Run("Update item via websocket request", func(t *testing.T) {
		todoStore := store.TodoStore{}
		todoStore.Create("Todo Item")
		id := todoStore.GetItems()[0].ID

		server := createTestServer(t, &todoStore)

		ws := createWebSocket(t, server)
		defer ws.Close()

		message := fmt.Sprintf("{\"action\":\"update\",\"id\":\"%s\",\"description\":\"Updated Item\"}", id)
		writeWebSocketMessage(t, ws, message)

		time.Sleep(10 * time.Millisecond)

		actual, err := todoStore.Read(id)
		helpers.AssertNoError(t, err)

		helpers.AssertStructEqual(t, actual.Description, "Updated Item")
	})

	t.Run("Toggle item status via websocket request", func(t *testing.T) {
		todoStore := store.TodoStore{}
		todoStore.Create("Todo Item")
		id := todoStore.GetItems()[0].ID

		server := createTestServer(t, &todoStore)

		ws := createWebSocket(t, server)
		defer ws.Close()

		message := fmt.Sprintf("{\"action\":\"toggle\",\"id\":\"%s\"}", id)
		writeWebSocketMessage(t, ws, message)

		time.Sleep(10 * time.Millisecond)

		actual, err := todoStore.Read(id)
		helpers.AssertNoError(t, err)

		helpers.AssertStructEqual(t, actual.Complete, true)
	})

	t.Run("Delete item via websocket request", func(t *testing.T) {
		todoStore := store.TodoStore{}
		todoStore.Create("Item 1")
		todoStore.Create("Item 2")
		id := todoStore.GetItems()[0].ID

		server := createTestServer(t, &todoStore)

		ws := createWebSocket(t, server)
		defer ws.Close()

		message := fmt.Sprintf("{\"action\":\"delete\",\"id\":\"%s\"}", id)
		writeWebSocketMessage(t, ws, message)

		time.Sleep(10 * time.Millisecond)

		actual := todoStore.GetItems()

		helpers.AssertStructEqual(t, len(actual), 1)
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

	return httptest.NewServer(web_app.NewServer(template, todoStore))
}

func createTemplate(t testing.TB) *template.Template {
	t.Helper()

	test_template, err := template.New("main.gohtml").Parse("<h1>Test Template</h1>")
	helpers.AssertNoError(t, err)

	return test_template
}
