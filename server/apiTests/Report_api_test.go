package apiTests_test

import (
	"net/http"
	"strconv"
	"time"

	"kadvisor/server/apiTests/ApiTestUtil"
	"kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ReportApi", func() {
	const (
		USER_ENDPOINT             = "/api/user"
		REPORT_ENDPOINT           = "/api/kadvisor/:uid/report"
		REPORT_AVAILABLE_ENDPOINT = "/api/kadvisor/:uid/reportavailable"
		ENTRY_ENDPOINT            = "/api/kadvisor/:uid/entry"
		CLASS_ENDPOINT            = "/api/kadvisor/:uid/class"
		FORECAST_ENDPOINT         = "/api/kadvisor/:uid/forecast"
		TEST_CLASS_NAME           = "testClass"
		TYPE_BALANCE              = "BALANCE"
		TYPE_YEAR_TO_DATE         = "YTD"
		TYPE_YEAR_FC              = "YFC"
	)
	var (
		today          time.Time
		testUser       s.User
		testClass      s.Class
		testEntry      s.Entry
		testForecast   s.Forecast
		defaultIncome  float64
		defaultExpense float64
	)

	BeforeEach(func() {
		today = time.Now()
		defaultIncome = float64(10)
		defaultExpense = float64(5)
	})

	postTestUser := func(user *s.User) {
		newUserReqBody := s.User{
			FirstName: "testLoginUser",
			IsPremium: false,
			Login: s.Login{
				RoleID:   2,
				Email:    KeiGenUtil.RandomString(8),
				UserName: "testUser",
				Password: "test",
			},
		}

		respStatus, respErr := kMakeRequest("POST", USER_ENDPOINT, newUserReqBody, &user, nil, nil)
		Expect(respErr).To(BeNil())
		Expect(respStatus).To(Equal(http.StatusOK))
	}

	postTestClass := func() {
		body := s.Class{
			UserID: testUser.Login.UserID,
			Name:   "testClass",
		}
		respStatus, respErr := kMakeRequest("POST", CLASS_ENDPOINT, body, &testClass, nil, testUser)
		Expect(respErr).To(BeNil())
		Expect(respStatus).To(Equal(http.StatusOK))
	}

	postTestEntry := func() {
		body := s.Entry{
			UserID:          testUser.Login.UserID,
			ClassID:         testClass.Base.ID,
			EntryTypeCodeID: "INCOME_ENTRY_TYPE",
			Date:            today,
			Amount:          defaultIncome,
		}

		respStatus, respErr := kMakeRequest("POST", ENTRY_ENDPOINT, body, &testEntry, nil, testUser)
		Expect(respErr).To(BeNil())
		Expect(respStatus).To(Equal(http.StatusOK))
	}

	postTestForecast := func() {
		respStatus, respErr := kMakeRequest(
			"POST",
			FORECAST_ENDPOINT,
			ApiTestUtil.CreateTestForecast(
				testUser.Base.ID,
				today.Year(),
				defaultIncome,
				defaultExpense,
			),
			&testForecast,
			nil,
			testUser,
		)
		Expect(respErr).To(BeNil())
		Expect(respStatus).To(Equal(http.StatusOK))
	}

	BeforeEach(func() {
		postTestUser(&testUser)
	})

	Describe("GetReport", func() {
		Context("Balance", func() {
			It("No error - should return balance with ok response", func() {
				expectedBalance := float64(10)
				postTestClass()
				postTestEntry()

				var result dtos.Balance
				params := map[string]string{"type": TYPE_BALANCE}
				respStatus, respErr := kMakeRequest("GET", REPORT_ENDPOINT, nil, &result, params, testUser)

				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.UserID).To(Equal(testUser.Base.ID))
				Expect(result.Balance).To(Equal(expectedBalance))
			})

			It("Error - not found response if user has not balance", func() {
				expectedErrMsg := "no balance is available"

				var result dtos.Balance
				params := map[string]string{"type": TYPE_BALANCE}
				respStatus, respErr := kMakeRequest("GET", REPORT_ENDPOINT, nil, &result, params, testUser)

				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})

		Context("YearToDate", func() {
			It("No error - should return report for the year with ok response", func() {
				postTestClass()
				postTestEntry()
				expected := []dtos.MonthReport{
					{
						Year:    today.Year(),
						Month:   int(today.Month()),
						Income:  defaultIncome,
						Expense: 0,
						Balance: defaultIncome,
					},
				}

				var result []dtos.MonthReport
				params := map[string]string{
					"type": TYPE_YEAR_TO_DATE,
					"year": strconv.Itoa(today.Year()),
				}
				respStatus, respErr := kMakeRequest("GET", REPORT_ENDPOINT, nil, &result, params, testUser)

				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result).To(Equal(expected))
			})

			It("Error - should return bad request response if year param is missing", func() {
				expectedErrMsg := "query param error"

				var result []dtos.MonthReport
				params := map[string]string{"type": TYPE_YEAR_TO_DATE}
				respStatus, respErr := kMakeRequest("GET", REPORT_ENDPOINT, nil, &result, params, testUser)

				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})

			It("Error - not found response if user has no report on a given year", func() {
				expectedErrMsg := "no report available"
				wrongYear := today.Year() + 1

				var result []dtos.MonthReport
				params := map[string]string{
					"type": TYPE_YEAR_TO_DATE,
					"year": strconv.Itoa(wrongYear),
				}
				respStatus, respErr := kMakeRequest("GET", REPORT_ENDPOINT, nil, &result, params, testUser)

				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})

		Context("YearToDateWithForecast", func() {
			It("No error - should return report for the year with ok response", func() {
				postTestClass()
				postTestEntry()
				postTestForecast()

				var result []dtos.MonthReport
				params := map[string]string{
					"type": TYPE_YEAR_FC,
					"year": strconv.Itoa(today.Year()),
				}
				respStatus, respErr := kMakeRequest("GET", REPORT_ENDPOINT, nil, &result, params, testUser)

				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(len(result)).To(Equal(12))
				verifyReportMonths(testEntry, testForecast, result)
			})

			It("Error - should return bad request response if year param is missing", func() {
				expectedErrMsg := "query param error"

				var result []dtos.MonthReport
				params := map[string]string{"type": TYPE_YEAR_FC}
				respStatus, respErr := kMakeRequest("GET", REPORT_ENDPOINT, nil, &result, params, testUser)

				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})

			It("Error - not found response if user has no report on a given year and no forecast", func() {
				expectedNoReportErrMsg := "no report available"
				expectedNoForecastErrMsg := "record not found"
				wrongYear := today.Year() + 1

				var result []dtos.MonthReport
				params := map[string]string{
					"type": TYPE_YEAR_FC,
					"year": strconv.Itoa(wrongYear),
				}
				respStatus, respErr := kMakeRequest("GET", REPORT_ENDPOINT, nil, &result, params, testUser)

				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(2))
				Expect(respErr[0].Error()).To(Equal(expectedNoReportErrMsg))
				Expect(respErr[1].Error()).To(Equal(expectedNoForecastErrMsg))
			})
		})
	})

	Describe("GetReportAvailable", func() {
		Context("No error", func() {
			var (
				result     []int
				respStatus int
				respErr    []error
			)
			AfterEach(func() {
				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(len(result)).To(Equal(1))
				Expect(result[0]).To(Equal(today.Year()))
			})

			It("should return available years if the user has an entry and no forecast", func() {
				postTestEntry()

				respStatus, respErr = kMakeRequest(
					"GET", REPORT_AVAILABLE_ENDPOINT, nil, &result, nil, testUser,
				)
			})

			It("should return available years if the user has a forecast and no entry", func() {
				postTestForecast()

				respStatus, respErr = kMakeRequest(
					"GET", REPORT_AVAILABLE_ENDPOINT, nil, &result, nil, testUser,
				)
			})

			It("with query - should return available years if the user has a forecast and no entry", func() {
				postTestForecast()

				params := map[string]string{"forecast": "true"}
				respStatus, respErr = kMakeRequest(
					"GET", REPORT_AVAILABLE_ENDPOINT, nil, &result, params, testUser,
				)
			})

		})

		Context("Error", func() {
			It("should return not found response if user has no entries and no forecast", func() {
				expectedNoReportErrMsg := "no available years found"

				var result []int
				respStatus, respErr := kMakeRequest(
					"GET", REPORT_AVAILABLE_ENDPOINT, nil, &result, nil, testUser,
				)
				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedNoReportErrMsg))

				params := map[string]string{"forecast": "true"}
				respWithParamStatus, respWithParamErr := kMakeRequest(
					"GET", REPORT_AVAILABLE_ENDPOINT, nil, &result, params, testUser,
				)
				Expect(respWithParamStatus).To(Equal(http.StatusNotFound))
				Expect(len(respWithParamErr)).To(Equal(1))
				Expect(respWithParamErr[0].Error()).To(Equal(expectedNoReportErrMsg))
			})
		})
	})
})

