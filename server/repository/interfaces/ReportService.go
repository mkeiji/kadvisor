package interfaces

import (
	"kadvisor/server/libs/dtos"
)

//go:generate mockgen -destination=mocks/mock_report_service.go -package=mocks . ReportService
type ReportService interface {
	GetBalance(userID int) dtos.KhttpResponse
	GetYearToDateReport(userID int, year int) dtos.KhttpResponse
	GetYearToDateWithForecastReport(userID int, year int) dtos.KhttpResponse
	GetReportForecastAvailable(userID int) dtos.KhttpResponse
	GetReportAvailable(userID int) dtos.KhttpResponse
}
