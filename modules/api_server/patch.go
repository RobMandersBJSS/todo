package api_server

import (
	"context"
	"net/http"
)

func (a *ApiServer) patchItemDescription(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	channel := make(chan string)

	go func() {
		if r.Body == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		request := unpackRequest(w, r)

		err := a.store.UpdateItem(request.ID, request.Description)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		channel <- request.ID
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case id := <-channel:
		a.sendItemAsResponse(ctx, w, id, http.StatusOK)
	}

}

func (a *ApiServer) patchItemStatus(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := unpackRequest(w, r)

	err := a.store.ToggleItemStatus(request.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	default:
		a.sendItemAsResponse(ctx, w, request.ID, http.StatusOK)
	}
}
