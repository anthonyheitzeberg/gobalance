package main

import (
	"gobalance/internal/balancer"
	"log"
	"net/http"
)

func main() {
	backends := []*balancer.Backend{
		{URL: "http://localhost:8081", Alive: true},
		{URL: "http://localhost:8082", Alive: true},
	}

	balancer := balancer.NewRoundRobinBalancer(backends)

	// Set up a handler that uses the balancer to forward requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		balancer.ForwardRequest(w, r)
	})

	// Start the load balancer server
    port := ":8080"
    log.Printf("GoBalance is running on port %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}