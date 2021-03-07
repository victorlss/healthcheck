package checker

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CreateHttpTestServer(statusCode int) *httptest.Server {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(statusCode)
		}),
	)

	return server
}

func TestIsAlive(t *testing.T) {
	t.Run("should return true when status is OK", func(t *testing.T) {
		server := CreateHttpTestServer(http.StatusOK)
		defer server.Close()

		got := IsAlive(server.URL)

		assert.True(t, got)
	})

	t.Run("should return false when status code is InternalServerError", func(t *testing.T) {
		server := CreateHttpTestServer(http.StatusInternalServerError)
		defer server.Close()

		got := IsAlive(server.URL)

		assert.False(t, got)
	})
}
