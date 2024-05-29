package api_server

import (
	"context"
	"net/http"
)

func (a *ApiServer) patchItemDescription(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	ok := make(chan string)
	fail := make(chan int)

	go func() {
		defer close(ok)
		defer close(fail)

		if r.Body == nil {
			fail <- http.StatusBadRequest
			return
		}

		request := unpackRequest(w, r)

		err := a.store.UpdateItem(request.ID, request.Description)
		if err != nil {
			fail <- http.StatusNotFound
			return
		}

		ok <- request.ID
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case id := <-ok:
		a.sendItemAsResponse(ctx, w, id, http.StatusOK)
	case status := <-fail:
		w.WriteHeader(status)
	}
}

func (a *ApiServer) patchItemStatus(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	ok := make(chan string)
	fail := make(chan int)

	go func() {
		defer close(ok)
		defer close(fail)

		if r.Body == nil {
			fail <- http.StatusBadRequest
			return
		}

		request := unpackRequest(w, r)

		err := a.store.ToggleItemStatus(request.ID)
		if err != nil {
			fail <- http.StatusNotFound
			return
		}

		ok <- request.ID
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case id := <-ok:
		a.sendItemAsResponse(ctx, w, id, http.StatusOK)
	case status := <-fail:
		w.WriteHeader(status)
	}
}
