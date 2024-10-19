package main

import (
	"gobalance/internal/balancer"
	"gobalance/internal/config"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	// Get the port from the environment variable, default to 8080 if not set
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port
    }

	// Load configuration
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("❌ Error loading config: %v", err)
	}
	log.Printf("✅ Loaded config with %d backends\n", len(config.Backends))

	// Initialize the round-robn balancer with the loaded backends
	balancer := balancer.NewRoundRobinBalancer(config.Backends)
	log.Println("✅ RoundRobinBalancer initialized")

	// Set up a handler that uses the balancer to forward requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		balancer.ForwardRequest(w, r)
		// Add request forwarding logic here
	})

    log.Printf("✅ GoBalance is running on port %s...\n", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}