package mappers

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
)

type LookupMapper struct {}

func (l LookupMapper) MapCodeToLookup(
	code structs.Code) dtos.LookupEntry {

	return dtos.LookupEntry{
		Id: code.ID,
		Text: code.Name,
		Code: code.CodeTypeID,
	}
}
