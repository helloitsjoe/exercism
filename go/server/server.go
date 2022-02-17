package main

import (
	"fmt"
	"log"
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
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/test", test)
	http.HandleFunc("/headers", headers)

	fmt.Println("Listening on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
