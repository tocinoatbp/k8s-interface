package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GetPodsResponse struct {
	Message string `json:"message"`
}

type GetPodsRequest struct {
	// Structural flags
	Namespace string `json:"namespace"`
	PodName   string `json:"pod_name"`
}

func RegisterGetPods(client K8sClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if r.URL.Path != "/hello" {
		// 	http.Error(w, "", http.StatusNotFound)
		// 	return
		// }

		if r.Method != "POST" {
			http.Error(w, "Method is not Supported", http.StatusNotFound)
			return
		}

		body := r.Body
		defer body.Close()

		var request GetPodsRequest
		if err := json.NewDecoder(body).Decode(&request); err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		fmt.Printf("%+v", request)
		get, err := client.GetPods(request)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		bytes, err := json.Marshal(get)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, string(bytes))
	}

}
