package main

import (
	"html/template"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Panicln(err)
	}
	t.Execute(w, "Hello World Top Page")
}

func about(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("about.html")
	if err != nil {
		log.Panicln(err)
	}
	t.Execute(w, "Hello World About Page")
}

func main() {
	http.HandleFunc("/top", top)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}
