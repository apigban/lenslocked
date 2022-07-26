package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>WELCOME</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact</h1><p>email: <a href=\"mailto:apigban@gmail.com\">apigban@gmail.com</a>.")
}

// func notFoundHandler(w http.ResponseWriter, r *http.Request) {

// 	w.WriteHeader(http.N)
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// }

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("starting the server on :3000")
	http.ListenAndServe(":3000", nil)
}
