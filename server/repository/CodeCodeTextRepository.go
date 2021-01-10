package repository

import (
	"errors"
	"kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"

	"gorm.io/gorm"
)

type CodeCodeTextRepository struct {
	Db *gorm.DB
}

func NewCodeCodeTextRepository() CodeCodeTextRepository {
	return CodeCodeTextRepository{Db: app.Db}
}

func (this CodeCodeTextRepository) FindAllByCodeGroup(
	codeGroup string,
) ([]structs.Code, error) {
	var codes []structs.Code
	var err error

	this.Db.Where("code_group=?", codeGroup).Find(&codes)
	if len(codes) <= 0 {
		err = errors.New("code_group not found")
	}

	return codes, err
}

func (this CodeCodeTextRepository) FindOne(
	codeTypeID string,
) (structs.Code, error) {
	var code structs.Code
	err := this.Db.Where("code_type_id=?", codeTypeID).First(&code).Error
	return code, err
}
