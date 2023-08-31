package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/analisis", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hola, esto es un microservicio!, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/sistemas", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Usando microservicio con Go y Docker!")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}