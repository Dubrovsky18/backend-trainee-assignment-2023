package app_test

import (
	"context"
	app2 "github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Application", func() {
	Describe("InitializeApplication()", func() {
		It("should create new Application", func() {
			app, _ := app2.InitializeApplication()

			Expect(app).NotTo(BeNil())
		})
	})

	Describe("methods", func() {
		var (
			app *app2.Application
		)

		BeforeEach(func() {
			app, _ = app2.InitializeApplication()
		})

		Describe("Start(), Stop()", func() {
			It("should start and stop application", func() {
				ctx, cancel := context.WithCancel(context.Background())
				app.Start(ctx, false)

				defer cancel()

				err := app.Stop()

				Expect(err).To(BeNil())
			})
		})
	})
})
