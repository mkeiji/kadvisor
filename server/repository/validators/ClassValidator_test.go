package validators_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
)

var _ = Describe("ClassValidator", func() {
	var (
		mockCtrl         *gomock.Controller
		mockTagValidator *mocks.MockTagValidator
		testClass        s.Class
		validator        v.ClassValidator
		errMsg           string
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockTagValidator = mocks.NewMockTagValidator(mockCtrl)

		errMsg = "test err msg"
		testClass = s.Class{}
		validator = v.ClassValidator{
			TagValidator: mockTagValidator,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("ClassValidator", func() {
		Context("Constructor", func() {
			It("should return an instance", func() {
				expected := v.ClassValidator{
					TagValidator: v.TagValidator{},
				}

				Expect(v.NewClassValidator()).To(Equal(expected))
			})

		})

		Context("Validate", func() {
			It("should call TagValidator and return error list", func() {
				mockTagValidator.EXPECT().
					ValidateStruct(testClass).
					Return(errors.New(errMsg)).
					Times(1)

				result := validator.Validate(testClass)
				Expect(result).To(HaveCap(1))
				Expect(result[0].Error()).To(Equal(errMsg))
			})
		})
	})
})
