package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/hello", nil)
	rec := httptest.NewRecorder()

	req.Header.Set("Content-Type", "application/json")

	// Act
	Hello(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"message\":\"hello\"}", rec.Body.String())
}
