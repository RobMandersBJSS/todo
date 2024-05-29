package api_server

import (
	"context"
	"net/http"
)

func (a *ApiServer) postItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	channel := make(chan string)

	go func() {
		if r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		request := unpackRequest(w, r)

		id, err := a.store.Create(request.Description)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		channel <- id
		close(channel)
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case id := <-channel:
		a.sendItemAsResponse(ctx, w, id, http.StatusCreated)
	}

}
