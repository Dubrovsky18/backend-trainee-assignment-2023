package initializers

import "github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/config"

func InitializeAppConfig() *config.AppConfig {
	return config.NewAppConfig()
}
