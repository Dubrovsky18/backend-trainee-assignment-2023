package initializers

import "github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/build"

// InitializeBuildInfo creates new build.Info
func InitializeBuildInfo() *build.Info {
	return build.NewInfo()
}
