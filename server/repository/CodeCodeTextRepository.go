package repository

import (
	"kadvisor/server/repository/mappers"
	"kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
)

type CodeCodeTextRepository struct {
	mapper mappers.LookupMapper
}

func (repo *CodeCodeTextRepository) FindAllByCodeGroup(
	codeGroup string) ([]structs.Code, error) {

	var codes []structs.Code
	query := structs.Code{CodeGroup: codeGroup}

	err := app.Db.Where(query).Find(&codes).Error
	return codes, err
}

func (repo *CodeCodeTextRepository) FindOne(
	codeTypeID string,
) (structs.Code, error) {
	var code structs.Code
	err := app.Db.Where("code_type_id=?", codeTypeID).First(&code).Error
	return code, err
}
