package repository_test

import (
	"database/sql"
	r "kadvisor/server/repository"
	h "kadvisor/server/repository/RepositoryTestHelper"
	m "kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("ForecastEntryRepository", func() {
	const (
		TEST_ID          = 1
		TEST_FORECAST_ID = 1
		TEST_MONTH       = 1
		TEST_INCOME      = 10.0
		TEST_EXPENSE     = 5.0
	)

	var (
		gormmockDB  *gorm.DB
		sqlmockDB   *sql.DB
		mockManager sqlmock.Sqlmock
		repo        r.ForecastEntryRepository
		testObject  s.ForecastEntry
	)

	BeforeEach(func() {
		gormmockDB, sqlmockDB, mockManager = h.SetupMockDB()
		repo = r.ForecastEntryRepository{
			Mapper: m.ForecastEntryMapper{},
			Db:     gormmockDB,
		}

		testObject = s.ForecastEntry{
			Base:       s.Base{ID: TEST_ID},
			ForecastID: TEST_FORECAST_ID,
			Month:      TEST_MONTH,
			Income:     TEST_INCOME,
			Expense:    TEST_EXPENSE,
		}
	})

	AfterEach(func() {
		sqlmockDB.Close()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := r.ForecastEntryRepository{
				Mapper: m.ForecastEntryMapper{},
				Db:     app.Db,
			}
			Expect(r.NewForecastEntryRepository()).To(Equal(expected))
		})
	})

	Describe("FindOne", func() {
		var expectedQuery string

		BeforeEach(func() {
			expectedQuery = regexp.QuoteMeta(
				"SELECT * FROM `forecast_entries` WHERE id=? ORDER BY `forecast_entries`.`id` LIMIT 1",
			)
		})

		It("should return expected row", func() {
			expectedRow := sqlmock.
				NewRows([]string{"id", "forecast_id", "month", "income", "expense"}).
				AddRow(TEST_ID, TEST_FORECAST_ID, TEST_MONTH, TEST_INCOME, TEST_EXPENSE)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.FindOne(TEST_ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(testObject))
		})

		It("should return error if none is found", func() {
			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.FindOne(TEST_ID)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})

	Describe("Update", func() {
		var expectedQuery string

		BeforeEach(func() {
			expectedQuery = regexp.QuoteMeta(
				"UPDATE `forecast_entries` SET `income`=?,`expense`=? WHERE `id` = ?",
			)
		})

		It("should return updated record", func() {
			newIncomeValue := 20.0
			updatedObject := testObject
			updatedObject.Income = newIncomeValue
			mappedExpense := -1 * TEST_EXPENSE
			expectedObject := updatedObject
			expectedObject.Expense = mappedExpense

			nInsertedID := int64(1)
			nAffectedRows := int64(1)

			storedEntry := sqlmock.
				NewRows([]string{"id", "forecast_id", "month", "income", "expense"}).
				AddRow(TEST_ID, TEST_FORECAST_ID, TEST_MONTH, TEST_INCOME, mappedExpense)

			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(storedEntry)
			mockManager.ExpectExec(expectedQuery).
				WithArgs(
					updatedObject.Income,
					mappedExpense,
					updatedObject.Base.ID).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Update(updatedObject)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(expectedObject))
		})

		It("should return error if none is found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.Update(testObject)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})
})
