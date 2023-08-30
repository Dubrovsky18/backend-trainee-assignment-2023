package router

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Router", func() {
	Describe("NewRouter()", func() {
		It("should create new router", func() {
			r := NewRouter()

			Expect(r).NotTo(BeNil())
		})
	})
})
