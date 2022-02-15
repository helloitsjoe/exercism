package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Testing...")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v/n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/headers", headers)

	fmt.Println("Listening on http://localhost:8080/test")

	http.ListenAndServe(":8080", nil)
}
