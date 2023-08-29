package dependencies

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/build"
	services "github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/services"
)

// Container is a DI container for application
type Container struct {
	BuildInfo *build.Info
	Service   services.TemplateServiceUserSlug
}
