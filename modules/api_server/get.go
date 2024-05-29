package api_server

import (
	"context"
	"encoding/json"
	"net/http"
)

func (a *ApiServer) getAllItems(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	ok := make(chan []byte)
	fail := make(chan int)

	go func() {
		defer close(ok)
		defer close(fail)

		items := a.store.GetItems()

		if len(items) == 0 {
			fail <- http.StatusNotFound
			return
		}

		responseJson, err := json.Marshal(items)
		if err != nil {
			fail <- http.StatusInternalServerError
			return
		}

		ok <- responseJson
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case responseJson := <-ok:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseJson))
	case status := <-fail:
		w.WriteHeader(status)
	}
}

func (a *ApiServer) getItem(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	id := r.PathValue("id")

	a.sendItemAsResponse(ctx, w, id, http.StatusOK)
}

type ItemStatus struct {
	Completed bool `json:"completed"`
}

func (a *ApiServer) getItemStatus(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	ok := make(chan []byte)
	fail := make(chan int)

	go func() {
		defer close(ok)
		defer close(fail)

		item, err := a.store.Read(r.PathValue("id"))
		if err != nil {
			fail <- http.StatusNotFound
			return
		}

		response := ItemStatus{item.Complete}
		responseJson, err := json.Marshal(response)
		if err != nil {
			fail <- http.StatusInternalServerError
			return
		}

		ok <- responseJson
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case responseJson := <-ok:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responseJson))
	case status := <-fail:
		w.WriteHeader(status)
	}
}
