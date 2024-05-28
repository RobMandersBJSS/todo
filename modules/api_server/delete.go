package api_server

import "net/http"

func (a *ApiServer) deleteItem(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	err := a.store.Delete(request.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
