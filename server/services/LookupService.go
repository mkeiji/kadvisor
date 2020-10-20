package services

import (
	"errors"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/mappers"
	"net/http"
)

type LookupService struct {
	mapper     mappers.LookupMapper
	repository repository.CodeCodeTextRepository
}

func (svc *LookupService) GetAllByCodeGroup(
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

	codes, err := svc.repository.FindAllByCodeGroup(codeGroup)
	if err == nil {
		for _, c := range codes {
			lookups = append(lookups, svc.mapper.MapCodeToLookup(c))
		}
		response = dtos.NewKresponse(http.StatusOK, lookups)
	} else {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	}

	return response
}
