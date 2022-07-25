package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>WELCOME TO MY SITE</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting the server on :3000")
	http.ListenAndServe(":3000", nil)
}
