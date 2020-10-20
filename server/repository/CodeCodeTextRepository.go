package repository

import (
	"errors"
	"kadvisor/server/repository/mappers"
	"kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
)

type CodeCodeTextRepository struct {
	mapper mappers.LookupMapper
}

func (repo *CodeCodeTextRepository) FindAllByCodeGroup(
	codeGroup string,
) ([]structs.Code, error) {
	var codes []structs.Code
	var err error

	app.Db.Where("code_group=?", codeGroup).Find(&codes)
	if len(codes) <= 0 {
		err = errors.New("code_group not found")
	}

	return codes, err
}

func (repo *CodeCodeTextRepository) FindOne(
	codeTypeID string,
) (structs.Code, error) {
	var code structs.Code
	err := app.Db.Where("code_type_id=?", codeTypeID).First(&code).Error
	return code, err
}
