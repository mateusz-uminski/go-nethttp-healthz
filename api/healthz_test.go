package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mateusz-uminski/go-nethttp-healthz/api"
	"github.com/mateusz-uminski/go-nethttp-healthz/internal/fake"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	t.Run("should return status healthy", func(t *testing.T) {
		handler := api.Healthz(fake.MakeLogger())
		req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
		rr := httptest.NewRecorder()

		// when
		handler(rr, req)

		// then
		expectedResponse := api.HealthzResponse{
			Status: "healthy",
		}
		actualResponse, err := actualResponse(rr.Body)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, expectedResponse, actualResponse)
	})
}

func actualResponse(body io.Reader) (api.HealthzResponse, error) {
	var response api.HealthzResponse
	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return api.HealthzResponse{}, err
	}
	return response, nil
}
