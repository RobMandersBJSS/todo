package api_server

import (
	"net/http"
	"time"
	"todo/modules/todo_store"
)

type ApiServer struct {
	http.Handler
	store todo_store.TodoStore
}

func NewApiServer(store todo_store.TodoStore, timout time.Duration) *ApiServer {
	server := new(ApiServer)

	router := http.NewServeMux()
	router.Handle("GET /api", http.HandlerFunc(server.getAllItems))
	router.Handle("GET /api/{id}", http.HandlerFunc(server.getItem))
	router.Handle("POST /api", http.HandlerFunc(server.postItem))
	router.Handle("PATCH /api/update", http.HandlerFunc(server.patchItemDescription))
	router.Handle("PATCH /api/toggle", http.HandlerFunc(server.patchItemStatus))
	router.Handle("DELETE /api", http.HandlerFunc(server.deleteItem))

	server.Handler = router
	server.store = store

	return server
}
