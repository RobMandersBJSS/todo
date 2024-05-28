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
	router.Handle("GET /api", http.HandlerFunc(server.getAllItems))
	router.Handle("GET /api/{id}", http.HandlerFunc(server.getItem))
	router.Handle("POST /api", http.HandlerFunc(server.postItem))
	router.Handle("PATCH /api/update", http.HandlerFunc(server.updateItemDescription))
	router.Handle("PATCH /api/toggle", http.HandlerFunc(server.toggleItemStatus))
	router.Handle("DELETE /api", http.HandlerFunc(server.deleteItem))

	server.Handler = router
	server.store = store

	return server
}

func (a *ApiServer) getAllItems(w http.ResponseWriter, r *http.Request) {
	items := a.store.GetItems()
	if len(items) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	responseJson, err := json.Marshal(items)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJson))
}

func (a *ApiServer) getItem(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	a.sendItemAsResponse(w, id, http.StatusOK)
}

func (a *ApiServer) postItem(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	id, err := a.store.Create(request.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	a.sendItemAsResponse(w, id, http.StatusCreated)
}

func (a *ApiServer) updateItemDescription(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	err := a.store.UpdateItem(request.ID, request.Description)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.sendItemAsResponse(w, request.ID, http.StatusOK)
}

func (a *ApiServer) toggleItemStatus(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	err := a.store.ToggleItemStatus(request.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.sendItemAsResponse(w, request.ID, http.StatusOK)
}

func (a *ApiServer) deleteItem(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	err := a.store.Delete(request.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
