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
	s "kadvisor/server/repository/structs"
	svc "kadvisor/server/services"
)

var _ = Describe("LoginService", func() {
	var (
		mockCtrl      *g.Controller
		mockLoginRepo *mocks.MockLoginRepository
		service       svc.LoginService
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockLoginRepo = mocks.NewMockLoginRepository(mockCtrl)

		service = svc.LoginService{
			Repository: mockLoginRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := svc.LoginService{
				Repository: r.LoginRepository{},
			}

			Expect(svc.NewLoginService()).To(Equal(expected))
		})
	})

	Describe("GetOneByEmail", func() {
		It("should call repository.FindOneByEmail and return response with status ok", func() {
			testEmail := "test@email.com"
			testLogin := s.Login{Base: s.Base{ID: 1}}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				testLogin,
			)

			mockLoginRepo.EXPECT().
				FindOneByEmail(testEmail).
				Return(testLogin, nil).
				Times(1)

			resultResponse := service.GetOneByEmail(testEmail)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should call return error if email not found", func() {
			testEmail := "test@email.com"
			testError := errors.New("test error")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testError,
			)

			mockLoginRepo.EXPECT().
				FindOneByEmail(testEmail).
				Return(s.Login{}, testError).
				Times(1)

			resultResponse := service.GetOneByEmail(testEmail)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("Put", func() {
		It("should call repository.Update and return response with status ok", func() {
			testLogin := s.Login{
				Base:  s.Base{ID: 1},
				Email: "updated@email.com",
			}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				testLogin,
			)

			mockLoginRepo.EXPECT().
				Update(testLogin).
				Return(testLogin, nil).
				Times(1)

			resultResponse := service.Put(testLogin)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if cannot update", func() {
			testLogin := s.Login{
				Base:  s.Base{ID: 1},
				Email: "updated@email.com",
			}
			testError := errors.New("test error")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testError,
			)

			mockLoginRepo.EXPECT().
				Update(testLogin).
				Return(s.Login{}, testError).
				Times(1)

			resultResponse := service.Put(testLogin)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("UpdateLoginStatus", func() {
		It("should call repository.UpdateLoginStatus and return response with status ok", func() {
			testLoginStatus := true
			testLogin := s.Login{
				Base:       s.Base{ID: 1},
				Email:      "updated@email.com",
				IsLoggedIn: testLoginStatus,
			}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				testLogin,
			)

			mockLoginRepo.EXPECT().
				UpdateLoginStatus(testLogin, testLoginStatus).
				Return(testLogin, nil).
				Times(1)

			resultResponse := service.UpdateLoginStatus(testLogin, testLoginStatus)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if cannot update", func() {
			testLoginStatus := false
			testLogin := s.Login{Base: s.Base{ID: 1}}
			testError := errors.New("test error")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testError,
			)

			mockLoginRepo.EXPECT().
				UpdateLoginStatus(testLogin, testLoginStatus).
				Return(s.Login{}, testError).
				Times(1)

			resultResponse := service.UpdateLoginStatus(testLogin, testLoginStatus)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})
})
