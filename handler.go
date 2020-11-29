package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "", http.StatusNotFound )
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not Supported", http.StatusNotFound )
		return
	}

	fmt.Fprintf(w, "Hello this is your first webserver!")
	
}