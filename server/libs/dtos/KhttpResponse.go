package dtos

import "kadvisor/server/libs/KeiGenUtil"

type KhttpResponse struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func NewBadKresponse(err error) KhttpResponse {
	return NewKresponse(400, err)
}

func NewKresponse(status int, body interface{}) KhttpResponse {
	_, isError := body.(error)
	if isError {
		newBody := KeiGenUtil.MapErrList([]error{body.(error)})
		return KhttpResponse{
			Status: status,
			Body:   newBody,
		}
	} else {
		return KhttpResponse{
			Status: status,
			Body:   body,
		}
	}
}
