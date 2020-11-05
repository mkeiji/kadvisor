package mappers_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	m "kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
)

var _ = Describe("ForecastMapper", func() {
	mapper := m.ForecastMapper{}
	forecast := generateTestForecast()

	Describe("MapForecastToMonthReportDto", func() {
		It("should map a forecast to forecast dto", func() {
			result := mapper.MapForecastToMonthReportDto(forecast)

			Expect(len(result)).To(Equal(12))
			for i, resultEntry := range result {
				forecastEntry := forecast.Entries[i]
				expectedBalance := forecastEntry.Income + forecastEntry.Expense

				Expect(resultEntry.Year).To(Equal(forecast.Year))
				Expect(resultEntry.Month).To(Equal(forecastEntry.Month))
				Expect(resultEntry.Income).To(Equal(forecastEntry.Income))
				Expect(resultEntry.Expense).To(Equal(forecastEntry.Expense))
				Expect(resultEntry.Balance).To(Equal(expectedBalance))
			}
		})
	})
})

func generateTestForecast() s.Forecast {
	fEntries := []s.ForecastEntry{}
	forecast := s.Forecast{
		UserID: 1,
		Year:   time.Now().Year(),
	}

	for i := 1; i < 13; i++ {
		nEntry := s.ForecastEntry{
			Month:   i,
			Income:  float64(i),
			Expense: float64(i),
		}
		fEntries = append(fEntries, nEntry)
	}

	forecast.Entries = fEntries
	return forecast
}
