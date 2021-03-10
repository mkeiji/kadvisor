package apiTests_test

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	s "kadvisor/server/repository/structs"
)

var _ = Describe("ForecastApi", func() {
	const (
		FORECAST_ENDPOINT = "/api/kadvisor/:uid/forecast"
	)

	var (
		testYear int
	)

	buildTestForecast := func() s.Forecast {
		return s.Forecast{
			UserID: testUserRegular.Login.UserID,
			Year:   testYear,
		}
	}

	postTestForecast := func(forecast *s.Forecast) {
		respStatus, postErr := kMakeRequest("POST", FORECAST_ENDPOINT, buildTestForecast(), &forecast, nil, nil)
		if postErr != nil {
			panic(fmt.Sprintf("\nError: %v", postErr))
		}

		Expect(respStatus).To(Equal(http.StatusOK))
	}

	deleteTestForecast := func(forecast s.Forecast) {
		params := map[string]string{"id": strconv.Itoa(forecast.Base.ID)}
		respStatus, deleteErr := kMakeRequest("DELETE", FORECAST_ENDPOINT, nil, nil, params, nil)
		if deleteErr != nil {
			panic(fmt.Sprintf("\nError: %v", deleteErr))
		}

		Expect(respStatus).To(Equal(http.StatusOK))
	}

	BeforeEach(func() {
		testYear = time.Now().Year()
	})

	Describe("GetOneForecast", func() {
		Context("No error", func() {
			It("should get forecast and return ok response", func() {
				var testForecast s.Forecast
				postTestForecast(&testForecast)

				params := map[string]string{"year": strconv.Itoa(testYear)}
				var result s.Forecast
				respStatus, _ := kMakeRequest(
					"GET", FORECAST_ENDPOINT, nil, &result, params, nil,
				)

				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result).To(Equal(testForecast))

				deleteTestForecast(testForecast)
			})
		})

		Context("Error", func() {
			It("should return not found if not exist for user", func() {
				expectedErrMsg := "record not found"
				params := map[string]string{"year": strconv.Itoa(testYear)}

				respStatus, respErr := kMakeRequest(
					"GET", FORECAST_ENDPOINT, nil, nil, params, nil,
				)

				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})

	Describe("PostForecast", func() {
		Context("No error", func() {
			It("should post a forecast and return ok response", func() {
				var result s.Forecast
				respStatus, postForecastErr := kMakeRequest(
					"POST", FORECAST_ENDPOINT, buildTestForecast(), &result, nil, nil,
				)

				Expect(postForecastErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.Base.ID).NotTo(BeNil())
				Expect(result.UserID).To(Equal(testUserRegular.Login.UserID))
				Expect(result.Year).To(Equal(testYear))

				deleteTestForecast(result)
			})
		})

		Context("Error", func() {
			var (
				respStatus     int
				respErr        []error
				expectedErrMsg string
			)

			It("should return bad request if year is missing", func() {
				expectedErrMsg = "Key: 'Forecast.Year' Error:Field validation for 'Year' failed on the 'required' tag"
				invalidForecast := s.Forecast{
					Base: s.Base{ID: testUserRegular.Login.UserID},
				}

				respStatus, respErr = kMakeRequest(
					"POST", FORECAST_ENDPOINT, invalidForecast, nil, nil, nil,
				)

				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})

			It("should return bad request if forecast already exist for the user", func() {
				expectedErrMsg = "Key: 'User.Forecast' Error:Field validation for 'Forecast' forecast already exists"
				var forecast s.Forecast
				postTestForecast(&forecast)

				respStatus, respErr = kMakeRequest(
					"POST", FORECAST_ENDPOINT, buildTestForecast(), nil, nil, nil,
				)

				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))

				deleteTestForecast(forecast)
			})
		})
	})

	Describe("DeleteForecast", func() {
		Context("No error", func() {
			It("should delete forecast and return ok response", func() {
				var forecast s.Forecast
				postTestForecast(&forecast)
				Expect(forecast.Base.ID).NotTo(BeNil())

				params := map[string]string{"id": strconv.Itoa(forecast.Base.ID)}
				var deleted s.Forecast
				respStatus, deleteErr := kMakeRequest("DELETE", FORECAST_ENDPOINT, nil, &deleted, params, nil)

				Expect(deleteErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(deleted.Base.ID).To(Equal(forecast.Base.ID))
			})
		})

		Context("Error", func() {
			It("should return error if id is invalid", func() {
				invalidID := 999
				expectedErrMsg := "record not found"

				params := map[string]string{"id": strconv.Itoa(invalidID)}
				var deleted s.Forecast
				respStatus, respErr := kMakeRequest("DELETE", FORECAST_ENDPOINT, nil, &deleted, params, nil)

				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})
})
