package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	called := false

	handler := func (w http.ResponseWriter, r *http.Request) {
		called = true

		params := r.Context().Value("params").(map[string]string)

		name, ok := params["name"]

		if !ok {
			t.Error("Params map does not contain `name`")
		}

		if name != "aldehir" {
			t.Error("Name is not `aldehir`")
		}
	}

	router := CreateRouter()
	router.HandleFunc("^/hello/(?P<name>[^/]+)/?$", handler)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost/hello/aldehir", nil)

	router.ServeHTTP(recorder, request)

	if !called {
		t.Error("Handler function not called")
	}
}
