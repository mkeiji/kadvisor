package services

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/mappers"
	"sort"
)

type ReportService struct {
	repository         repository.ReportRepository
	forecastRepository repository.ForecastRepository
	forecastMapper     mappers.ForecastMapper
}

func (svc *ReportService) GetBalance(
	userID int) (dtos.Balance, error) {
	return svc.repository.FindBalance(userID)
}

func (svc *ReportService) GetYearToDateReport(
	userID int, year int) ([]dtos.MonthReport, error) {
	return svc.repository.FindYearToDateReport(userID, year)
}

func (svc *ReportService) GetYearToDateWithForecastReport(
	userID int, year int) ([]dtos.MonthReport, []error) {

	var result []dtos.MonthReport
	errors := []error{}
	ytdMonths, ytdErr := svc.repository.FindYearToDateReport(userID, year)
	forecast, _ := svc.forecastRepository.FindOne(userID, year, true)
	forecastMonts := svc.forecastMapper.MapForecastToMonthReportDto(forecast)

	if ytdErr != nil {
		errors = append(errors, ytdErr)
	}

	result = svc.combineYtdWithForecast(ytdMonths, forecastMonts)
	return result, errors
}

func (svc *ReportService) GetReportAvailable(userID int) ([]int, error) {
	result, err := svc.repository.GetAvailableYears(userID)
	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result, err
}

func (svc *ReportService) combineYtdWithForecast(
	ytdMonths []dtos.MonthReport, forecastMonts []dtos.MonthReport) []dtos.MonthReport {

	result := []dtos.MonthReport{}
	for i := 1; i <= 12; i++ {
		month := svc.findMonthInMonthReport(i, ytdMonths)
		if svc.isEmptyMonthReport(month) {
			result = append(result, svc.findMonthInMonthReport(i, forecastMonts))
		} else {
			result = append(result, month)
		}
	}
	return result
}

func (svc *ReportService) findMonthInMonthReport(
	month int, monthReports []dtos.MonthReport) dtos.MonthReport {

	result := dtos.MonthReport{}

	for _, mr := range monthReports {
		if mr.Month == month {
			result = mr
		}
	}

	return result
}

func (svc *ReportService) isEmptyMonthReport(month dtos.MonthReport) bool {
	return (dtos.MonthReport{}) == month
}
