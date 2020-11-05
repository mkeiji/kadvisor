package validators_test

import (
	"errors"
	"strings"

	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	r "kadvisor/server/repository"
	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
)

var _ = Describe("EntryValidator", func() {
	var (
		mockCtrl         *g.Controller
		mockTagValidator *mocks.MockTagValidator
		mockClassRepo    *mocks.MockClassRepository
		mockCodeRepo     *mocks.MockCodeCodeTextRepository
		testEntry        s.Entry
		validator        v.EntryValidator
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockTagValidator = mocks.NewMockTagValidator(mockCtrl)
		mockClassRepo = mocks.NewMockClassRepository(mockCtrl)
		mockCodeRepo = mocks.NewMockCodeCodeTextRepository(mockCtrl)

		testEntry = s.Entry{}
		validator = v.EntryValidator{
			TagValidator:    mockTagValidator,
			ClassRepository: mockClassRepo,
			CodeRepository:  mockCodeRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("EntryValidator", func() {
		Context("Constructor", func() {
			It("should return an instance", func() {
				expected := v.EntryValidator{
					TagValidator:    v.TagValidator{},
					ClassRepository: r.ClassRepository{},
					CodeRepository:  r.CodeCodeTextRepository{},
				}

				Expect(v.NewEntryValidator()).To(Equal(expected))
			})
		})

		Context("Tag validation", func() {
			It("should call TagValidator and return error list", func() {
				tagName := "ispositive"
				errMsg := "test err msg"

				mockClassRepo.EXPECT().FindOne(g.Any()).AnyTimes()
				mockCodeRepo.EXPECT().FindOne(g.Any()).AnyTimes()
				mockTagValidator.EXPECT().
					ValidateStruct(testEntry).
					Return(errors.New(errMsg)).
					Times(1)
				mockTagValidator.EXPECT().
					RegisterTag(tagName, g.Any()).
					Times(1)

				result := validator.Validate(testEntry)
				Expect(result).To(HaveCap(1))
				Expect(result[0].Error()).To(Equal(errMsg))
			})
		})

		Context("Property validation", func() {
			BeforeEach(func() {
				mockTagValidator.EXPECT().ValidateStruct(g.Any()).AnyTimes()
				mockTagValidator.EXPECT().RegisterTag(g.Any(), g.Any()).AnyTimes()
			})

			It("should call ClassRepository.FindOne and return error if not found", func() {
				expectedKey := "Entry.ClassID"
				expectedErrMsg := "Field validation for 'ClassID' invalid classID"

				mockCodeRepo.EXPECT().FindOne(g.Any()).AnyTimes()
				mockClassRepo.EXPECT().FindOne(g.Any()).
					Return(s.Class{}, errors.New(""))

				result := validator.Validate(testEntry)
				resultErrMsg := result[0].Error()

				Expect(result).To(HaveCap(1))
				Expect(strings.Contains(resultErrMsg, expectedKey)).To(BeTrue())
				Expect(strings.Contains(resultErrMsg, expectedErrMsg)).To(BeTrue())
			})

			It("should call CodeRepository.FindOne and return error if not found", func() {
				expectedKey := "Entry.EntryTypeCodeID"
				expectedErrMsg := "Field validation for 'EntryTypeCodeID' invalid lookup"

				mockClassRepo.EXPECT().FindOne(g.Any()).AnyTimes()
				mockCodeRepo.EXPECT().FindOne(g.Any()).
					Return(s.Code{}, errors.New(""))

				result := validator.Validate(testEntry)
				resultErrMsg := result[0].Error()

				Expect(result).To(HaveCap(1))
				Expect(strings.Contains(resultErrMsg, expectedKey)).To(BeTrue())
				Expect(strings.Contains(resultErrMsg, expectedErrMsg)).To(BeTrue())
			})
		})
	})
})
