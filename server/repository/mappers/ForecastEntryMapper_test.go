package mappers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	m "kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
)

var _ = Describe("ForecastEntryMapper", func() {
	var mapper m.ForecastEntryMapper

	entry := s.ForecastEntry{}

	BeforeEach(func() {
		mapper = m.ForecastEntryMapper{}
	})

	Describe("MapForecastEntry", func() {
		It("should return income positive and expense negative", func() {
			expectedIncome := float64(5)
			expectedExpense := float64(-5)

			entry.Income = -5
			entry.Expense = 5
			result := mapper.MapForecastEntry(entry)

			Expect(result.Income).To(Equal(expectedIncome))
			Expect(result.Expense).To(Equal(expectedExpense))
		})
	})
})
