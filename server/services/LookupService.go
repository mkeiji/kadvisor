package services

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/mappers"
)

type LookupService struct {
	mapper	 	mappers.LookupMapper
	repository 	repository.CodeCodeTextRepository
}

func (svc *LookupService) GetAllByCodeGroup (
	codeGroup string) ([]dtos.LookupEntry, error) {

	var lookups []dtos.LookupEntry

	codes, err := svc.repository.FindAllByCodeGroup(codeGroup)
	if err == nil {
		for _, c := range codes {
			lookups = append(lookups, svc.mapper.MapCodeToLookup(c))
		}
	}

	return lookups, err
}
