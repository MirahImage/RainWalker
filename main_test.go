package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MirahImage/RainWalker"
	soda "github.com/SebastiaanKlippert/go-soda"
)

var _ = Describe("Main", func() {
	var (
		shedReq       *soda.GetRequest
		shedsEndpoint string
		appToken      string
	)

	BeforeSuite(func() {
		shedsEndpoint = "https://data.cityofnewyork.us/resource/2jy7-cddj"
		appToken = "Hl8EuekONWAZmQj7osuAIKPTg"
		shedReq = soda.NewGetRequest(shedsEndpoint, appToken)
	})

	Describe("NumberSheds", func() {
		var (
			numSheds uint
			err      error
		)

		BeforeEach(func() {
			numSheds, err = NumberSheds(shedReq)
		})

		It("should not produce an error", func() {
			Expect(err).To(BeNil())
		})

		It("should return the number of sheds in NY", func() {
			Expect(numSheds).To(BeNumerically(">", 0))
		})
	})

	Describe("IsCovered", func() {
		var (
			isCovered   bool
			err         error
			borough     string
			fakeBorough string
			streetName  string
			houseNumber string
		)

		BeforeEach(func() {
			borough = "BRONX"
			fakeBorough = "JERSEY"
			streetName = "HULL AVENUE"
			houseNumber = "3232"
			isCovered, err = IsCovered(shedReq, borough, streetName, houseNumber)
		})

		It("should not produce an error", func() {
			Expect(err).To(BeNil())
		})

		It("should return true when there is a shed at the address", func() {
			Expect(isCovered).To(BeTrue())
		})

		It("should return false when there is not a shed at the address", func() {
			isCovered, err = IsCovered(shedReq, fakeBorough, streetName, houseNumber)
			Expect(isCovered).To(BeFalse())
			Expect(err).To(BeNil())
		})

	})
})
