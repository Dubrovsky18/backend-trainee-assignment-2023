package config

import "github.com/gobuffalo/envy"

// LoggerConfig
type LoggerConfig struct {
	Namespace string `json:"namespace,omitempty"`
	BuildMode string `json:"build_mode,omitempty"`
}

func DefaultLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		Namespace: envy.Get("LOG_NAMESPACE", "fabric-svc"),
		BuildMode: envy.Get("LOG_MODE", "development"),
	}
}
