package api_server

import (
	"context"
	"encoding/json"
	"net/http"
)

const OnlineMessage = "Server Online"

type StatusMessage struct {
	Message string `json:"message"`
}

func (a *ApiServer) getServerStatus(w http.ResponseWriter, r *http.Request) {
	ctx, cancelCtx := context.WithTimeout(r.Context(), a.timeout)
	defer cancelCtx()

	ok := make(chan []byte)
	fail := make(chan int)

	go func() {
		defer close(ok)
		defer close(fail)

		responseJson, err := json.Marshal(StatusMessage{OnlineMessage})
		if err != nil {
			fail <- http.StatusInternalServerError
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
