package mappers_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	m "kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	c "kadvisor/server/resources/constants"
)

var _ = Describe("EntryMapper", func() {
	var mapper m.EntryMapper

	incomeEntry := s.Entry{}

	BeforeEach(func() {
		mapper = m.EntryMapper{}
	})

	Describe("MapEntry", func() {
		Context("Map date", func() {
			It("should return date in utc", func() {
				date := time.Now()
				estLocation, _ := time.LoadLocation("EST")
				utcLocation, _ := time.LoadLocation("UTC")
				wrongDate := date.In(estLocation)
				expectedDate := date.In(utcLocation)

				incomeEntry.Amount = 5
				incomeEntry.Date = wrongDate
				result := mapper.MapEntry(incomeEntry)

				Expect(result.Date).To(Equal(expectedDate))
			})
		})

		Context("Map amount", func() {
			It("should return a positive amount for type income", func() {
				expectedAmount := float64(5)

				incomeEntry.Amount = -5
				incomeEntry.EntryTypeCodeID = c.INCOME_ENTRY_TYPE

				result := mapper.MapEntry(incomeEntry)
				Expect(result.Amount).To(Equal(expectedAmount))
			})

			It("should return a negative amount for type expense", func() {
				expectedAmount := float64(-5)

				incomeEntry.Amount = 5
				incomeEntry.EntryTypeCodeID = c.EXPENSE_ENTRY_TYPE

				result := mapper.MapEntry(incomeEntry)
				Expect(result.Amount).To(Equal(expectedAmount))
			})
		})
	})
})
