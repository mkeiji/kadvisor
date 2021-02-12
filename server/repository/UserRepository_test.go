package repository_test

import (
	"database/sql"
	"errors"
	"fmt"
	r "kadvisor/server/repository"
	h "kadvisor/server/repository/RepositoryTestHelper"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
	"regexp"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"kadvisor/server/resources/enums"
)

var _ = Describe("UserRepository", func() {
	const (
		TEST_ID_1    = 1
		TEST_ID_2    = 2
		TEST_USER_ID = 1
	)

	var (
		gormmockDB    *gorm.DB
		sqlmockDB     *sql.DB
		mockManager   sqlmock.Sqlmock
		repo          r.UserRepository
		today         int64
		yesterday     int64
		nInsertedID   int64
		nAffectedRows int64
		storedUser    *sqlmock.Rows
	)

	BeforeEach(func() {
		today = h.GetTodayUTCUnix()
		yesterday = h.GetYesterdayUTCUnix()
		nInsertedID = int64(1)
		nAffectedRows = int64(1)
		storedUser = createMockUserDbRow(TEST_ID_1, yesterday, yesterday)

		gormmockDB, sqlmockDB, mockManager = h.SetupMockDB()
		repo = r.UserRepository{
			Db: gormmockDB,
		}
	})

	AfterEach(func() {
		sqlmockDB.Close()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := r.UserRepository{
				Db: app.Db,
			}
			Expect(r.NewUserRepository()).To(Equal(expected))
		})
	})

	Describe("FindAll", func() {
		var (
			expectedLoginQuery string
			expectedUserQuery  string
			expectedClassQuery string
			expectedEntryQuery string
			isPreLoaded        bool
			expectedRows       *sqlmock.Rows
		)

		BeforeEach(func() {
			expectedLoginQuery = regexp.QuoteMeta(
				"SELECT * FROM `logins` WHERE `logins`.`user_id` IN (?,?)",
			)
			expectedUserQuery = regexp.QuoteMeta(
				"SELECT * FROM `users`",
			)
			expectedClassQuery = regexp.QuoteMeta(
				"SELECT * FROM `classes` WHERE `classes`.`user_id` IN (?,?)",
			)
			expectedEntryQuery = regexp.QuoteMeta(
				"SELECT * FROM `entries` WHERE `entries`.`user_id` IN (?,?)",
			)
			expectedRows = sqlmock.
				NewRows([]string{"id"}).
				AddRow(TEST_ID_1).
				AddRow(TEST_ID_2)
		})

		Context("No error", func() {
			It("no preload - should query users table and logins", func() {
				isPreLoaded = false

				mockManager.ExpectQuery(expectedUserQuery).
					WillReturnRows(expectedRows)
				mockManager.ExpectQuery(expectedLoginQuery).
					WillReturnRows(expectedRows)

				result, err := repo.FindAll(isPreLoaded)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(result)).Should(Equal(2))
				Expect(result[0].Base.ID).Should(Equal(TEST_ID_1))
				Expect(result[1].Base.ID).Should(Equal(TEST_ID_2))
			})

			It("with preload - should query users, logins, entries and classes tables", func() {
				isPreLoaded = true

				mockManager.ExpectQuery(expectedUserQuery).
					WillReturnRows(expectedRows)
				mockManager.ExpectQuery(expectedClassQuery).
					WillReturnRows(expectedRows)
				mockManager.ExpectQuery(expectedEntryQuery).
					WillReturnRows(expectedRows)
				mockManager.ExpectQuery(expectedLoginQuery).
					WillReturnRows(expectedRows)

				_, err := repo.FindAll(isPreLoaded)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("Error", func() {
			var (
				expectedErr error
			)

			BeforeEach(func() {
				expectedErr = errors.New("records not found")
			})

			AfterEach(func() {
				mockManager.ExpectQuery(expectedUserQuery).
					WillReturnRows(sqlmock.NewRows(nil))

				_, err := repo.FindAll(isPreLoaded)
				Expect(err).Should(Equal(expectedErr))
			})

			It("no preload - should return error if user not found", func() {
				isPreLoaded = false
			})

			It("with preload - should return error if user not found", func() {
				isPreLoaded = true
			})
		})
	})

	Describe("FindOne", func() {
		var (
			expectedQuery string
			expectedRow   *sqlmock.Rows
			testName      string
			isPreLoaded   bool
		)

		BeforeEach(func() {
			testName = "test"
			expectedQuery = regexp.QuoteMeta(
				"SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1",
			)
			expectedRow = sqlmock.
				NewRows([]string{"first_name"}).
				AddRow(testName)
		})

		Context("No error", func() {
			AfterEach(func() {
				mockManager.ExpectQuery(expectedQuery).
					WithArgs(TEST_USER_ID).
					WillReturnRows(expectedRow)

				result, err := repo.FindOne(TEST_USER_ID, isPreLoaded)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result.FirstName).Should(Equal(testName))
			})

			It("no preload - should return user", func() {
				isPreLoaded = false
			})

			It("with preload - should return user", func() {
				isPreLoaded = true
			})
		})

		Context("Error", func() {
			AfterEach(func() {
				mockManager.ExpectQuery(expectedQuery).
					WillReturnRows(sqlmock.NewRows(nil))

				_, err := repo.FindOne(TEST_USER_ID, isPreLoaded)
				Expect(err).Should(Equal(gorm.ErrRecordNotFound))
			})

			It("no preload - should return user", func() {
				isPreLoaded = false
			})

			It("with preload - should return user", func() {
				isPreLoaded = true
			})
		})
	})

	Describe("Create", func() {
		var (
			expectedQuery string
		)

		BeforeEach(func() {
			expectedQuery = regexp.QuoteMeta(
				fmt.Sprintf(
					"%v %v %v",
					"INSERT INTO `users`",
					" (`created_at`,`updated_at`,`is_premium`,`first_name`,`last_name`,`phone`,`address`)",
					"VALUES (?,?,?,?,?,?,?)",
				),
			)
		})

		It("no error - should create a User", func() {
			testIsPremium := false
			testName := "test"
			testLast := "last"
			testPhone := "123"
			testAddr := "testAddr"
			testObj := s.User{
				FirstName: testName,
				LastName:  testLast,
				Phone:     testPhone,
				Address:   testAddr,
			}

			mockManager.ExpectExec(expectedQuery).
				WithArgs(today, today, testIsPremium, testName, testLast, testPhone, testAddr).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Create(testObj)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.FirstName).Should(Equal(testName))
			Expect(result.LastName).Should(Equal(testLast))
			Expect(result.Phone).Should(Equal(testPhone))
			Expect(result.Address).Should(Equal(testAddr))
		})

		It("no error - should create a User", func() {
			today := h.GetTodayUTCUnix()
			testErr := errors.New("test error")
			testObj := s.User{}

			mockManager.ExpectExec(expectedQuery).
				WithArgs(today, today, false, "", "", "", "").
				WillReturnError(testErr)

			_, err := repo.Create(testObj)
			Expect(err).Should(Equal(testErr))
		})
	})

	Describe("Update", func() {
		var (
			expectedFindUserQuery  string
			expectedFindLoginQuery string
			expectedUpdateQuery    string
			storedLogin            *sqlmock.Rows
		)

		BeforeEach(func() {
			expectedFindUserQuery = regexp.QuoteMeta(
				"SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1",
			)
			expectedFindLoginQuery = regexp.QuoteMeta(
				"SELECT * FROM `logins` WHERE `logins`.`user_id` = ?",
			)
			expectedUpdateQuery = regexp.QuoteMeta(
				"UPDATE `users` SET `updated_at`=?,`first_name`=? WHERE `id` = ?",
			)
			storedLogin = createMockLoginDbRow(TEST_ID_1, TEST_USER_ID, time.Unix(yesterday, 0))
		})

		It("no error - should update user", func() {
			updateName := "newName"
			testObj := s.User{FirstName: updateName}

			mockManager.ExpectQuery(expectedFindUserQuery).
				WillReturnRows(storedUser)
			mockManager.ExpectQuery(expectedFindLoginQuery).
				WillReturnRows(storedLogin)
			mockManager.ExpectExec(expectedUpdateQuery).
				WithArgs(h.GetTodayUTCUnix(), updateName, TEST_ID_1).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))
			mockManager.ExpectExec(h.AnyInsertQuery()).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Update(testObj)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.FirstName).Should(Equal(updateName))
		})

		It("error - should return error if cannot find user", func() {
			testObj := s.User{}

			mockManager.ExpectQuery(expectedFindUserQuery).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.Update(testObj)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})

	Describe("Delete", func() {
		It("no error - should return deleted user", func() {
			expectedQuery := regexp.QuoteMeta(
				"DELETE FROM `users` WHERE `users`.`id` = ?",
			)

			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(storedUser)
			mockManager.ExpectExec(expectedQuery).
				WithArgs(TEST_ID_1).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Delete(TEST_ID_1)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Base.ID).Should(Equal(TEST_ID_1))
		})

		It("error - should return error if cannot find user", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.Delete(TEST_ID_1)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})
})

func createMockUserDbRow(
	id int, createdAt int64, updatedAt int64,
) *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"id", "created_at", "updated_at", "is_premium", "first_name", "last_name", "phone", "address",
	}).AddRow(
		id, createdAt, updatedAt, false, "testName", "testLast", "123", "testaddr",
	)
}

func createMockLoginDbRow(id int, userID int, lastLogin time.Time) *sqlmock.Rows {
	return sqlmock.
		NewRows(
			[]string{
				"id", "user_id", "role_id", "email", "user_name", "password", "is_logged_in", "last_login",
			}).
		AddRow(
			id, userID, enums.REGULAR, "testEmail", "testUserName", "testPass", false, lastLogin,
		)
}
