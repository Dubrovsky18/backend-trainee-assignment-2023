package initializers_test

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/config"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/initializers"
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/web/router"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("HttpServer", func() {
	Describe("InitializeHTTPServer()", func() {
		var (
			r   *gin.Engine
			cfg *config.HTTPConfig
		)

		BeforeEach(func() {
			r = router.NewRouter()
			cfg = initializers.InitializeAppConfig().HTTP
		})

		It("should initialize HTTP server", func() {
			srv := initializers.InitializeHTTPServer(r, cfg)

			Expect(srv).NotTo(BeNil())
		})
	})
})
