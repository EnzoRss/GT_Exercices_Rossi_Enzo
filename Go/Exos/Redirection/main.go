package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Mauvaise Route")
}

func redir(w http.ResponseWriter, r *http.Request) {
	tmpl_index := template.Must(template.ParseFiles("index.html"))
	tmpl_redir1 := template.Must(template.ParseFiles("redir1.html"))
	tmpl_redir2 := template.Must(template.ParseFiles("redir2.html"))
	tmpl_redir3 := template.Must(template.ParseFiles("redir3.html"))

	if r.URL.Path != "/index" || r.URL.Path != "/redir1" || r.URL.Path != "/redir2" || r.URL.Path != "/redir3" {

	}

	http.HandleFunc("/redir1", func(w http.ResponseWriter, r *http.Request) {
		tmpl_redir1.Execute(w, nil)
	})
	http.HandleFunc("/redir2", func(w http.ResponseWriter, r *http.Request) {
		tmpl_redir2.Execute(w, nil)
	})
	http.HandleFunc("/redir3", func(w http.ResponseWriter, r *http.Request) {
		tmpl_redir3.Execute(w, nil)
	})
	tmpl_index.Execute(w, nil)
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
}

func main() {

	http.HandleFunc("/index", redir)
	http.HandleFunc("/", NotFoundHandler)

	http.ListenAndServe(":8080", nil)
}
