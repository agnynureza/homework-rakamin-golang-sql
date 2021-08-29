package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	handler TokenHandler
)

func TestTokenHandler_GetNewAccessToken(t *testing.T) {
	app.Get("/login", handler.GetNewAccessToken)

	t.Run("should return success", func(t *testing.T) {
		url := "/login"
		req := httptest.NewRequest("GET", url, nil)
		resp, err := app.Test(req)
		require.Nil(t, err)
		require.Equal(t, resp.StatusCode, http.StatusOK)
	})

	t.Run("should return error", func(t *testing.T) {
		url := "/salah-url"
		req := httptest.NewRequest("GET", url, nil)
		resp, err := app.Test(req)
		require.Nil(t, err)
		require.Equal(t, resp.StatusCode, http.StatusNotFound)
	})
}
