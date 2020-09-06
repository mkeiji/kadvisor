package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type ForecastEntryService struct {
	repository repository.ForecastEntryRepository
}

func (svc *ForecastEntryService) Put(
	entry structs.ForecastEntry) (structs.ForecastEntry, error) {
	return svc.repository.Update(entry)
}
