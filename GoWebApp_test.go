package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {

}

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
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
	}
}

/*
type HandleTester func(
    method string,
    params url.Values,
) *httptest.ResponseRecorder

func GenerateHandleTester(t *testing.T,handleFunc http.Handler,) HandleTester {

    // Given a method type ("GET", "POST", etc) and
    // parameters, serve the response against the handler and
    // return the ResponseRecorder.

    return func(
        method string,
        params url.Values,
    ) *httptest.ResponseRecorder {

        req, err := http.NewRequest(
            method,
            "",
            strings.NewReader(params.Encode()),
        )
        if err != nil {
            t.Errorf("%v", err)
        }
        req.Header.Set(
            "Content-Type",
            "application/x-www-form-urlencoded; param=value",
        )
        w := httptest.NewRecorder()
        handleFunc.ServeHTTP(w, req)
        return w
	}
}
function IndexTest(t *testing.T) {
    mockDb := MockDb{}
    homeHandle := homeHandler(mockDb)
    test := GenerateHandleTester(t, homeHandle)
    w := test("GET", url.Values{})
    if w.Code != http.StatusOK {
        t.Errorf("Home page didn't return %v", http.StatusOK)
    }
}
*/
