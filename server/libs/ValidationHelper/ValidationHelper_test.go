package ValidationHelper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/ValidationHelper"
)

var _ = Describe("ValidationHelper", func() {
	const (
		property      = "Test.Property"
		validationMsg = "Test msg"
	)
	var expectedMsg string

	BeforeEach(func() {
		expectedMsg = "Key: 'Test.Property' Error:Field validation for 'Property' Test msg"
	})

	Describe("GetValidationMsg", func() {
		Context("one level property", func() {
			It("should return message", func() {
				result := ValidationHelper.GetValidationMsg(
					property,
					validationMsg,
				)

				Expect(result).To(Equal(expectedMsg))
			})
		})
	})
})
