package main

import (
	"html/template"
	"net/http"
)

// WebData - Statische Meldungen
type WebData struct {
	Title   string
	Page    string
	Content string
}

//HomeHandler - Bereitstellen der home.html
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("layout.html", "home.html")
	wd := WebData{
		Title:   "WebApp",
		Page:    "Home",
		Content: "This is a simple Webapp written in Go.",
	}
	tmpl.Execute(w, &wd)
}

<<<<<<< HEAD
//PageHandler - Bereitstellen der page.html
func pageHandler(w http.ResponseWriter, r *http.Request) {
=======
//PageHandler -Bereitstellen der page.html
func PageHandler(w http.ResponseWriter, r *http.Request) {
>>>>>>> development
	tmpl, _ := template.ParseFiles("layout.html", "page.html")
	wd := WebData{
		Title:   "WebApp",
		Page:    "Page",
		Content: "It has been deployed throw a Continuous Delivery Pipeline in a Kubernetes Cluster.",
	}
	tmpl.Execute(w, &wd)
}

func main() {
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/page", PageHandler)
	http.HandleFunc("/", HomeHandler)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.ListenAndServe(":80", nil)
}
