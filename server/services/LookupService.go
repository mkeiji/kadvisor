package services

import (
	"errors"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/mappers"
	"net/http"
)

type LookupService struct {
	Mapper     mappers.LookupMapper
	Repository i.CodeCodeTextRepository
}

func NewLookupService() LookupService {
	return LookupService{
		Mapper:     mappers.LookupMapper{},
		Repository: repository.CodeCodeTextRepository{},
	}
}

func (svc LookupService) GetAllByCodeGroup(
	codeGroup string,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse
	var lookups []dtos.LookupEntry

	if codeGroup == "" {
		return dtos.NewKresponse(
			http.StatusBadRequest,
			errors.New("missing codeGroup param"),
		)
	}

	codes, err := svc.Repository.FindAllByCodeGroup(codeGroup)
	if err == nil {
		for _, c := range codes {
			lookups = append(lookups, svc.Mapper.MapCodeToLookup(c))
		}
		response = dtos.NewKresponse(http.StatusOK, lookups)
	} else {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	}

	return response
}
