package build_test

import (
	"github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/build"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Info", func() {
	Describe("NewInfo()", func() {
		It("should create new info object", func() {
			info := build.NewInfo()

			Expect(info).NotTo(BeNil())
		})
	})
})
