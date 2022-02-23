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

func formHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse error: %v", err)
		return
	}
	fmt.Fprint(w, "Handling form...\n")
	name := req.FormValue("name")
	email := req.FormValue("email")

	fmt.Fprint(w, "Name = %s\n", name)
	fmt.Fprint(w, "Email = %s\n", email)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/test", test)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/headers", headers)

	fmt.Println("Listening on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
