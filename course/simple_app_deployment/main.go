package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hello world!\n")
}

func main() {
	http.HandleFunc("/hello", getRoot)
	http.ListenAndServe(":3333", nil)
}
