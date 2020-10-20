package services

import (
	"kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/mappers"
	"net/http"
	"sort"
)

type ReportService struct {
	repository         repository.ReportRepository
	forecastRepository repository.ForecastRepository
	forecastMapper     mappers.ForecastMapper
}

func (svc *ReportService) GetBalance(
	userID int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	balance, err := svc.repository.FindBalance(userID)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, balance)
	}

	return response
}

func (svc *ReportService) GetYearToDateReport(
	userID int,
	year int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	report, err := svc.repository.FindYearToDateReport(userID, year)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, report)
	}

	return response
}

func (svc *ReportService) GetYearToDateWithForecastReport(
	userID int,
	year int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	errors := []error{}
	ytdMonths, ytdErr := svc.repository.FindYearToDateReport(userID, year)
	forecast, fcErr := svc.forecastRepository.FindOne(userID, year, true)
	forecastMonts := svc.forecastMapper.MapForecastToMonthReportDto(forecast)

	if ytdErr != nil && fcErr != nil {
		errors = append(errors, ytdErr)
		errors = append(errors, fcErr)
	}
	if len(errors) > 0 {
		response = dtos.NewKresponse(
			http.StatusNotFound,
			KeiGenUtil.MapErrList(errors),
		)
	} else {
		combined := svc.combineYtdWithForecast(ytdMonths, forecastMonts)
		response = dtos.NewKresponse(
			http.StatusOK,
			svc.getCombinedWithAccumulatedBalance(combined),
		)
	}

	return response
}

func (svc *ReportService) GetReportAvailable(userID int) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	result, err := svc.repository.GetAvailableYears(userID)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
		response = dtos.NewKresponse(http.StatusOK, result)
	}

	return response
}

func (svc *ReportService) getCombinedWithAccumulatedBalance(
	combinedMonths []dtos.MonthReport,
) []dtos.MonthReport {
	updated := []dtos.MonthReport{}

	accBalance := 0.0
	for _, month := range combinedMonths {
		accBalance = accBalance + month.Balance
		month.Balance = accBalance
		updated = append(updated, month)
	}

	return updated
}

func (svc *ReportService) combineYtdWithForecast(
	ytdMonths []dtos.MonthReport,
	forecastMonts []dtos.MonthReport,
) []dtos.MonthReport {

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
