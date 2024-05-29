package api_server

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func (a *ApiServer) sendItemAsResponse(ctx context.Context, w http.ResponseWriter, id string, status int) {
	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()

	channel := make(chan []byte)

	go func() {
		item, err := a.store.Read(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		responseJson, err := json.Marshal(item)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		channel <- responseJson
		close(channel)
	}()

	select {
	case <-ctx.Done():
		cancelCtx()
		w.WriteHeader(http.StatusRequestTimeout)
	case responseJson := <-channel:
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(responseJson))
	}
}

func unpackRequest(w http.ResponseWriter, r *http.Request) requestBody {
	defer r.Body.Close()

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var contents requestBody
	if err := json.Unmarshal(bytes, &contents); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return contents
}
