package main

import (
	"fmt"
	"net/http"
)

const portname = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is about page and 2 + 2 = %d", sum)

}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application n port %s", portname)
	http.ListenAndServe(portname, nil)
}
