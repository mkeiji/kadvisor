package repository

import (
	"errors"
	"fmt"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/resources/application"
)

type Year struct {
	Year int `json:"year"`
}

type ReportRepository struct{}

func (repo *ReportRepository) GetAvailableYears(userID int) ([]int, error) {
	var years []Year
	var result []int

	query := fmt.Sprintf(`
        select distinct year(date) as year
        from entries where user_id=%v
        union
        select distinct year as year
        from forecasts where user_id=%v
	`, userID, userID)

	err := application.Db.Raw(query).Scan(&years).Error
	if len(years) <= 0 {
		err = errors.New("no available years found")
	} else {
		for _, y := range years {
			result = append(result, y.Year)
		}
	}

	return result, err
}

func (repo *ReportRepository) FindBalance(userID int) (dtos.Balance, error) {
	var balance dtos.Balance

	err := application.Db.Table("entries").Select(
		"user_id as user_id, sum(amount) as balance").Group(
		"user_id").Where(
		"user_id=?", userID).Scan(&balance).Error
	if balance.UserID == 0 && balance.Balance == 0 {
		err = errors.New("no balance is available")
	}

	return balance, err
}

func (repo *ReportRepository) FindYearToDateReport(userID int, year int) ([]dtos.MonthReport, error) {
	var mReport []dtos.MonthReport

	query := fmt.Sprintf(`
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
				and year(date)=%d
		) yearly
		group by year(date), month(date);
	`, year)

	err := application.Db.Raw(query, userID).Scan(&mReport).Error

	if len(mReport) <= 0 {
		err = errors.New("no report available")
	}

	return mReport, err
}
