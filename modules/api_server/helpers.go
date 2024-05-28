package api_server

import (
	"encoding/json"
	"io"
	"net/http"
)

func (a *ApiServer) sendItemAsResponse(w http.ResponseWriter, id string) {
	item, err := a.store.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	responseJson, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(responseJson))
}

func unpackRequest(w http.ResponseWriter, r *http.Request) RequestBody {
	defer r.Body.Close()

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var contents RequestBody
	if err := json.Unmarshal(bytes, &contents); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	return contents
}
