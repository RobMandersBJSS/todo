package api_server

import (
	"encoding/json"
	"net/http"
	"todo/modules/store"
)

type RequestBody struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type ApiServer struct {
	http.Handler
	store *store.TodoStore
}

func NewApiServer(store *store.TodoStore) *ApiServer {
	server := new(ApiServer)

	router := http.NewServeMux()
	router.Handle("GET /api/", http.HandlerFunc(server.getItems))
	router.Handle("POST /api/", http.HandlerFunc(server.postItem))
	router.Handle("PATCH /api/update/", http.HandlerFunc(server.updateItemDescription))
	router.Handle("PATCH /api/toggle/", http.HandlerFunc(server.toggleItemStatus))

	server.Handler = router
	server.store = store

	return server
}

func (a *ApiServer) getItems(w http.ResponseWriter, r *http.Request) {
	responseJson, err := json.Marshal(a.store.GetItems())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJson))
}

func (a *ApiServer) postItem(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	id, err := a.store.Create(request.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	a.sendItemAsResponse(w, id)
}

func (a *ApiServer) updateItemDescription(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	err := a.store.UpdateItem(request.ID, request.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	a.sendItemAsResponse(w, request.ID)
}

func (a *ApiServer) toggleItemStatus(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	err := a.store.ToggleItemStatus(request.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	a.sendItemAsResponse(w, request.ID)
}
