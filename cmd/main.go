package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
    "log"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
    }).Methods("GET")

    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
