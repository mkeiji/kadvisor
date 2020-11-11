package services_test

import (
	"errors"
	"net/http"
	"time"

	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/dtos"
	r "kadvisor/server/repository"
	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	svc "kadvisor/server/services"
)

var _ = Describe("ForecastService", func() {
	const (
		okResponse = 200
	)

	var (
		mockCtrl               *g.Controller
		mockForecastRepository *mocks.MockForecastRepository
		service                svc.ForecastService
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockForecastRepository = mocks.NewMockForecastRepository(mockCtrl)

		service = svc.ForecastService{
			Repository: mockForecastRepository,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := svc.ForecastService{
				Repository: r.ForecastRepository{},
			}

			Expect(svc.NewForecastService()).To(Equal(expected))
		})
	})

	Describe("GetOne", func() {
		It("should call Repository.FindOne and return forecast", func() {
			testUserID := 1
			testYear := time.Now().Year()
			testIsPreloaded := true
			expectedForecast := s.Forecast{Base: s.Base{ID: 1}}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				expectedForecast,
			)

			mockForecastRepository.EXPECT().
				FindOne(testUserID, testYear, testIsPreloaded).
				Return(expectedForecast, nil).
				Times(1)

			resultResponse := service.GetOne(testUserID, testYear, testIsPreloaded)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if year is not present", func() {
			testUserID := 1
			testYear := 0
			testIsPreloaded := true
			expectedResponse := dtos.NewKresponse(
				http.StatusBadRequest,
				errors.New("year param is required"),
			)

			mockForecastRepository.EXPECT().
				FindOne(testUserID, testYear, testIsPreloaded).
				Times(0)

			resultResponse := service.GetOne(testUserID, testYear, testIsPreloaded)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if forecast is not found", func() {
			testUserID := 1
			testYear := time.Now().Year()
			testIsPreloaded := true
			testErr := errors.New("year param is required")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testErr,
			)

			mockForecastRepository.EXPECT().
				FindOne(testUserID, testYear, testIsPreloaded).
				Return(s.Forecast{}, testErr).
				Times(1)

			resultResponse := service.GetOne(testUserID, testYear, testIsPreloaded)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("Post", func() {
		It("should call Repository.Create and return forecast", func() {
			testForecast := s.Forecast{}
			expectedForecast := s.Forecast{Base: s.Base{ID: 1}}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				expectedForecast,
			)

			mockForecastRepository.EXPECT().
				Create(testForecast).
				Return(expectedForecast, nil).
				Times(1)

			resultResponse := service.Post(testForecast)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if save fails", func() {
			testForecast := s.Forecast{}
			expectedErr := errors.New("failed to save")
			expectedResponse := dtos.NewKresponse(
				http.StatusBadRequest,
				expectedErr,
			)

			mockForecastRepository.EXPECT().
				Create(testForecast).
				Return(s.Forecast{}, expectedErr).
				Times(1)

			resultResponse := service.Post(testForecast)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("Delete", func() {
		It("should call Repository.Delete", func() {
			testID := 1
			deletedForecast := s.Forecast{}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				deletedForecast,
			)

			mockForecastRepository.EXPECT().
				Delete(testID).
				Return(deletedForecast, nil).
				Times(1)

			resultResponse := service.Delete(testID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if cannot find forecast to delete", func() {
			testID := 2
			expectedErr := errors.New("not found error")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				expectedErr,
			)

			mockForecastRepository.EXPECT().
				Delete(testID).
				Return(s.Forecast{}, expectedErr).
				Times(1)

			resultResponse := service.Delete(testID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})
})
