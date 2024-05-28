package api_server

import (
	"context"
	"encoding/json"
	"net/http"
)

func (a *ApiServer) getAllItems(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	items := a.store.GetItems()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	default:
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
}

func (a *ApiServer) getItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	id := r.PathValue("id")

	a.sendItemAsResponse(ctx, w, id, http.StatusOK)
}
