package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func UnmarshalBody[T any](t testing.TB, body []byte) T {
	t.Helper()

	var response T
	err := json.Unmarshal(body, &response)
	AssertNoError(t, err)

	return response
}

func NewRequestResponse(t testing.TB, method, url string, body io.Reader) (*http.Request, *httptest.ResponseRecorder) {
	t.Helper()

	request, err := http.NewRequest(method, url, body)
	AssertNoError(t, err)

	response := httptest.NewRecorder()

	return request, response
}
