package api_server

import "net/http"

type requestBody struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func (a *ApiServer) patchItemDescription(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	err := a.store.UpdateItem(request.ID, request.Description)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.sendItemAsResponse(w, request.ID, http.StatusOK)
}

func (a *ApiServer) patchItemStatus(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	err := a.store.ToggleItemStatus(request.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.sendItemAsResponse(w, request.ID, http.StatusOK)
}
