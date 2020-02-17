package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	called := false

	handler := func(w http.ResponseWriter, r *http.Request) {
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

	router := NewRouter()
	router.HandleFunc("^/hello/(?P<name>[^/]+)/?$", handler)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost/hello/aldehir", nil)

	router.ServeHTTP(recorder, request)

	if !called {
		t.Error("Handler function not called")
	}
}

func TestRouterFallthrough(t *testing.T) {
	helloCalled := false
	goodbyeCalled := false
	noopCalled := false

	noopHandler := func(w http.ResponseWriter, r *http.Request) {
		noopCalled = true
	}

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		helloCalled = true
	}

	goodbyeHandler := func(w http.ResponseWriter, r *http.Request) {
		goodbyeCalled = true
	}

	router := NewRouter()
	router.HandleFunc("^/goodbye/(?P<name>[^/]+)/?$", goodbyeHandler)
	router.HandleFunc("^/hello/(?P<name>[^/]+)/?$", helloHandler)
	router.HandleFunc("^/(?P<name>[^/]+)/?$", noopHandler)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://localhost/hello/aldehir", nil)

	router.ServeHTTP(recorder, request)

	if !helloCalled {
		t.Error("Hello handler not called")
	}

	if noopCalled || goodbyeCalled {
		t.Error("Other handlers were called")
	}

	recorder = httptest.NewRecorder()
	request = httptest.NewRequest("GET", "http://localhost/goodbye/aldehir", nil)

	helloCalled = false

	router.ServeHTTP(recorder, request)

	if !goodbyeCalled {
		t.Error("Goodbye Handler not called")
	}

	if noopCalled || helloCalled {
		t.Error("Other handlers were called")
	}
}
