package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	utilhttp "github.com/mateusz-uminski/go-nethttp-healthz/util/http"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	t.Run("should register an endpoint", func(t *testing.T) {
		r := utilhttp.NewRouter()
		path := "/test"
		response := "example"
		handler := testHandler(response)

		// when
		r.RegisterEndpoint(path, handler)
		rr := sendRequest(r, http.MethodGet, path)

		// then
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, response, rr.Body.String())
	})

	t.Run("should register many endpoints", func(t *testing.T) {
		r := utilhttp.NewRouter()
		endpoints := []utilhttp.Endpoint{
			{
				Path:    "/test",
				Handler: testHandler("example"),
			},
			{
				Path:    "/example",
				Handler: testHandler("example"),
			},
		}

		// when
		r.RegisterEndpoints("/api", endpoints)

		// then
		for _, endpoint := range endpoints {
			rr := sendRequest(r, http.MethodGet, "/api"+endpoint.Path)

			assert.Equal(t, http.StatusOK, rr.Code)
			assert.Equal(t, "example", rr.Body.String())
		}
	})

	t.Run("should return 404 if endpoint is not registered", func(t *testing.T) {
		r := utilhttp.NewRouter()
		path := "/test"

		// when
		rr := sendRequest(r, http.MethodGet, path)

		// then
		expectedResponse := "404 page not found\n"
		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, expectedResponse, rr.Body.String())
	})
}

func testHandler(response string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response)) // nolint:errcheck
	}
}

func sendRequest(router utilhttp.Router, method string, path string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, nil)
	rr := httptest.NewRecorder()
	router.ServeMux().ServeHTTP(rr, req)
	return rr
}
