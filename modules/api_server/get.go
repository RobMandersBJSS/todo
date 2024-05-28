package api_server

import (
	"encoding/json"
	"net/http"
)

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
