package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/RainWalker"
)

var _ = Describe("Main", func() {
	Describe("NumberSheds", func() {
		var (
			numSheds uint
			err      error
		)

		BeforeEach(func() {
			numSheds, err = NumberSheds()
		})

		It("should not produce an error", func() {
			Expect(err).To(BeNil())
		})

		It("should return the number of sheds in NY", func() {
			Expect(numSheds).To(BeNumerically(">", 0))
		})
	})
})
