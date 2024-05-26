package errors_test

import (
	"dashboard/internal/errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIError(t *testing.T) {
	err := errors.APIError{
		StatusCode: 400,
		Message: "Requisição inválida",
	}

	assert.Equal(t, "Requisição inválida", err.Error())

	expectedJSON := `{
	  "statusCode": 400,
	  "message": "Requisição inválida"
	}`
	jsonResponse, _ := err.ToJSON()
	assert.JSONEq(t, expectedJSON, string(jsonResponse))

	w := &mockResponseWriter{}
	err.RespondWithError(w, err)
	assert.Equal(t, http.StatusBadRequest, w.statusCode)
	assert.JSONEq(t, expectedJSON, w.response)
}

type mockResponseWriter struct {
	statusCode int
	response   string
}

func (w *mockResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (w *mockResponseWriter) Write(data []byte) (int, error) {
	w.response = string(data)
	return len(data), nil
}

func (w *mockResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}