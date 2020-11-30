package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// add route for static files
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// add route to endpoint
	http.HandleFunc("/getPods", RegisterGetPods(NewK8sClient()))

	// Run webserver
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
