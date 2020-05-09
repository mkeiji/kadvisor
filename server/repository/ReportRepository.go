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

func (repo *ReportRepository) FindYearToDateReport(userID int) ([]dtos.MonthReport, error) {
	var mReport []dtos.MonthReport

	query := `
		select 
			year(date) year, 
			month(date) month, 
			sum(income) income, 
			sum(expense) expense, 
			(sum(income) + sum(expense)) balance 
		from (
			select date, 
				case when entry_type_code_id='INCOME_ENTRY_TYPE' then amount else 0 end income, 
				case when entry_type_code_id='EXPENSE_ENTRY_TYPE' then amount else 0 end expense 
			from entries 
			where user_id=? 
				and year(date)=year(now())
		) yearly
		group by year(date), month(date);
	`

	err := application.Db.Raw(query, userID).Scan(&mReport).Error

	return mReport, err
}