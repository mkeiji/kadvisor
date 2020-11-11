package interfaces

import (
	"kadvisor/server/libs/dtos"
)

//go:generate mockgen -destination=mocks/mock_report_repository.go -package=mocks . ReportRepository
type ReportRepository interface {
	GetAvailableForecastYears(userID int) ([]int, error)
	GetAvailableYears(userID int) ([]int, error)
	FindBalance(userID int) (dtos.Balance, error)
	FindYearToDateReport(userID int, year int) ([]dtos.MonthReport, error)
}
