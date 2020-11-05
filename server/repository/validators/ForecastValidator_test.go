package validators_test

import (
	"errors"
	"strings"
	"time"

	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	r "kadvisor/server/repository"
	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
)

var _ = Describe("ForecastValidator", func() {
	var (
		mockCtrl         *g.Controller
		mockTagValidator *mocks.MockTagValidator
		mockForecastRepo *mocks.MockForecastRepository
		validator        v.ForecastValidator
	)

	BeforeEach(func() {
		mockCtrl = g.NewController(GinkgoT())
		mockTagValidator = mocks.NewMockTagValidator(mockCtrl)
		mockForecastRepo = mocks.NewMockForecastRepository(mockCtrl)

		validator = v.ForecastValidator{
			TagValidator: mockTagValidator,
			ForecastRepo: mockForecastRepo,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("ForecastValidator", func() {
		Context("Constructor", func() {
			It("should return an instance", func() {
				expected := v.ForecastValidator{
					TagValidator: v.TagValidator{},
					ForecastRepo: r.ForecastRepository{},
				}

				Expect(v.NewForecastValidator()).To(Equal(expected))
			})
		})

		Context("Tag Validation", func() {
			It("should call TagValidator and return error list", func() {
				testObj := s.Forecast{}
				errMsg := "test err msg"

				mockForecastRepo.EXPECT().
					FindOne(g.Any(), g.Any(), g.Any()).
					Return(s.Forecast{}, errors.New("error")).
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

		Context("Validation with repository", func() {
			BeforeEach(func() {
				mockTagValidator.EXPECT().ValidateStruct(g.Any()).AnyTimes()
			})

			It("should call ForecastRepository.FindOne and return error if not unique", func() {
				testID := 1
				testYear := time.Now().Year()
				isPreloaded := false
				expectedMsg := "forecast already exists"
				testObj := s.Forecast{
					UserID: testID,
					Year:   testYear,
				}

				mockForecastRepo.EXPECT().
					FindOne(testID, testYear, isPreloaded).
					Return(s.Forecast{}, nil).
					AnyTimes()

				result := validator.Validate(testObj)
				resultErrMsg := result[0].Error()
				Expect(result).To(HaveCap(1))
				Expect(strings.Contains(resultErrMsg, expectedMsg)).To(BeTrue())
			})
		})

		Context("Validate Entries months", func() {
			BeforeEach(func() {
				mockTagValidator.EXPECT().ValidateStruct(g.Any()).AnyTimes()
				mockForecastRepo.EXPECT().FindOne(g.Any(), g.Any(), g.Any()).
					Return(s.Forecast{}, errors.New("error")).
					AnyTimes()
			})

			It("should return error if entries have repeated months", func() {
				testObj := generateTestForecast()
				testObj.Entries[1].Month = 1
				expectedMsg := "repeated month not allowed"

				result := validator.Validate(testObj)
				resultErrMsg := result[0].Error()
				Expect(result).To(HaveCap(1))
				Expect(strings.Contains(resultErrMsg, expectedMsg)).To(BeTrue())
			})
		})
	})
})

func generateTestForecast() s.Forecast {
	fEntries := []s.ForecastEntry{}
	forecast := s.Forecast{
		UserID: 1,
		Year:   time.Now().Year(),
	}

	for i := 1; i < 13; i++ {
		nEntry := s.ForecastEntry{
			Month:   i,
			Income:  float64(i),
			Expense: float64(i),
		}
		fEntries = append(fEntries, nEntry)
	}

	forecast.Entries = fEntries
	return forecast
}
