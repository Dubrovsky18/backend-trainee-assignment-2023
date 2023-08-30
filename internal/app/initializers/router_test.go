package initializers_test

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/initializers"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/config"
	"github.com/gobuffalo/envy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/dependencies"
)

var _ = Describe("Router", func() {
	Describe("InitializeRouter()", func() {
		var (
			c *dependencies.Container
		)

		BeforeEach(func() {
			envy.Set("CONFIG_PATH", "../../../configs")

			c = &dependencies.Container{
				Configuration: config.GetConfig(),
			}
		})

		It("should initialize router", func() {
			r := initializers.InitializeRouter(c)

			Expect(r).NotTo(BeNil())
		})
	})
})
