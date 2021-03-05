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

	entry1 := s.Entry{}
	entry2 := s.Entry{}

	BeforeEach(func() {
		mapper = m.EntryMapper{}
	})

	Describe("MapEntry", func() {
		Context("Map amount", func() {
			It("should return a positive amount for type income", func() {
				expectedAmount := float64(5)

				entry1.Amount = -5
				entry1.EntryTypeCodeID = c.INCOME_ENTRY_TYPE

				result := mapper.MapEntry(entry1)
				Expect(result.Amount).To(Equal(expectedAmount))
			})

			It("should return a negative amount for type expense", func() {
				expectedAmount := float64(-5)

				entry1.Amount = 5
				entry1.EntryTypeCodeID = c.EXPENSE_ENTRY_TYPE

				result := mapper.MapEntry(entry1)
				Expect(result.Amount).To(Equal(expectedAmount))
			})
		})
	})

	Describe("MapEntryDate", func() {
		It("should return date in utc with ISO 8601 (RFC 3339) format for a single obj", func() {
			date := time.Now()
			estLocation, _ := time.LoadLocation("EST")
			utcLocation, _ := time.LoadLocation("UTC")
			wrongDate := date.In(estLocation)
			expectedDate, _ := time.Parse(time.RFC3339, date.In(utcLocation).Format(time.RFC3339))

			entry1.Amount = 5
			entry1.Date = wrongDate
			result := mapper.MapEntry(entry1)

			Expect(result.Date).To(Equal(expectedDate))
		})
	})

	Describe("MapEntriesDates", func() {
		It("should return empty array if list is empty", func() {
			entryList := []s.Entry{}

			result := mapper.MapEntriesDates(entryList)
			Expect(result).To(BeEmpty())
		})

		It("should map dates for multiple objs", func() {
			date := time.Now()
			estLocation, _ := time.LoadLocation("EST")
			utcLocation, _ := time.LoadLocation("UTC")
			wrongDate := date.In(estLocation)
			expectedDate, _ := time.Parse(time.RFC3339, date.In(utcLocation).Format(time.RFC3339))

			entry1.Date = wrongDate
			entry2.Date = wrongDate
			entryList := []s.Entry{entry1, entry2}

			result := mapper.MapEntriesDates(entryList)
			Expect(result[0].Date).To(Equal(expectedDate))
			Expect(result[1].Date).To(Equal(expectedDate))
		})
	})
})
