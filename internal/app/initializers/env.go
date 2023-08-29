package initializers

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/pkg/logger"
	"github.com/gobuffalo/envy"
)

func InitializeEnv() {
	if err := envy.Load(); err != nil {
		logger.Info("can not load .env file", "error", err)
		envy.Reload()
	}
}
