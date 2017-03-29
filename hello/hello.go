package main

import (
	"fmt"
	"github.com/asgrim/go-playground/string"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	reverse := r.URL.Query().Get("reverse")

	if reverse == "" {
		fmt.Fprint(w, "You need to provide the 'reverse' query parameter...")
		return
	}

	fmt.Fprintf(w, "Hey: %s", string.Reverse(reverse))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
