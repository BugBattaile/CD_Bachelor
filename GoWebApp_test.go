package main

import (
	"net/http"
	"testing"
)

func TestMain(t *testing.T) {

}

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	//req, err := http.NewRequest("GET", "/health-check", nil)
	req, err := http.Get("http://localhost/home")
	if err != nil {
		t.Fatal(err)
	} else {
		print(string(req.StatusCode) + req.Status)
	}
	/*
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HealthCheckHandler)
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := `{"alive": true}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}*/
}
