package main

import (
	"fmt"
	"goo/pkg/handlers"
	"net/http"
)

const portname = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portname)
	http.ListenAndServe(portname, nil)
}
