package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GoBalance Load Balancer!")
}

func main() {
	http.HandleFunc("/", handler)

	port := ":8080"
	log.Printf("GoBalance is running on port %s...\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}