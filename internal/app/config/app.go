package config

var config *AppConfig

type AppConfig struct {
	HTTP   *HTTPConfig   `json:"http"`
	Logger *LoggerConfig `json:"logger"`
}

// DefaultAppConfig
func NewAppConfig() *AppConfig {
	if config != nil {
		return config
	}

	config = &AppConfig{
		HTTP:   DefaultHTTPConfig(),
		Logger: DefaultLoggerConfig(),
	}

	return config
}
