package services_test

import (
	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/dtos"
	r "kadvisor/server/repository"
	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	svc "kadvisor/server/services"
)

var _ = Describe("ForecastEntryService", func() {
	const (
		okResponse = 200
	)

	var (
		mockCtrl              *g.Controller
		mockForecastEntryRepo *mocks.MockForecastEntryRepository
		service               svc.ForecastEntryService
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockForecastEntryRepo = mocks.NewMockForecastEntryRepository(mockCtrl)

		service = svc.ForecastEntryService{
			Repository: mockForecastEntryRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := svc.ForecastEntryService{
				Repository: r.ForecastEntryRepository{},
			}

			Expect(svc.NewForecastEntryService()).To(Equal(expected))
		})
	})

	Describe("Put", func() {
		It("should call Repository.Update", func() {
			testID := s.Base{ID: 1}
			testForecastEntry := s.ForecastEntry{
				Base:    testID,
				Month:   1,
				Income:  2,
				Expense: 1,
			}
			expectedResponse := dtos.NewKresponse(okResponse, testForecastEntry)

			mockForecastEntryRepo.EXPECT().
				Update(testForecastEntry).
				Return(testForecastEntry, nil).
				Times(1)

			resultResponse := service.Put(testForecastEntry)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})
})
