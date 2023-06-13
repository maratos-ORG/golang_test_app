package backend

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	// Start the pinger function as a goroutine
	go pinger()

	req, err := http.NewRequest("GET", "/?ip=127.0.0.1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pingHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is non-empty
	if rr.Body.String() == "" {
		t.Errorf("handler returned empty body")
	}
}