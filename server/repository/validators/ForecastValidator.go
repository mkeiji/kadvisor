package validators

import (
	"errors"
	util "kadvisor/server/libs/ValidationHelper"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"

	"github.com/go-playground/validator/v10"
)

type ForecastValidator struct {
	tagValidator *validator.Validate
	forecastRepo repository.ForecastRepository
}

func (f ForecastValidator) Validate(obj interface{}) []error {
	errList := []error{}
	forecast, _ := obj.(structs.Forecast)
	f.validateProperties(forecast, &errList)
	f.validateIsUnique(forecast, &errList)
	f.validateEntriesMonth(forecast, &errList)
	return errList
}

func (f ForecastValidator) validateProperties(
	forecast structs.Forecast,
	errList *[]error,
) {
	f.tagValidator = validator.New()

	err := f.tagValidator.Struct(forecast)
	if err != nil {
		*errList = append(*errList, err)
	}
}

func (f ForecastValidator) validateIsUnique(
	forecast structs.Forecast,
	errList *[]error,
) {
	_, fErr := f.forecastRepo.FindOne(forecast.UserID, forecast.Year, false)
	if fErr == nil {
		*errList = append(
			*errList,
			errors.New(util.GetValidationMsg(
				"User.Forecast",
				"forecast already exists",
			)),
		)
	}
}

func (f ForecastValidator) validateEntriesMonth(
	forecast structs.Forecast,
	errList *[]error,
) {
	var entriesMonth []int
	checked := map[int]bool{}

	for _, entry := range forecast.Entries {
		entriesMonth = append(entriesMonth, entry.Month)
	}

	for _, month := range entriesMonth {
		if checked[month] != true {
			checked[month] = true
		} else {
			*errList = append(
				*errList,
				errors.New(util.GetValidationMsg(
					"User.Forecast.Entries",
					"repeated month not allowed",
				)),
			)
		}
	}
}
