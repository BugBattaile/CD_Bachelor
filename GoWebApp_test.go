package main

import (
	"net/http"
	"testing"
	"os/exec"
)

func TestMain(t *testing.T) {

}

func TestUrlHome(t *testing.T) {
	// create server
	cmd := exec.Command("go", "run", "GoWebApp.go")
    cmd.Start()

	// test url home
	req, err := http.Get("http://localhost/home")
	cmd.Wait()
	status := req.StatusCode

	if err != nil{
		t.Fatal(err)
	} 

	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}


func TestServer(t *testing.T) {
	// create server
	cmd := exec.Command("go", "run", "GoWebApp.go")
    cmd.Start()

	// test url home
	req, err := http.Get("http://localhost/randomURLwichisntincludedsoweget404")
	cmd.Wait()
	status := req.StatusCode

	if err != nil{
		t.Fatal(err)
	} 

	if status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}