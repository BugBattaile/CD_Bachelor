package main

import (
	"html/template"
	"io"
	"net/http"
)

type WebData struct {
	Title string
	Page  string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("layout.html", "home.html")
	wd := WebData{
		Title: "WebApp",
		Page:  "Home",
	}
	tmpl.Execute(w, &wd)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("layout.html", "page.html")
	wd := WebData{
		Title: "WebApp",
		Page:  "Page",
	}
	tmpl.Execute(w, &wd)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/page", pageHandler)
	http.ListenAndServe(":80", nil)
}

// HealthCheckHandler as a simple HTTP Health check
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
	//	w.Header().Set("Content-Type", "application/json")
}
