package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl_index := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Value := r.FormValue("formulaire")
			fmt.Println(Value)
		}

		tmpl_index.Execute(w, nil)
	})

	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	http.ListenAndServe(":8080", nil)
}
