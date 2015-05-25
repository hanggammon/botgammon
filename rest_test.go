package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

/*
 * Adapted from:
 * http://www.markjberger.com/testing-web-apps-in-golang/
 *
 * Given a method type ("GET", "POST", etc) and parameters, serve the
 * response against the handler and return the ResponseRecorder.
 */
func test(t *testing.T, method string, urlStr string, params url.Values) *httptest.ResponseRecorder {
	r, err := http.NewRequest(
		method,
		urlStr,
		strings.NewReader(params.Encode()),
	)
	if err != nil {
		t.Errorf("%v", err)
	}
	r.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded; param=value",
	)
	w := httptest.NewRecorder()
	MuxRouter.ServeHTTP(w, r)
	return w
}

func TestV1CreateGame(t *testing.T) {
	w := test(t, "POST", "/api/v1/games", url.Values{})
	if w.Code != http.StatusCreated {
		t.Errorf("CreateGame API; expected %v, got %v", http.StatusCreated, w.Code)
	}
}

func TestV1ResetGame(t *testing.T) {
	w := test(t, "PATCH", "/api/v1/games/0/reset", url.Values{})
	if w.Code != http.StatusOK {
		t.Errorf("ResetGame API on existing game; expected %v, got %v", http.StatusOK, w.Code)
	}

	w = test(t, "PATCH", "/api/v1/games/0/reset/", url.Values{})
	if w.Code != http.StatusOK {
		t.Errorf("ResetGame API on existing game with URL ending in slash: expected %v, got %v", http.StatusOK, w.Code)
	}

	w = test(t, "PATCH", "/api/v1/games/10/reset", url.Values{})
	if w.Code != http.StatusNotFound {
		t.Errorf("ResetGame API on non-existent game: expected %v, got %v", http.StatusNotFound, w.Code)
	}
}
