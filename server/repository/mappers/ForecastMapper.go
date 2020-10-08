package mappers

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
)

type ForecastMapper struct{}

func (f *ForecastMapper) MapForecastToMonthReportDto(
	forecast structs.Forecast) []dtos.MonthReport {

	mapped := []dtos.MonthReport{}

	year := forecast.Year
	for _, entry := range forecast.Entries {
		mapped = append(mapped, dtos.MonthReport{
			Year:    year,
			Month:   entry.Month,
			Income:  entry.Income,
			Expense: entry.Expense,
			Balance: entry.Income + entry.Expense,
		})
	}

	return mapped
}