func verifyReportMonths(entry s.Entry, forecast s.Forecast, report []dtos.MonthReport) {
	getmonth := func(month int) dtos.MonthReport {
		for _, monthReport := range report {
			if monthReport.Month == month {
				return monthReport
			}
		}
		return dtos.MonthReport{}
	}

	entryMonth := int(entry.Date.Month())
	for _, reportEntry := range report {
		if reportEntry.Month < entryMonth {
			Expect(reportEntry.Income).To(Equal(float64(0)))
			Expect(reportEntry.Expense).To(Equal(float64(0)))
			Expect(reportEntry.Balance).To(Equal(float64(0)))
		} else if reportEntry.Month == entryMonth {
			if entry.EntryTypeCodeID == "INCOME_ENTRY_TYPE" {
				Expect(reportEntry.Income).To(Equal(entry.Amount))
			} else {
				Expect(reportEntry.Expense).To(Equal(entry.Amount * -1))
			}
		} else {
			Expect(reportEntry.Income).To(Equal(forecast.Entries[reportEntry.Month-1].Income))
			Expect(reportEntry.Expense).To(Equal(forecast.Entries[reportEntry.Month-1].Expense))
			Expect(reportEntry.Balance).To(Equal(
				getmonth(reportEntry.Month-1).Balance + reportEntry.Income + reportEntry.Expense,
			)) // use (+) expense because the value on the report is already (-)
		}
	}
}
