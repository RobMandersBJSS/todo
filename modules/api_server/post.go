package api_server

import "net/http"

func (a *ApiServer) postItem(w http.ResponseWriter, r *http.Request) {
	request := unpackRequest(w, r)

	id, err := a.store.Create(request.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	a.sendItemAsResponse(w, id, http.StatusCreated)
}
