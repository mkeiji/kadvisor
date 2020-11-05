package validators_test

import (
	"errors"
	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"

	r "kadvisor/server/repository"
	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
)

var _ = Describe("UserValidator", func() {
	var (
		mockCtrl         *g.Controller
		mockTagValidator *mocks.MockTagValidator
		mockLoginRepo    *mocks.MockLoginRepository
		validator        v.UserValidator
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockTagValidator = mocks.NewMockTagValidator(mockCtrl)
		mockLoginRepo = mocks.NewMockLoginRepository(mockCtrl)

		validator = v.UserValidator{
			TagValidator:    mockTagValidator,
			LoginRepository: mockLoginRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := v.UserValidator{
				TagValidator:    v.TagValidator{},
				LoginRepository: r.LoginRepository{},
			}

			Expect(v.NewUserValidator()).To(Equal(expected))
		})
	})

	Describe("Tag validation", func() {
		It("should call TagValidator and return error list", func() {
			errMsg := "test err msg"
			testObj := s.User{
				Login: s.Login{
					Email: "test@email.com",
				},
			}

			mockLoginRepo.EXPECT().
				FindOneByEmail(g.Any()).
				Return(s.Login{}, errors.New("not found")).
				AnyTimes()
			mockTagValidator.EXPECT().
				ValidateStruct(testObj).
				Return(errors.New(errMsg)).
				Times(1)

			result := validator.Validate(testObj)
			Expect(result).To(HaveCap(1))
			Expect(result[0].Error()).To(Equal(errMsg))
		})

	})

	Describe("Validate with repository", func() {
		It(`should call LoginRepository.FindOneByEmail 
            and return error if an email is found`, func() {
			mockTagValidator.EXPECT().ValidateStruct(g.Any()).AnyTimes()

			expectedErrMsg := "email already exists"
			testObj := s.User{
				Login: s.Login{
					Email: "test@email.com",
				},
			}

			mockLoginRepo.EXPECT().
				FindOneByEmail(g.Any()).
				Return(s.Login{}, nil).
				Times(1)

			result := validator.Validate(testObj)
			resultErrMsg := result[0].Error()

			Expect(result).To(HaveCap(1))
			Expect(strings.Contains(resultErrMsg, expectedErrMsg)).To(BeTrue())
		})
	})
})
