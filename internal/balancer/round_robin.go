package balancer

import (
	"gobalance/internal/shared"
	"net/http"
	"sync"
)

// RoundRobinBalancer implements the round-robin load balancing algorithm
type RoundRobinBalancer struct {
	backends []*shared.Backend
	current int
	mutex sync.Mutex
}

// NewRoundRobinBalancer initializes a new RoundRobinBalancer
func NewRoundRobinBalancer(backends []*shared.Backend) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		backends: backends,
		current: 0,
	}
}

// NextBackend returns the next backend to which traffic should be forwarded
func (b *RoundRobinBalancer) NextBackend() *shared.Backend {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	backend := b.backends[b.current]

	// Update to the next backend using the round-robin approach
	b.current = (b.current + 1) % len(b.backends)

	return backend
}

// ForwardRequest forwards the HTTP request to the selected backend
func (b *RoundRobinBalancer) ForwardRequest(w http.ResponseWriter, r *http.Request) {
	backend := b.NextBackend()

	// Create the proxy request to the backend
	proxyReq, err := http.NewRequest(r.Method, backend.URL, r.Body)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Forward all headers
	proxyReq.Header = r.Header

	// Send the request to the backend
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Error forwarding request to backend", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// Write the backend's response to the client
	w.WriteHeader(resp.StatusCode)
	if _,err := w.Write([]byte(backend.URL + " responded.")); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}