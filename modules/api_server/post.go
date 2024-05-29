package api_server

import (
	"context"
	"net/http"
)

func (a *ApiServer) postItem(w http.ResponseWriter, r *http.Request) {
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

		id, err := a.store.Create(request.Description)
		if err != nil {
			fail <- http.StatusInternalServerError
			return
		}

		ok <- id
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case id := <-ok:
		a.sendItemAsResponse(ctx, w, id, http.StatusCreated)
	case status := <-fail:
		w.WriteHeader(status)
	}
}
