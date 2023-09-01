package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
)

func main() {
    http.HandleFunc("/analisis", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hola! este es un Microservicio, te encuentras solicitando el path: %q", r.URL.Path)
    })

    http.HandleFunc("/colors", func(w http.ResponseWriter, r *http.Request) {

        data, err := ioutil.ReadFile("colors.json")
        if err != nil {
            http.Error(w, "Error al leer el archivo JSON", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.WriteHeader(http.StatusOK)

        w.Write(data)

    })

    log.Fatal(http.ListenAndServe(":8081", nil))
}