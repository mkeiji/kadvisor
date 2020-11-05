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

var _ = Describe("LoginValidator", func() {
	var (
		mockCtrl         *g.Controller
		mockTagValidator *mocks.MockTagValidator
		mockLoginRepo    *mocks.MockLoginRepository
		validator        v.LoginValidator
		testEmail        string
		testPassword     string
		testLogin        s.Login
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockTagValidator = mocks.NewMockTagValidator(mockCtrl)
		mockLoginRepo = mocks.NewMockLoginRepository(mockCtrl)

		testEmail = "test@email.com"
		testPassword = "password"
		testLogin = s.Login{
			Email:    testEmail,
			Password: testPassword,
		}

		validator = v.LoginValidator{
			TagValidator:    mockTagValidator,
			LoginRepository: mockLoginRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := v.LoginValidator{
				TagValidator:    v.TagValidator{},
				LoginRepository: r.LoginRepository{},
			}

			Expect(v.NewLoginValidator()).To(Equal(expected))
		})
	})

	Describe("Validation with repository", func() {
		It("should call LoginRepository.FindOneByEmail and return error if not found", func() {
			emptyObj := s.Login{}
			testEmail := "test@email.com"
			expectedMsg := "invalid email"

			mockLoginRepo.EXPECT().
				FindOneByEmail(testEmail).
				Return(emptyObj, errors.New("not found")).
				Times(1)

			result := validator.Validate(testLogin)
			resultErrMsg := result[0].Error()
			Expect(result).To(HaveCap(1))
			Expect(strings.Contains(resultErrMsg, expectedMsg)).To(BeTrue())
		})
	})

	Describe("Password validation", func() {
		It("should call KeiPassUtil and return an error if pwd is invalid", func() {
			expectedMsg := "wrong password"
			testStoredLogin := s.Login{
				Email:    testEmail,
				Password: "differentPwd",
			}

			mockLoginRepo.EXPECT().
				FindOneByEmail(g.Any()).
				Return(testStoredLogin, nil).
				Times(1)

			result := validator.Validate(testLogin)
			resultErrMsg := result[0].Error()
			Expect(result).To(HaveCap(1))
			Expect(strings.Contains(resultErrMsg, expectedMsg)).To(BeTrue())
		})
	})
})
