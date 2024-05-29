package api_server

import (
	"context"
	"net/http"
)

func (a *ApiServer) deleteItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	channel := make(chan int)

	go func() {
		if r.Body == nil {
			channel <- http.StatusBadRequest
			return
		}

		request := unpackRequest(w, r)

		err := a.store.Delete(request.ID)
		if err != nil {
			channel <- http.StatusNotFound
			return
		}

		channel <- http.StatusOK
		close(channel)
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case status := <-channel:
		w.WriteHeader(status)
	}
}
