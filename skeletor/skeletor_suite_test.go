package skeletor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSkeletor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Skeletor Suite")
}
