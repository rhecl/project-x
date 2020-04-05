package main

import (
	"fmt"
	"net/http"
)

func healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/", healthz)

	http.ListenAndServe(":3001", nil)
}
