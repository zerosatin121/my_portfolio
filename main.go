package main

import (
    "net/http"
    "html/template"
)

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, nil)
    })

    println("Server running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
