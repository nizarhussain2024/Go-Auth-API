package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "Go Auth API Running")
    })
    fmt.Println("Server started on :8080")
    http.ListenAndServe(":8080", nil)
}
