package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler )
	
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
// handler security
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "asan display neto", http.StatusNotFound )
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not Supported", http.StatusNotFound )
		return
	}

	fmt.Fprintf(w, "Hello!")
	
}