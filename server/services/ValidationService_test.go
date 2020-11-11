package services_test

import (
	"errors"
	"net/http"

	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/interfaces/mocks"
	svc "kadvisor/server/services"
)

var _ = Describe("ValidationService", func() {
	const (
		okResponse = 200
	)

	var (
		mockCtrl      *g.Controller
		mockValidator *mocks.MockValidator
		service       svc.ValidationService
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockValidator = mocks.NewMockValidator(mockCtrl)

		service = svc.ValidationService{}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("GetResponse", func() {
		It("should return an error list with status badRequest if validator finds error", func() {
			testErr := errors.New("test error")
			testErrList := []error{testErr}
			expectedErrList := KeiGenUtil.MapErrList(testErrList)
			expectedResponse := dtos.NewKresponse(
				http.StatusBadRequest,
				expectedErrList,
			)

			mockValidator.EXPECT().
				Validate(g.Any()).
				Return(testErrList).
				Times(1)

			resultResponse := service.GetResponse(mockValidator, g.Any())
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return empty list with status ok if no error is found", func() {
			emptyErrList := []error{}
			expectedErrList := KeiGenUtil.MapErrList(emptyErrList)
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				expectedErrList,
			)

			mockValidator.EXPECT().
				Validate(g.Any()).
				Return(emptyErrList).
				Times(1)

			resultResponse := service.GetResponse(mockValidator, g.Any())
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})
})
