package mappers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	m "kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
)

var _ = Describe("LookupMapper", func() {
	mapper := m.LookupMapper{}
	testID := 1
	testCodeType := "TEST"
	testCodeGroup := "TEST_GROUP"
	testName := "test name"
	testCode := s.Code{
		Base: s.Base{
			ID: testID,
		},
		CodeTypeID: testCodeType,
		CodeGroup:  testCodeGroup,
		Name:       testName,
	}

	Describe("LookupMapper", func() {
		It("MapCodeToLookup", func() {
			result := mapper.MapCodeToLookup(testCode)

			Expect(result.Id).To(Equal(testID))
			Expect(result.Text).To(Equal(testName))
			Expect(result.Code).To(Equal(testCodeType))
		})
	})
})
