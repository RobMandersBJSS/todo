package api_server

import (
	"context"
	"net/http"
)

func (a *ApiServer) deleteItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := unpackRequest(w, r)

	err := a.store.Delete(request.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	default:
		w.WriteHeader(http.StatusOK)
	}
}
