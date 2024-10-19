package shared

// Backend represents a backend server configuration.
type Backend struct {
    URL   string `json:"url"`
    Alive bool   `json:"alive"`
}