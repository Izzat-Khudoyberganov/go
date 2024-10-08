package main

import (
	"fmt"
	"net/http"
)

const portname = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", portname)
	http.ListenAndServe(portname, nil)
}
