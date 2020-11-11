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

var _ = Describe("EntryService", func() {
	const (
		okResponse = 200
	)

	var (
		mockCtrl      *g.Controller
		mockEntryRepo *mocks.MockEntryRepository
		service       svc.EntryService
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockEntryRepo = mocks.NewMockEntryRepository(mockCtrl)

		service = svc.EntryService{
			Repository: mockEntryRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := svc.EntryService{
				Repository: r.EntryRepository{},
			}

			Expect(svc.NewEntryService()).To(Equal(expected))
		})
	})

	Describe("GetManyByUserId", func() {
		It("should call Repository.FindAllByUserId", func() {
			testUserID := 1
			testLimit := 1
			testEntries := []s.Entry{}
			expectedResponse := dtos.NewKresponse(okResponse, testEntries)

			mockEntryRepo.EXPECT().
				FindAllByUserId(g.Any(), g.Any()).
				Return(testEntries, nil).
				Times(1)

			resultResponse := service.GetManyByUserId(testUserID, testLimit)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("GetManyByClassId", func() {
		It("should call Repository.FindAllByClassId", func() {
			testClassID := 1
			testLimit := 1
			testEntries := []s.Entry{}
			expectedResponse := dtos.NewKresponse(okResponse, testEntries)

			mockEntryRepo.EXPECT().
				FindAllByUserId(g.Any(), g.Any()).
				Return(testEntries, nil).
				Times(1)

			resultResponse := service.GetManyByUserId(testClassID, testLimit)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("GetOneById", func() {
		It("should call Repository.FindOne", func() {
			testEntryID := 1
			testEntry := s.Entry{Base: s.Base{ID: testEntryID}}
			expectedResponse := dtos.NewKresponse(okResponse, testEntry)

			mockEntryRepo.EXPECT().
				FindOne(testEntryID).
				Return(testEntry, nil).
				Times(1)

			resultResponse := service.GetOneById(testEntryID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("Post", func() {
		It("should call Repository.Create", func() {
			testDescription := "test"
			testEntry := s.Entry{Description: testDescription}
			testCreatedEntry := s.Entry{
				Base:        s.Base{ID: 1},
				Description: testDescription,
			}
			expectedResponse := dtos.NewKresponse(okResponse, testCreatedEntry)

			mockEntryRepo.EXPECT().
				Create(testEntry).
				Return(testCreatedEntry, nil).
				Times(1)

			resultResponse := service.Post(testEntry)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("Put", func() {
		It("should call Repository.Update", func() {
			testID := s.Base{ID: 1}
			testDescription := "updated"
			testEntry := s.Entry{
				Base:        testID,
				Description: testDescription,
			}
			expectedResponse := dtos.NewKresponse(okResponse, testEntry)

			mockEntryRepo.EXPECT().
				Update(testEntry).
				Return(testEntry, nil).
				Times(1)

			resultResponse := service.Put(testEntry)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})

	Describe("Delete", func() {
		It("should call Repository.Delete", func() {
			testEntryID := 1
			expectedResponse := dtos.NewKresponse(okResponse, testEntryID)

			mockEntryRepo.EXPECT().
				Delete(testEntryID).
				Return(testEntryID, nil).
				Times(1)

			resultResponse := service.Delete(testEntryID)
			Expect(resultResponse).To(Equal(expectedResponse))
		})
	})
})
