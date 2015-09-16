package skeletor_test

import (
	. "github.com/crowdriff/skeletor/skeletor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conf", func() {

	BeforeEach(func() {
		SaveConf(&Conf{})
	})

	Context("SaveConf", func() {
		It("should be able to save and retrieve a Conf struct", func() {
			expected := Conf{
				Prod: true,
			}
			SaveConf(&expected)

			actual := GetConf()

			Ω(expected.Prod).Should(Equal(actual.Prod))
		})
	})

})
