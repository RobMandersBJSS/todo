package api_server

import (
	"context"
	"net/http"
)

func (a *ApiServer) deleteItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	channel := make(chan bool)

	go func() {
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

		channel <- true
		close(channel)
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case <-channel:
		w.WriteHeader(http.StatusOK)
	}
}
