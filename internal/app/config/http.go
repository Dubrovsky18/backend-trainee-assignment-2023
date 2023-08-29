package config

import (
	"github.com/gobuffalo/envy"
)

// HTTPConfig
type HTTPConfig struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}

// DefaultHTTPConfig
func DefaultHTTPConfig() *HTTPConfig {
	return &HTTPConfig{
		Host: envy.Get("HTTP_HOST", "0.0.0.0"),
		Port: envy.Get("HTTP_PORT", "4443"),
	}
}
