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

var _ = Describe("UserService", func() {
	const (
		testIsPreloaded = false
	)

	var (
		mockCtrl     *g.Controller
		mockUserRepo *mocks.MockUserRepository
		service      svc.UserService
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockUserRepo = mocks.NewMockUserRepository(mockCtrl)

		service = svc.UserService{
			Repository: mockUserRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := svc.UserService{
				Repository: r.UserRepository{},
			}

			Expect(svc.NewUserService()).To(Equal(expected))
		})
	})

	Describe("GetMany", func() {
		It("should call repository.FindAll and return response with status ok", func() {
			testUsers := []s.User{
				{
					Base: s.Base{ID: 1},
				},
				{
					Base: s.Base{ID: 2},
				},
			}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				testUsers,
			)

			mockUserRepo.EXPECT().
				FindAll(testIsPreloaded).
				Return(testUsers, nil).
				Times(1)

			resultResponse := service.GetMany(testIsPreloaded)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if cannot find users", func() {
			testErr := errors.New("not found")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testErr,
			)

			mockUserRepo.EXPECT().
				FindAll(testIsPreloaded).
				Return([]s.User{}, testErr).
				Times(1)

			resultResponse := service.GetMany(testIsPreloaded)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("GetOne", func() {
		var testID int

		BeforeEach(func() {
			testID = 1
		})

		It("should call repository.FindOne and return response with status ok", func() {
			testUser := s.User{Base: s.Base{ID: testID}}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				testUser,
			)

			mockUserRepo.EXPECT().
				FindOne(testID, testIsPreloaded).
				Return(testUser, nil).
				Times(1)

			resultResponse := service.GetOne(testID, testIsPreloaded)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if cannot find user", func() {
			testErr := errors.New("not found")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testErr,
			)

			mockUserRepo.EXPECT().
				FindOne(testID, testIsPreloaded).
				Return(s.User{}, testErr).
				Times(1)

			resultResponse := service.GetOne(testID, testIsPreloaded)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("Post", func() {
		var (
			testID      int
			testAddress string
			testNewUser s.User
		)

		BeforeEach(func() {
			testID = 1
			testAddress = "test address"
			testNewUser = s.User{Address: testAddress}
		})

		It("should call repository.Create and return response with status ok", func() {
			testExpectedUser := testNewUser
			testExpectedUser.Base = s.Base{ID: testID}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				testExpectedUser,
			)

			mockUserRepo.EXPECT().
				Create(testNewUser).
				Return(testExpectedUser, nil).
				Times(1)

			resultResponse := service.Post(testNewUser)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if cannot create user", func() {
			testErr := errors.New("error")
			expectedResponse := dtos.NewKresponse(
				http.StatusBadRequest,
				testErr,
			)

			mockUserRepo.EXPECT().
				Create(testNewUser).
				Return(s.User{}, testErr).
				Times(1)

			resultResponse := service.Post(testNewUser)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("Put", func() {
		var (
			testID            int
			testUpdateNewUser s.User
		)

		BeforeEach(func() {
			testID = 1
			testUpdateNewUser = s.User{
				Base:    s.Base{ID: testID},
				Address: "test updated addr",
			}
		})

		It("should call repository.Update and return response with status ok", func() {
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				testUpdateNewUser,
			)

			mockUserRepo.EXPECT().
				Update(testUpdateNewUser).
				Return(testUpdateNewUser, nil).
				Times(1)

			resultResponse := service.Put(testUpdateNewUser)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if cannot update user", func() {
			testErr := errors.New("error")
			expectedResponse := dtos.NewKresponse(
				http.StatusBadRequest,
				testErr,
			)

			mockUserRepo.EXPECT().
				Update(testUpdateNewUser).
				Return(s.User{}, testErr).
				Times(1)

			resultResponse := service.Put(testUpdateNewUser)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("Delete", func() {
		var (
			testID int
		)

		BeforeEach(func() {
			testID = 1
		})

		It("should call repository.Update and return response with status ok", func() {
			testDeletedUser := s.User{Base: s.Base{ID: testID}}
			expectedResponse := dtos.NewKresponse(
				http.StatusOK,
				testDeletedUser,
			)

			mockUserRepo.EXPECT().
				Delete(testID).
				Return(testDeletedUser, nil).
				Times(1)

			resultResponse := service.Delete(testID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

		It("should return error if cannot update user", func() {
			testErr := errors.New("error")
			expectedResponse := dtos.NewKresponse(
				http.StatusNotFound,
				testErr,
			)

			mockUserRepo.EXPECT().
				Delete(testID).
				Return(s.User{}, testErr).
				Times(1)

			resultResponse := service.Delete(testID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})
})
