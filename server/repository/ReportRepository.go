package repository

import (
	"errors"
	"fmt"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"

	"gorm.io/gorm"
)

type Year struct {
	Year int `json:"year"`
}

type ReportRepository struct {
	Db *gorm.DB
}

func NewReportRepository() ReportRepository {
	return ReportRepository{
		Db: app.Db,
	}
}

func (this ReportRepository) GetAvailableForecastYears(userID int) ([]int, error) {
	var forecasts []structs.Forecast
	var result []int

	err := this.Db.
		Where("user_id=?", userID).
		Distinct("year").
		Order("year desc").
		Find(&forecasts).
		Error
	if err == nil {
		if len(forecasts) <= 0 {
			err = errors.New("no available years found")
		} else {
			for _, f := range forecasts {
				result = append(result, f.Year)
			}
		}
	}

	return result, err
}

func (this ReportRepository) GetAvailableYears(userID int) ([]int, error) {
	var years []Year
	var result []int

	query := fmt.Sprintf(`
        SELECT DISTINCT year(date) as year
        FROM entries 
		WHERE user_id=%v
        UNION
        SELECT DISTINCT year as year
        FROM forecasts
		WHERE user_id=%v
	`, userID, userID)

	err := this.Db.Raw(query).Scan(&years).Error
	if err == nil {
		if len(years) <= 0 {
			err = errors.New("no available years found")
		} else {
			for _, y := range years {
				result = append(result, y.Year)
			}
		}
	}

	return result, err
}

func (this ReportRepository) FindBalance(userID int) (dtos.Balance, error) {
	var balance dtos.Balance

	err := this.Db.Table("entries").Select(
		"user_id as user_id, sum(amount) as balance").Group(
		"user_id").Where(
		"user_id=?", userID).Scan(&balance).Error
	if err == nil {
		if balance.UserID == 0 && balance.Balance == 0 {
			err = errors.New("no balance is available")
		}
	}

	return balance, err
}

func (this ReportRepository) FindYearToDateReport(userID int, year int) ([]dtos.MonthReport, error) {
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

	err := this.Db.Raw(query, userID).Scan(&mReport).Error
	if err == nil {
		if len(mReport) <= 0 {
			err = errors.New("no report available")
		}
	}

	return mReport, err
}
