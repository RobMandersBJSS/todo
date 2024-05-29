package api_server

import (
	"context"
	"encoding/json"
	"net/http"
	"todo/modules/todo_store"
)

func (a *ApiServer) getAllItems(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	channel := make(chan []todo_store.Todo)

	go func() {
		channel <- a.store.GetItems()
		close(channel)
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case items := <-channel:
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
