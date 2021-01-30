package repository_test

import (
	"database/sql"
	"fmt"
	r "kadvisor/server/repository"
	h "kadvisor/server/repository/RepositoryTestHelper"
	m "kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("ForecastRepository", func() {
	const (
		TEST_ID = 1
		USER_ID = 1
	)

	var (
		gormmockDB  *gorm.DB
		sqlmockDB   *sql.DB
		mockManager sqlmock.Sqlmock
		repo        r.ForecastRepository
		year        int
	)

	BeforeEach(func() {
		year = time.Now().Year()
		gormmockDB, sqlmockDB, mockManager = h.SetupMockDB()
		repo = r.ForecastRepository{
			EntryMapper: m.ForecastEntryMapper{},
			Db:          gormmockDB,
		}
	})

	AfterEach(func() {
		sqlmockDB.Close()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := r.ForecastRepository{
				EntryMapper: m.ForecastEntryMapper{},
				Db:          app.Db,
			}
			Expect(r.NewForecastRepository()).To(Equal(expected))
		})
	})

	Describe("FindOne", func() {
		var (
			expectedForecastQuery      string
			expectedForecastEntryQuery string
		)

		BeforeEach(func() {
			expectedForecastQuery = regexp.QuoteMeta(
				"SELECT * FROM `forecasts` WHERE user_id=? AND year=? ORDER BY `forecasts`.`id` LIMIT 1",
			)

			expectedForecastEntryQuery = regexp.QuoteMeta(
				"SELECT * FROM `forecast_entries` WHERE `forecast_entries`.`forecast_id` = ?",
			)
		})

		Context("No error", func() {
			const (
				month   = 1
				income  = 10.0
				expense = -5.0
			)

			var (
				expectedForecastRow      *sqlmock.Rows
				expectedForecastEntryRow *sqlmock.Rows
				isPreloaded              bool
			)
			BeforeEach(func() {
				expectedForecastRow = sqlmock.
					NewRows([]string{"id", "user_id", "year"}).
					AddRow(TEST_ID, USER_ID, year)
				expectedForecastEntryRow = sqlmock.
					NewRows([]string{"id", "forecast_id", "month", "income", "expense"}).
					AddRow(TEST_ID, TEST_ID, month, income, expense)

				mockManager.ExpectQuery(expectedForecastQuery).
					WillReturnRows(expectedForecastRow)
				mockManager.ExpectQuery(expectedForecastEntryQuery).
					WillReturnRows(expectedForecastEntryRow)
			})

			AfterEach(func() {
				result, err := repo.FindOne(USER_ID, year, isPreloaded)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.Base.ID).Should(Equal(TEST_ID))
				Expect(result.UserID).Should(Equal(TEST_ID))
				Expect(result.Year).Should(Equal(year))
				if isPreloaded == true {
					Expect(len(result.Entries)).Should(Equal(1))
					Expect(result.Entries[0].ForecastID).Should(Equal(TEST_ID))
					Expect(result.Entries[0].Month).Should(Equal(month))
					Expect(result.Entries[0].Income).Should(Equal(income))
					Expect(result.Entries[0].Expense).Should(Equal(expense))
				} else {
					Expect(len(result.Entries)).Should(Equal(0))
				}
			})

			It("preload - should return expected row with entries", func() {
				isPreloaded = true
			})

			It("no preload - should return expected row without entries", func() {
				isPreloaded = false
			})
		})

		Context("Error", func() {
			var isPreloaded bool

			BeforeEach(func() {
				mockManager.ExpectQuery(h.AnySelectQuery()).
					WillReturnRows(sqlmock.NewRows(nil))
			})

			AfterEach(func() {
				_, err := repo.FindOne(USER_ID, year, isPreloaded)
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})

			It("preload - should return error if none is found", func() {
				isPreloaded = true
			})

			It("no preload - should return error if none is found", func() {
				isPreloaded = false
			})
		})
	})

	Describe("Create", func() {
		It("should create one forecast", func() {
			expectedQuery := regexp.QuoteMeta(
				fmt.Sprintf(
					"%v %v",
					"INSERT INTO `forecasts` (`created_at`,`updated_at`,`user_id`,`year`)",
					"VALUES (?,?,?,?)",
				),
			)
			date := time.Now()
			nInsertedID := int64(1)
			nAffectedRows := int64(1)
			testObject := s.Forecast{
				Base:   s.Base{CreatedAt: date, UpdatedAt: date},
				UserID: TEST_ID,
				Year:   year,
			}
			expected := testObject
			expected.Base.ID = 1

			mockManager.ExpectExec(expectedQuery).
				WithArgs(
					testObject.Base.CreatedAt,
					testObject.Base.UpdatedAt,
					testObject.UserID,
					testObject.Year).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Create(testObject)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(expected))
		})
	})

	Describe("Delete", func() {
		expectedQuery := regexp.QuoteMeta(
			"DELETE FROM `forecasts` WHERE `forecasts`.`id` = ?",
		)
		It("should delete one forecast", func() {
			nInsertedID := int64(0)
			nAffectedRows := int64(1)
			testObject := s.Forecast{
				Base:   s.Base{ID: TEST_ID},
				UserID: TEST_ID,
				Year:   year,
			}
			storedForecast := sqlmock.
				NewRows([]string{"id", "user_id", "year"}).
				AddRow(TEST_ID, USER_ID, year)

			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(storedForecast)
			mockManager.ExpectExec(expectedQuery).
				WithArgs(TEST_ID).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Delete(TEST_ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(testObject))
		})
	})
})
