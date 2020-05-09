package repository

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/resources/application"
)

type ReportRepository struct {}

func (repo *ReportRepository) FindBalance(userID int) (dtos.Balance, error) {
	var balance dtos.Balance

	err := application.Db.Table("entries").Select(
		"user_id as user_id, sum(amount) as balance").Group(
			"user_id").Where(
				"user_id=?", userID).Scan(&balance).Error

	return balance, err
}