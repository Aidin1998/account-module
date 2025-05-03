package integration

import (
	"account-module/cmd"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	e := cmd.NewServer()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/healthz", nil)
	e.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Code)
	assert.JSONEq(t, `{"status":"OK"}`, rec.Body.String())
}
