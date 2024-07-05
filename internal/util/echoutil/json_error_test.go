package echoutil

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestJsonError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Call JsonError with an error and a message
	err := JsonError(c, http.StatusBadRequest, errors.New("test error"), "test message")

	// Assert that no error was returned from JsonError
	assert.NoError(t, err)

	// Assert that the response status is 400
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Assert that the response body is as expected
	expectedBody := `{"error":"test error","message":"test message"}`
	assert.Equal(t, expectedBody, strings.TrimSuffix(rec.Body.String(), "\n"))

	// Call JsonError with an error and no message
	rec.Body.Reset() // Reset the response body
	err = JsonError(c, http.StatusBadRequest, errors.New("test error"))

	// Assert that no error was returned from JsonError
	assert.NoError(t, err)

	// Assert that the response status is 400
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	// Assert that the response body is as expected
	expectedBody = `{"error":"test error"}` + "\n"
	assert.Equal(t, expectedBody, rec.Body.String())
}
