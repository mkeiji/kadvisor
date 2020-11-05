package validators

import (
	"errors"
	util "kadvisor/server/libs/ValidationHelper"
	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
	s "kadvisor/server/repository/structs"
)

type ForecastValidator struct {
	TagValidator i.TagValidator
	ForecastRepo i.ForecastRepository
}

func NewForecastValidator() ForecastValidator {
	return ForecastValidator{
		TagValidator: TagValidator{},
		ForecastRepo: r.ForecastRepository{},
	}
}

func (f ForecastValidator) Validate(obj interface{}) []error {
	errList := []error{}
	forecast, _ := obj.(s.Forecast)
	f.validateProperties(forecast, &errList)
	f.validateIsUnique(forecast, &errList)
	f.validateEntriesMonth(forecast, &errList)
	return errList
}

func (f ForecastValidator) validateProperties(
	forecast s.Forecast,
	errList *[]error,
) {
	err := f.TagValidator.ValidateStruct(forecast)
	if err != nil {
		*errList = append(*errList, err)
	}
}

func (f ForecastValidator) validateIsUnique(
	forecast s.Forecast,
	errList *[]error,
) {
	_, fErr := f.ForecastRepo.FindOne(forecast.UserID, forecast.Year, false)
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
	forecast s.Forecast,
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
