package services

import (
	"kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
	m "kadvisor/server/repository/mappers"
	"kadvisor/server/resources/enums"
	"net/http"
	"sort"
)

type ReportService struct {
	Repository         i.ReportRepository
	ForecastRepository i.ForecastRepository
	ForecastMapper     m.ForecastMapper
}

func NewReportService() ReportService {
	return ReportService{
		Repository:         r.ReportRepository{},
		ForecastRepository: r.ForecastRepository{},
		ForecastMapper:     m.ForecastMapper{},
	}
}

func (svc ReportService) GetBalance(
	userID int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	balance, err := svc.Repository.FindBalance(userID)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, balance)
	}

	return response
}

func (svc ReportService) GetYearToDateReport(
	userID int,
	year int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	report, err := svc.Repository.FindYearToDateReport(userID, year)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, report)
	}

	return response
}

func (svc ReportService) GetYearToDateWithForecastReport(
	userID int,
	year int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse
	errors := []error{}

	ytdMonths, ytdErr := svc.Repository.FindYearToDateReport(userID, year)
	if ytdErr != nil {
		errors = append(errors, ytdErr)
	}

	forecast, fcErr := svc.ForecastRepository.FindOne(userID, year, true)
	if fcErr != nil {
		errors = append(errors, fcErr)
	}

	if len(errors) > 0 {
		response = dtos.NewKresponse(
			http.StatusNotFound,
			KeiGenUtil.MapErrList(errors),
		)
	} else {
		forecastMonths := svc.ForecastMapper.MapForecastToMonthReportDto(forecast)
		combined := svc.combineYtdWithForecast(ytdMonths, forecastMonths)
		response = dtos.NewKresponse(
			http.StatusOK,
			svc.getCombinedWithAccumulatedBalance(combined),
		)
	}

	return response
}

func (svc ReportService) GetReportForecastAvailable(userID int) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	result, err := svc.Repository.GetAvailableForecastYears(userID)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
		response = dtos.NewKresponse(http.StatusOK, result)
	}

	return response
}

func (svc ReportService) GetReportAvailable(userID int) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	result, err := svc.Repository.GetAvailableYears(userID)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
		response = dtos.NewKresponse(http.StatusOK, result)
	}

	return response
}

func (svc ReportService) getCombinedWithAccumulatedBalance(
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

func (svc ReportService) combineYtdWithForecast(
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

func (svc ReportService) findMonthInMonthReport(
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

func (svc ReportService) isEmptyMonthReport(month dtos.MonthReport) bool {
	return (dtos.MonthReport{}) == month
}

func (svc ReportService) getLatestActualMonth(
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
