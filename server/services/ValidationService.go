package services

import (
	"kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/interfaces"
	"net/http"
)

type ValidationService struct{}

func NewValidationService() ValidationService {
	return ValidationService{}
}

func (v ValidationService) GetResponse(
	validator interfaces.Validator,
	obj interface{},
) dtos.KhttpResponse {
	var status int

	errList := validator.Validate(obj)
	if len(errList) > 0 {
		status = http.StatusBadRequest
	} else {
		status = http.StatusOK
	}

	return dtos.NewKresponse(
		status,
		KeiGenUtil.MapErrList(errList),
	)
}
