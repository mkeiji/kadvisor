package services

import (
	"kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/mappers"
	"kadvisor/server/resources/enums"
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
	latestActual := svc.getLatestActualMonth(combinedMonths)

	accBalance := 0.0
	for _, month := range combinedMonths {
		if (latestActual == dtos.MonthReport{}) ||
			(month.Type == enums.ACTUAL.ToString()) ||
			(month.Month >= latestActual.Month) {
			accBalance = accBalance + month.Balance
		} else {
			month.Income = 0.0
			month.Expense = 0.0
		}
		month.Balance = accBalance
		updated = append(updated, month)
	}

	return updated
}

func (svc *ReportService) combineYtdWithForecast(
	ytdMonths []dtos.MonthReport,
	forecastMonths []dtos.MonthReport,
) []dtos.MonthReport {

	result := []dtos.MonthReport{}
	for i := 1; i <= 12; i++ {
		aMonth := svc.findMonthInMonthReport(i, ytdMonths)
		fMonth := svc.findMonthInMonthReport(i, forecastMonths)
		if svc.isEmptyMonthReport(aMonth) {
			fMonth.Type = enums.FORECAST.ToString()
			result = append(result, fMonth)
		} else {
			aMonth.Type = enums.ACTUAL.ToString()
			result = append(result, aMonth)
		}
	}
	return result
}

func (svc *ReportService) findMonthInMonthReport(
	month int,
	monthReports []dtos.MonthReport,
) dtos.MonthReport {

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

func (svc *ReportService) getLatestActualMonth(
	monthReports []dtos.MonthReport,
) dtos.MonthReport {
	var result dtos.MonthReport

	for _, month := range monthReports {
		if month.Type == enums.ACTUAL.ToString() {
			result = month
		}
	}

	return result
}
