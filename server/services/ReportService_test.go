package services_test

import (
	"errors"
	"net/http"
	"time"

	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	r "kadvisor/server/repository"
	"kadvisor/server/repository/interfaces/mocks"
	m "kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	svc "kadvisor/server/services"
)

var _ = Describe("ReportService", func() {
	const (
		userID = 1
	)

	var (
		mockCtrl         *g.Controller
		mockReportRepo   *mocks.MockReportRepository
		mockForecastRepo *mocks.MockForecastRepository
		service          svc.ReportService
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockReportRepo = mocks.NewMockReportRepository(mockCtrl)
		mockForecastRepo = mocks.NewMockForecastRepository(mockCtrl)

		service = svc.ReportService{
			Repository:         mockReportRepo,
			ForecastRepository: mockForecastRepo,
			ForecastMapper:     m.ForecastMapper{},
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := svc.ReportService{
				Repository:         r.ReportRepository{},
				ForecastRepository: r.ForecastRepository{},
				ForecastMapper:     m.ForecastMapper{},
			}

			Expect(svc.NewReportService()).To(Equal(expected))
		})
	})

	Describe("GetBalance", func() {
		It("should call repository.FindBalance and return response with status ok", func() {
			expectedBalance := dtos.Balance{UserID: userID, Balance: 1.0}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				expectedBalance,
			)

			mockReportRepo.EXPECT().
				FindBalance(userID).
				Return(expectedBalance, nil).
				Times(1)

			resultResponse := service.GetBalance(userID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should call repository.FindBalance and return response with error if not found", func() {
			testErr := errors.New("not found")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testErr,
			)

			mockReportRepo.EXPECT().
				FindBalance(userID).
				Return(dtos.Balance{}, testErr).
				Times(1)

			resultResponse := service.GetBalance(userID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("GetYearToDateReport", func() {
		var (
			year  int
			month time.Month
		)

		BeforeEach(func() {
			year, month, _ = time.Now().Date()
		})

		It("should call repository.FindYearToDateReport ok response", func() {
			expectedReports := []dtos.MonthReport{
				{
					Year:    year,
					Month:   int(month),
					Income:  90,
					Expense: 10,
					Balance: 80,
				},
			}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				expectedReports,
			)

			mockReportRepo.EXPECT().
				FindYearToDateReport(userID, year).
				Return(expectedReports, nil).
				Times(1)

			resultResponse := service.GetYearToDateReport(userID, year)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if not found", func() {
			testErr := errors.New("not found")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testErr,
			)

			mockReportRepo.EXPECT().
				FindYearToDateReport(userID, year).
				Return([]dtos.MonthReport{}, testErr).
				Times(1)

			resultResponse := service.GetYearToDateReport(userID, year)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("GetYearToDateWithForecastReport", func() {
		var (
			year  int
			month int
		)

		BeforeEach(func() {
			year = 2020
			month = 1
		})

		It(`should call repository FindYearToDateReport,
            ForecastRepository.FindONe and return ok response`, func() {
			testForecast := createMockForecast(userID, year)
			testYearToDateReport := []dtos.MonthReport{
				{
					Year:    year,
					Month:   int(month),
					Income:  90,
					Expense: 10,
					Balance: 80,
				},
			}

			expectedReport := getExpectedYearToDateReport()
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				expectedReport,
			)

			mockReportRepo.EXPECT().
				FindYearToDateReport(userID, year).
				Return(testYearToDateReport, nil).
				Times(1)
			mockForecastRepo.EXPECT().
				FindOne(userID, year, true).
				Return(testForecast, nil).
				Times(1)

			resultResponse := service.GetYearToDateWithForecastReport(userID, year)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return one error if repository.FindYearToDateReport cannot be found", func() {
			testErr := errors.New("FindYearToDateReport not found")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				KeiGenUtil.MapErrList([]error{testErr}),
			)

			mockReportRepo.EXPECT().
				FindYearToDateReport(userID, year).
				Return([]dtos.MonthReport{}, testErr).
				Times(1)
			mockForecastRepo.EXPECT().
				FindOne(userID, year, true).
				Return(s.Forecast{}, nil).
				Times(1)

			resultResponse := service.GetYearToDateWithForecastReport(userID, year)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return one error if forecastRepository.FindOne cannot be found", func() {
			testErr := errors.New("forecastRepository not found")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				KeiGenUtil.MapErrList([]error{testErr}),
			)

			mockReportRepo.EXPECT().
				FindYearToDateReport(userID, year).
				Return([]dtos.MonthReport{}, nil).
				Times(1)
			mockForecastRepo.EXPECT().
				FindOne(userID, year, true).
				Return(s.Forecast{}, testErr).
				Times(1)

			resultResponse := service.GetYearToDateWithForecastReport(userID, year)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It(`should return two errors if repository.FindYearToDateReport
            and forecastRepository.FindOne cannot be found`, func() {
			testRepositoryErr := errors.New("FindYearToDateReport not found")
			testForecastRepoErr := errors.New("forecastRepository not found")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				KeiGenUtil.MapErrList(
					[]error{testRepositoryErr, testForecastRepoErr},
				),
			)

			mockReportRepo.EXPECT().
				FindYearToDateReport(userID, year).
				Return([]dtos.MonthReport{}, testRepositoryErr).
				Times(1)
			mockForecastRepo.EXPECT().
				FindOne(userID, year, true).
				Return(s.Forecast{}, testForecastRepoErr).
				Times(1)

			resultResponse := service.GetYearToDateWithForecastReport(userID, year)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("GetReportForecastAvailable", func() {
		It("should call repository.GetAvailableForecastYears and return response ok", func() {
			expectedYears := []int{2020, 2021}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				expectedYears,
			)

			mockReportRepo.EXPECT().
				GetAvailableForecastYears(userID).
				Return(expectedYears, nil).
				Times(1)

			resultResponse := service.GetReportForecastAvailable(userID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if no year is found", func() {
			testErr := errors.New("forecastRepository not found")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testErr,
			)

			mockReportRepo.EXPECT().
				GetAvailableForecastYears(userID).
				Return([]int{}, testErr).
				Times(1)

			resultResponse := service.GetReportForecastAvailable(userID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("GetReportAvailable", func() {
		It("should call repository.GetAvailableYears and return response ok", func() {
			expectedYears := []int{2021}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				expectedYears,
			)

			mockReportRepo.EXPECT().
				GetAvailableYears(userID).
				Return(expectedYears, nil).
				Times(1)

			resultResponse := service.GetReportAvailable(userID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if no year is found", func() {
			testErr := errors.New("not found")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testErr,
			)

			mockReportRepo.EXPECT().
				GetAvailableYears(userID).
				Return([]int{}, testErr).
				Times(1)

			resultResponse := service.GetReportAvailable(userID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})
})

func createMockForecast(userID int, year int) s.Forecast {
	entries := []s.ForecastEntry{}

	for i := 1; i < 13; i++ {
		entries = append(entries, s.ForecastEntry{
			ForecastID: 1,
			Month:      i,
			Income:     100,
			Expense:    10,
		})
	}

	return s.Forecast{
		UserID:  userID,
		Year:    year,
		Entries: entries,
	}
}

func getExpectedYearToDateReport() []dtos.MonthReport {
	return []dtos.MonthReport{
		{Year: 2020, Month: 1, Income: 90, Expense: 10, Balance: 80, Type: "ACTUAL"},
		{Year: 2020, Month: 2, Income: 100, Expense: 10, Balance: 190, Type: "FORECAST"},
		{Year: 2020, Month: 3, Income: 100, Expense: 10, Balance: 300, Type: "FORECAST"},
		{Year: 2020, Month: 4, Income: 100, Expense: 10, Balance: 410, Type: "FORECAST"},
		{Year: 2020, Month: 5, Income: 100, Expense: 10, Balance: 520, Type: "FORECAST"},
		{Year: 2020, Month: 6, Income: 100, Expense: 10, Balance: 630, Type: "FORECAST"},
		{Year: 2020, Month: 7, Income: 100, Expense: 10, Balance: 740, Type: "FORECAST"},
		{Year: 2020, Month: 8, Income: 100, Expense: 10, Balance: 850, Type: "FORECAST"},
		{Year: 2020, Month: 9, Income: 100, Expense: 10, Balance: 960, Type: "FORECAST"},
		{Year: 2020, Month: 10, Income: 100, Expense: 10, Balance: 1070, Type: "FORECAST"},
		{Year: 2020, Month: 11, Income: 100, Expense: 10, Balance: 1180, Type: "FORECAST"},
		{Year: 2020, Month: 12, Income: 100, Expense: 10, Balance: 1290, Type: "FORECAST"},
	}
}
