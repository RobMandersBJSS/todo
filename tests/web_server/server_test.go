package server_test

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/modules/web_server"
	"todo/tests/helpers"

	approvals "github.com/approvals/go-approval-tests"
)

func TestServer(t *testing.T) {
	t.Run("GET / returns 200", func(t *testing.T) {
		server := web_server.NewServer(createTemplate(t))

		request, err := http.NewRequest(http.MethodGet, "/", nil)
		helpers.AssertNoError(t, err)

		response := httptest.NewRecorder()

		server.Handle(response, request)

		helpers.AssertEqual(t, response.Code, http.StatusOK)
	})

	t.Run("GET / returns a body from a given template", func(t *testing.T) {
		server := web_server.NewServer(createTemplate(t))

		request, err := http.NewRequest(http.MethodGet, "/", nil)
		helpers.AssertNoError(t, err)

		response := httptest.NewRecorder()

		server.Handle(response, request)

		approvals.VerifyString(t, response.Body.String())
	})
}

func createTemplate(t testing.TB) *template.Template {
	t.Helper()

	test_template, err := template.New("main.gohtml").Parse("<h1>Test Template</h1>")
	helpers.AssertNoError(t, err)

	return test_template
}
