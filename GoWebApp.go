package main

import (
	"html/template"
	"net/http"
)

type WebData struct {
	Title string
	Page  string
}

//HomeHandler - Bereitstellen der home.html
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("layout.html", "home.html")
	wd := WebData{
		Title: "WebApp",
		Page:  "Home",
	}
	tmpl.Execute(w, &wd)
}

//PageHandler -Bereitstellen der page.html
func PageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("layout.html", "page.html")
	wd := WebData{
		Title: "WebApp",
		Page:  "Page",
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
