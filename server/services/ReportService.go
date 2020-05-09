package services

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
)

type ReportService struct {
	repository 	repository.ReportRepository
}

func (svc *ReportService) GetBalance(
	userID int) (dtos.Balance, error) {
	return svc.repository.FindBalance(userID)
}

func (svc *ReportService) GetYearToDateReport(
	userID int) ([]dtos.MonthReport, error) {
	return svc.repository.FindYearToDateReport(userID)
}