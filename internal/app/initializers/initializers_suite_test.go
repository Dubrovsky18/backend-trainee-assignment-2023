package initializers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestInitializers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Initializers Suite")
}
