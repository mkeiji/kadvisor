package repository_test

import (
	"database/sql"
	"errors"
	"fmt"
	r "kadvisor/server/repository"
	h "kadvisor/server/repository/RepositoryTestHelper"
	app "kadvisor/server/resources/application"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("ReportRepository", func() {
	const (
		TEST_USER_ID = 1
	)

	var (
		gormmockDB  *gorm.DB
		sqlmockDB   *sql.DB
		mockManager sqlmock.Sqlmock
		repo        r.ReportRepository
		date        time.Time
	)

	BeforeEach(func() {
		date = h.GetTodayUTC()
		gormmockDB, sqlmockDB, mockManager = h.SetupMockDB()
		repo = r.ReportRepository{
			Db: gormmockDB,
		}
	})

	AfterEach(func() {
		sqlmockDB.Close()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := r.ReportRepository{
				Db: app.Db,
			}
			Expect(r.NewReportRepository()).To(Equal(expected))
		})
	})

	Describe("GetAvailableForecastYears", func() {
		var (
			expectedQuery string
			expectedRow   *sqlmock.Rows
		)

		BeforeEach(func() {
			expectedQuery = regexp.QuoteMeta(
				"SELECT DISTINCT `year` FROM `forecasts` ORDER BY year desc",
			)
		})

		It("no error - should return availble forecast years", func() {
			expectedYear := date.Year()
			expectedRow = sqlmock.
				NewRows([]string{"user_id", "year"}).
				AddRow(TEST_USER_ID, expectedYear)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.GetAvailableForecastYears(TEST_USER_ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(result)).Should(Equal(1))
			Expect(result[0]).Should(Equal(expectedYear))
		})

		It("error - should return error if no forecast year is found", func() {
			expectedRow = sqlmock.NewRows([]string{"user_id", "year"})
			expectedError := errors.New("no available years found")

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			_, err := repo.GetAvailableForecastYears(TEST_USER_ID)
			Expect(err).Should(Equal(expectedError))
		})
	})

	Describe("GetAvailableYears", func() {
		var (
			expectedQuery string
			expectedRow   *sqlmock.Rows
		)

		BeforeEach(func() {
			expectedQuery = regexp.QuoteMeta(
				fmt.Sprintf(
					"%v %v %v %v %v %v %v",
					"SELECT DISTINCT year(date) as year",
					"FROM entries",
					fmt.Sprintf("WHERE user_id=%v", TEST_USER_ID),
					"UNION",
					"SELECT DISTINCT year as year",
					"FROM forecasts",
					fmt.Sprintf("WHERE user_id=%v", TEST_USER_ID),
				),
			)
		})

		It("no error - should return available years of forecast or entries", func() {
			expectedYear := date.Year()
			expectedRow = sqlmock.
				NewRows([]string{"year"}).
				AddRow(expectedYear)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.GetAvailableYears(TEST_USER_ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(result)).Should(Equal(1))
			Expect(result[0]).Should(Equal(expectedYear))
		})

		It("error - should return error if no year is found", func() {
			expectedRow = sqlmock.NewRows([]string{"year"})
			expectedError := errors.New("no available years found")

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			_, err := repo.GetAvailableYears(TEST_USER_ID)
			Expect(err).Should(Equal(expectedError))
		})
	})

	Describe("FindBalance", func() {
		var (
			expectedQuery   string
			expectedRow     *sqlmock.Rows
			expectedBalance float64
		)

		BeforeEach(func() {
			expectedBalance = 10.0
			expectedQuery = regexp.QuoteMeta(
				fmt.Sprintf(
					"%v %v %v %v",
					"SELECT user_id as user_id, sum(amount) as balance",
					"FROM `entries`",
					"WHERE user_id=?",
					"GROUP BY `user_id`",
				),
			)
		})

		It("no error - should return balance", func() {
			expectedRow = sqlmock.
				NewRows([]string{"user_id", "balance"}).
				AddRow(TEST_USER_ID, expectedBalance)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.FindBalance(TEST_USER_ID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Balance).Should(Equal(expectedBalance))
		})

		It("error - should return error if not found", func() {
			expectedRow = sqlmock.
				NewRows([]string{"user_id", "balance"}).
				AddRow(0, 0)
			expectedError := errors.New("no balance is available")

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			_, err := repo.FindBalance(TEST_USER_ID)
			Expect(err).Should(Equal(expectedError))
		})
	})

	Describe("FindYearToDateReport", func() {
		var (
			expectedQuery   string
			expectedRow     *sqlmock.Rows
			expectedYear    int
			expectedMonth   int
			expectedIncome  float64
			expectedExpense float64
			expectedBalance float64
			expectedType    string
		)

		BeforeEach(func() {
			expectedYear = date.Year()
			expectedMonth = int(date.Month())
			expectedIncome = 10.0
			expectedExpense = 5.0
			expectedBalance = expectedIncome - expectedExpense
			expectedType = "YTD"
			expectedQuery = regexp.QuoteMeta(fmt.Sprintf(`
				select
					year(date) year,
					month(date) month,
					sum(income) income, 
					sum(expense) expense,
					(sum(income) + sum(expense)) balance
				from (
					select date,
						case when entry_type_code_id='INCOME_ENTRY_TYPE' then amount else 0 end income, 
						case when entry_type_code_id='EXPENSE_ENTRY_TYPE' then amount else 0 end expense 
					from entries
					where user_id=?
						and year(date)=%d
				) yearly
				group by year(date), month(date);
			`, expectedYear))
		})

		It("no error - should return balance", func() {
			expectedRow = sqlmock.
				NewRows([]string{"year", "month", "income", "expense", "balance", "type"}).
				AddRow(expectedYear, expectedMonth, expectedIncome, expectedExpense, expectedBalance, expectedType)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.FindYearToDateReport(TEST_USER_ID, date.Year())
			Expect(err).ShouldNot(HaveOccurred())
			Expect(len(result)).Should(Equal(1))
			Expect(result[0].Year).Should(Equal(expectedYear))
			Expect(result[0].Month).Should(Equal(expectedMonth))
			Expect(result[0].Income).Should(Equal(expectedIncome))
			Expect(result[0].Expense).Should(Equal(expectedExpense))
			Expect(result[0].Balance).Should(Equal(expectedBalance))
			Expect(result[0].Type).Should(Equal(expectedType))
		})

		It("error - should return error if no report is found", func() {
			expectedRow = sqlmock.
				NewRows([]string{"year", "month", "income", "expense", "balance", "type"})
			expectedError := errors.New("no report available")

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			_, err := repo.FindYearToDateReport(TEST_USER_ID, date.Year())
			Expect(err).Should(Equal(expectedError))
		})
	})
})
