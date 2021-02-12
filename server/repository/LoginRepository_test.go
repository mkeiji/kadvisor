package repository_test

import (
	"database/sql"
	"errors"
	r "kadvisor/server/repository"
	h "kadvisor/server/repository/RepositoryTestHelper"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
	"kadvisor/server/resources/enums"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("LoginRepository", func() {
	const (
		TEST_ID       = 1
		TEST_USER_ID  = 1
		TEST_EMAIL    = "test@email.com"
		TEST_USERNAME = "testUserName"
		TEST_PASSWORD = "testPassword"
	)

	var (
		gormmockDB       *gorm.DB
		sqlmockDB        *sql.DB
		mockManager      sqlmock.Sqlmock
		repo             r.LoginRepository
		yesterday        time.Time
		isLoggedIn       bool
		getTestStoredRow func() *sqlmock.Rows
	)

	BeforeEach(func() {
		yesterday = h.GetYesterdayUTC()
		isLoggedIn = true
		gormmockDB, sqlmockDB, mockManager = h.SetupMockDB()
		repo = r.LoginRepository{
			Db: gormmockDB,
		}
	})

	AfterEach(func() {
		sqlmockDB.Close()
	})

	getTestStoredRow = func() *sqlmock.Rows {
		return sqlmock.
			NewRows(
				[]string{
					"id", "user_id", "role_id", "email", "user_name", "password", "is_logged_in", "last_login",
				}).
			AddRow(
				TEST_ID, TEST_USER_ID, enums.REGULAR, TEST_EMAIL, TEST_USERNAME, TEST_PASSWORD, isLoggedIn, yesterday,
			)
	}

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := r.LoginRepository{
				Db: app.Db,
			}
			Expect(r.NewLoginRepository()).To(Equal(expected))
		})
	})

	Describe("FindOneByEmail", func() {
		var (
			expectedQuery string
			expectedRow   *sqlmock.Rows
		)

		BeforeEach(func() {
			expectedQuery = regexp.QuoteMeta(
				"SELECT * FROM `logins` WHERE email=? ORDER BY `logins`.`id` LIMIT 1",
			)

			expectedRow = getTestStoredRow()
		})

		It("no error - should return expected row", func() {
			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.FindOneByEmail("")
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Base.ID).Should(Equal(TEST_ID))
			Expect(result.UserID).Should(Equal(TEST_USER_ID))
			Expect(result.RoleID).Should(Equal(int(enums.REGULAR)))
			Expect(result.Email).Should(Equal(TEST_EMAIL))
			Expect(result.UserName).Should(Equal(TEST_USERNAME))
			Expect(result.Password).Should(Equal(TEST_PASSWORD))
			Expect(result.IsLoggedIn).Should(Equal(true))
			Expect(result.LastLogin).Should(Equal(&yesterday))
		})

		It("error - should return error if not found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.FindOneByEmail(TEST_EMAIL)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})

	Describe("Update", func() {
		var (
			expectedFindOneQuery string
			expectedQuery        string
			storedRow            *sqlmock.Rows
			updatedEmail         string
			testObject           s.Login
		)

		BeforeEach(func() {
			updatedEmail = "updated@email.com"
			testObject = s.Login{Email: updatedEmail}
			expectedFindOneQuery = regexp.QuoteMeta(
				"SELECT * FROM `logins` WHERE id=? ORDER BY `logins`.`id` LIMIT 1",
			)
			expectedQuery = regexp.QuoteMeta(
				"UPDATE `logins` SET `updated_at`=?,`email`=? WHERE `id` = ?",
			)

			storedRow = getTestStoredRow()
		})

		It("no error - should return expected row", func() {
			nInsertedID := int64(1)
			nAffectedRows := int64(1)

			mockManager.ExpectQuery(expectedFindOneQuery).
				WillReturnRows(storedRow)
			mockManager.ExpectExec(expectedQuery).
				WithArgs(h.GetTodayUTCUnix(), testObject.Email, TEST_ID).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Update(testObject)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Email).Should(Equal(updatedEmail))
		})

		It("error - should return error if none is found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.Update(testObject)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})

	Describe("UpdateLoginStatus", func() {
		var (
			expectedFindQuery string
			expectedQuery     string
			storedRow         *sqlmock.Rows
			testObject        s.Login
		)

		BeforeEach(func() {
			testObject = s.Login{
				Email:      TEST_EMAIL,
				IsLoggedIn: true,
			}
			expectedFindQuery = regexp.QuoteMeta(
				"SELECT * FROM `logins` WHERE email=? ORDER BY `logins`.`id` LIMIT 1",
			)
			expectedQuery = regexp.QuoteMeta(
				"UPDATE `logins` SET `is_logged_in`=?,`updated_at`=? WHERE email=? AND `id` = ?",
			)

			storedRow = getTestStoredRow()
		})

		It("no error - should update isLoggedIn", func() {
			nInsertedID := int64(1)
			nAffectedRows := int64(1)
			expectedIsLoggedIn := false

			mockManager.ExpectQuery(expectedFindQuery).
				WillReturnRows(storedRow)
			mockManager.ExpectExec(expectedQuery).
				WithArgs(expectedIsLoggedIn, h.GetTodayUTCUnix(), TEST_EMAIL, TEST_ID).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.UpdateLoginStatus(testObject, expectedIsLoggedIn)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.IsLoggedIn).Should(BeFalse())
		})

		It("error - should return error if update fail", func() {
			expectedIsLoggedIn := false
			expectedError := errors.New("test error")

			mockManager.ExpectQuery(expectedFindQuery).
				WillReturnRows(storedRow)
			mockManager.ExpectExec(expectedQuery).
				WithArgs(expectedIsLoggedIn, h.GetTodayUTCUnix(), TEST_EMAIL, TEST_ID).
				WillReturnError(expectedError)

			result, err := repo.UpdateLoginStatus(testObject, expectedIsLoggedIn)
			Expect(err).Should(Equal(expectedError))
			Expect(result).Should(Equal(s.Login{}))
		})

		It("error - should return error if findOneByEmail fail", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.Update(testObject)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})
})
