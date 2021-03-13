package apiTests_test

import (
	"net/http"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	s "kadvisor/server/repository/structs"
)

var _ = Describe("ForecastEntryApi", func() {
	const (
		FORECAST_ENTRY_ENDPOINT = "/api/kadvisor/:uid/forecastentry"
		FORECAST_ENDPOINT       = "/api/kadvisor/:uid/forecast"
	)

	var (
		testForecast s.Forecast
	)

	buildTestForecast := func() s.Forecast {
		forecast := s.Forecast{
			UserID: testUserRegular.Login.UserID,
			Year:   time.Now().Year(),
		}

		for i := 1; i <= 12; i++ {
			entry := s.ForecastEntry{
				Month:   i,
				Income:  10.0,
				Expense: 5.0,
			}
			forecast.Entries = append(forecast.Entries, entry)
		}

		return forecast
	}

	postTestForecast := func() {
		respStatus, postForecastErr := kMakeRequest(
			"POST", FORECAST_ENDPOINT, buildTestForecast(), &testForecast, nil, nil,
		)
		Expect(postForecastErr).To(BeNil())
		Expect(respStatus).To(Equal(http.StatusOK))
	}

	deleteTestForecast := func(forecast s.Forecast) {
		params := map[string]string{"id": strconv.Itoa(forecast.Base.ID)}
		respStatus, deleteErr := kMakeRequest("DELETE", FORECAST_ENDPOINT, nil, nil, params, nil)
		Expect(deleteErr).To(BeNil())
		Expect(respStatus).To(Equal(http.StatusOK))
	}

	BeforeEach(func() {
		postTestForecast()
	})

	AfterEach(func() {
		deleteTestForecast(testForecast)
	})

	Describe("PutForecastEntry", func() {
		Context("No error", func() {
			It("should return updated forecastEntry with ok response", func() {
				entryID := testForecast.Entries[0].Base.ID
				newIncome := 50.0
				update := s.ForecastEntry{
					Base:   s.Base{ID: entryID},
					Income: newIncome,
				}

				var result s.ForecastEntry
				respStatus, respErr := kMakeRequest("PUT", FORECAST_ENTRY_ENDPOINT, update, &result, nil, nil)

				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.Base.ID).To(Equal(entryID))
				Expect(result.Income).To(Equal(newIncome))
			})
		})

		Context("Error", func() {
			It("should return not found if id is invalid", func() {
				invalidID := 9999
				newIncome := 50.0
				expectedErrMsg := "record not found"
				update := s.ForecastEntry{
					Base:   s.Base{ID: invalidID},
					Income: newIncome,
				}

				var result s.ForecastEntry
				respStatus, respErr := kMakeRequest("PUT", FORECAST_ENTRY_ENDPOINT, update, &result, nil, nil)

				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})
})
