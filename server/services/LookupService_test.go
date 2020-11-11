package services_test

import (
	"errors"
	"net/http"

	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/dtos"
	r "kadvisor/server/repository"
	"kadvisor/server/repository/interfaces/mocks"
	"kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	svc "kadvisor/server/services"
)

var _ = Describe("LookupService", func() {
	const (
		okResponse = 200
	)

	var (
		mockCtrl         *g.Controller
		mockCodeTextRepo *mocks.MockCodeCodeTextRepository
		service          svc.LookupService
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockCodeTextRepo = mocks.NewMockCodeCodeTextRepository(mockCtrl)

		service = svc.LookupService{
			Mapper:     mappers.LookupMapper{},
			Repository: mockCodeTextRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := svc.LookupService{
				Mapper:     mappers.LookupMapper{},
				Repository: r.CodeCodeTextRepository{},
			}

			Expect(svc.NewLookupService()).To(Equal(expected))
		})
	})

	Describe("GetAllByCodeGroup", func() {
		It("should return error if codeGroup is empty", func() {
			testCodeGroup := ""
			expectedResponse := dtos.NewKresponse(
				http.StatusBadRequest,
				errors.New("missing codeGroup param"),
			)

			mockCodeTextRepo.EXPECT().
				FindAllByCodeGroup(testCodeGroup).
				Times(0)

			resultResponse := service.GetAllByCodeGroup(testCodeGroup)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should call Repository.FindAllByCodeGroup and return lookups with status ok", func() {
			testCodeGroup := "testCode"
			testCodes := []s.Code{
				{
					Base:       s.Base{ID: 1},
					Name:       "test text",
					CodeTypeID: "testCode",
				},
			}
			expectedLookups := []dtos.LookupEntry{
				{
					Id:   testCodes[0].ID,
					Text: testCodes[0].Name,
					Code: testCodes[0].CodeTypeID,
				},
			}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				expectedLookups,
			)

			mockCodeTextRepo.EXPECT().
				FindAllByCodeGroup(testCodeGroup).
				Return(testCodes, nil).
				Times(1)

			resultResponse := service.GetAllByCodeGroup(testCodeGroup)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if codeGroup cannot be found", func() {
			testCodeGroup := "testCodeGroup"
			testErr := errors.New("testError")
			expectedResponse := dtos.NewKresponse(
				http.StatusBadRequest,
				testErr,
			)

			mockCodeTextRepo.EXPECT().
				FindAllByCodeGroup(testCodeGroup).
				Return([]s.Code{}, testErr).
				Times(1)

			resultResponse := service.GetAllByCodeGroup(testCodeGroup)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})
})
