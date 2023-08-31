package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/analisis", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hola! este es un Microservicio, te encuentras solicitando el path: %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/sistemas", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Microservicio con Go y Docker By jsuhula!")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))
    
}