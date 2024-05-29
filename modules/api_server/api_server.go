package api_server

import (
	"net/http"
	"time"
	"todo/modules/todo_store"
)

type requestBody struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type ApiServer struct {
	store   todo_store.TodoStore
	timeout time.Duration
	http.Handler
}

func NewApiServer(store todo_store.TodoStore, timeout time.Duration) *ApiServer {
	server := new(ApiServer)

	router := http.NewServeMux()
	router.Handle("GET /status", http.HandlerFunc(server.getServerStatus))
	router.Handle("GET /api", http.HandlerFunc(server.getAllItems))
	router.Handle("GET /api/{id}", http.HandlerFunc(server.getItem))
	router.Handle("GET /api/status/{id}", http.HandlerFunc(server.getItemStatus))
	router.Handle("POST /api", http.HandlerFunc(server.postItem))
	router.Handle("PATCH /api/update", http.HandlerFunc(server.patchItemDescription))
	router.Handle("PATCH /api/toggle", http.HandlerFunc(server.patchItemStatus))
	router.Handle("DELETE /api", http.HandlerFunc(server.deleteItem))

	server.store = store
	server.timeout = timeout
	server.Handler = router

	return server
}
