
# GoBalance - Simple Load Balancer in Go 🐹⚖️

**GoBalance** is a simple and minimal load balancer written in Go. It implements the **round-robin** load balancing algorithm to distribute traffic across multiple backend servers. This is a learning project to understand the basics of load balancing, concurrency, and HTTP request forwarding in Go.

## Features
- **Round-robin load balancing**: Requests are distributed evenly across backend servers in a circular fashion.
- **Concurrency-safe with mutex**: Ensures that multiple requests are handled correctly using Go’s concurrency primitives.

## Definitions
**[Round robin](https://www.cloudflare.com/learning/performance/types-of-load-balancing-algorithms/#:~:text=Round%20robin%3A%20Round%20robin%20load,response%20to%20each%20DNS%20query.)**: Round robin load balancing distributes traffic to a list of servers in rotation using the Domain Name System (DNS). An authoritative nameserver will have a list of different A records for a domain and provides a different one in response to each DNS query. Source: 

## How It Works
The load balancer forwards incoming HTTP requests to a pool of backend servers, distributing the traffic evenly by sending each request to the next backend in line. It uses a mutex to make sure that the backend selection process is concurrency-safe.

## Getting Started

### Prerequisites
- Go 1.18 or higher
- Backend servers (can be simple Python HTTP servers or your own backend)

### Project Structure
```
gobalance/
│
├── main.go
├── go.mod
└── internal/
    └── balancer/
        └── round_robin.go
```

### Running the Load Balancer

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/gobalance.git
   ```

2. Go into the project directory:
   ```bash
   cd gobalance
   ```

3. Start two simple backend servers (for testing) using Python:
   ```bash
   python3 -m http.server 8081
   python3 -m http.server 8082
   ```

4. Run the load balancer:
   ```bash
   go run main.go
   ```

5. Visit `http://localhost:8080` in your browser. The load balancer will forward your requests to `http://localhost:8081` and `http://localhost:8082` alternately using the round-robin algorithm.

## Next Steps
- **Health checks**: Add functionality to automatically check the status of backend servers and ensure that traffic is only sent to healthy servers.
- **Additional algorithms**: Implement other load balancing algorithms like least-connections or weighted round-robin.

## License
This project is open-source and available under the MIT License.
