package config

import (
	"encoding/json"
	"gobalancer/internal/shared"
	"io/ioutil"
	"os"
)

// Config represents the entire configuration for the load balancer
type Config struct {
	Backends []*shared.Backend `json:"backends"`
}

// LoadConfig loads the configuration from the given filename.
func LoadConfig(filename string) (*Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    byteValue, err := ioutil.ReadAll(file)
    if err != nil {
        return nil, err
    }

	// Unmarshal the entire config into a struct
    var config Config
    if err := json.Unmarshal(byteValue, &config); err != nil {
        return nil, err
    }

    return &config, nil
}