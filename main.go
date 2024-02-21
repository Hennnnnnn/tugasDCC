package main

import (
	"fmt"
	"net/http"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "page awal")
	case "/login":
		fmt.Fprint(w, "page login")
	case "/register":
		fmt.Fprint(w, "page register")
	default:
		fmt.Fprint(w, "page invalid")
	}
}

func main() {
	http.HandleFunc("/", helloWorldPage)
	http.ListenAndServe("", nil)
}
