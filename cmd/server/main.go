package main

import (
    "fmt"
    "net/http"

    "github.com/a-h/templ"
    "github.com/CaffeineOrDeath/banner-gen/internal/views"
)

func main(){
    component := views.Layout("Home")
    http.Handle("/", templ.Handler(component))

    fmt.Println("Listening on :42069")
    http.ListenAndServe(":42069", nil) //iykyk
}
