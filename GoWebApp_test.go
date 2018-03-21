package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {

	req, err := http.NewRequest("Get", "localhost:80/", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}
	rec := httptest.NewRecorder()
	HomeHandler(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Got %v", res.Status)
	}
}

func TestComponents(t *testing.T) {
	reqHome, errHome := http.Get("http://localhost/home")
	if errHome != nil {
		t.Fatal(errHome)
	} else {
		if reqHome.StatusCode != http.StatusOK {
			t.Fatal(reqHome.Status)
		} else {
			println("Home:" + reqHome.Status)
		}
	}
	reqPage, errPage := http.Get("http://localhost/page")
	if errPage != nil {
		t.Fatal(errPage)
	} else {
		if reqPage.StatusCode != http.StatusOK {
			t.Fatal(reqPage.Status)
		} else {
			println("Page:" + reqPage.Status)
		}
	}
	reqCSS, errCSS := http.Get("http://localhost/css/materialize.min.css")
	if errCSS != nil {
		t.Fatal(errCSS)
	} else {
		if reqCSS.StatusCode != http.StatusOK {
			t.Fatal(reqCSS.Status)
		} else {
			println("CSS:" + reqCSS.Status)
			println(http.StatusOK)
		}
	}
}
