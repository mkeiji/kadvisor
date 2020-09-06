package repository

import (
	"kadvisor/server/repository/mappers"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type CodeCodeTextRepository struct {
	mapper mappers.LookupMapper
}

func (repo *CodeCodeTextRepository) FindAllByCodeGroup(
	codeGroup string) ([]structs.Code, error) {

	var codes []structs.Code
	query := structs.Code{CodeGroup: codeGroup}

	err := application.Db.Where(query).Find(&codes).Error
	return codes, err
}
