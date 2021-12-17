package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Server)
	http.ListenAndServe(":8888", nil)
}

func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World %s", r.URL.Path[1:])

	//fmt.Fprintln(nil)
}
