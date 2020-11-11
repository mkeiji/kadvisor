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

var _ = Describe("ClassService", func() {
	const (
		okResponse = 200
		testUserID = 1
	)

	var (
		mockCtrl      *g.Controller
		mockClassRepo *mocks.MockClassRepository
		service       svc.ClassService
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockClassRepo = mocks.NewMockClassRepository(mockCtrl)

		service = svc.ClassService{
			Repository: mockClassRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := svc.ClassService{
				Repository: r.ClassRepository{},
			}

			Expect(svc.NewClassService()).To(Equal(expected))
		})
	})

	Describe("GetClass", func() {
		It("should call FindOne if classID is not empty (zero)", func() {
			testClassID := 1
			mockClassRepo.EXPECT().
				FindOne(g.Any()).
				Return(s.Class{}, nil).
				Times(1)

			service.GetClass(testUserID, testClassID)
		})

		It("should call FindAllByUserId if classID is empty (zero)", func() {
			testClassID := 0
			mockClassRepo.EXPECT().
				FindAllByUserId(g.Any()).
				Return([]s.Class{}, nil).
				Times(1)

			service.GetClass(testUserID, testClassID)
		})
	})

	Describe("GetOneById", func() {
		It("should call FindOne", func() {
			testClassID := 1
			testClass := s.Class{
				Base: s.Base{ID: testClassID},
			}
			expectedResponse := dtos.NewKresponse(okResponse, testClass)

			mockClassRepo.EXPECT().
				FindOne(testClassID).
				Return(testClass, nil).
				Times(1)

			resultResponse := service.GetOneById(testClassID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("GetManyByUserId", func() {
		It("should call FindAllByUserId", func() {
			testClasses := []s.Class{
				{Base: s.Base{ID: 1}},
				{Base: s.Base{ID: 2}},
			}

			expectedResponse := dtos.NewKresponse(okResponse, testClasses)

			mockClassRepo.EXPECT().
				FindAllByUserId(testUserID).
				Return(testClasses, nil).
				Times(1)

			resultResponse := service.GetManyByUserId(testUserID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})

	})
})
