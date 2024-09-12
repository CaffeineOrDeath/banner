package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/CaffeineOrDeath/banner-gen/internal/views"
)

func main(){
    component := views.Layout("Home")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        component.Render(r.Context(), w)
    })

    fmt.Println("Listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil)) // iykyk
}
