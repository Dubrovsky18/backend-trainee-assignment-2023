package apiv1

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestApiv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Apiv1 Suite")
}
