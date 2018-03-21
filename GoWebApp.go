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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
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
	http.HandleFunc("/home", HomeHandler)
	http.HandleFunc("/page", pageHandler)
	http.HandleFunc("/", HomeHandler)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.ListenAndServe(":80", nil)
}
