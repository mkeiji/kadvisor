package repository_test

import (
	"database/sql"
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
)

var _ = Describe("ClassRepository", func() {
	var (
		gormmockDB  *gorm.DB
		sqlmockDB   *sql.DB
		mockManager sqlmock.Sqlmock
		repo        r.ClassRepository
	)

	BeforeEach(func() {
		gormmockDB, sqlmockDB, mockManager = h.SetupMockDB()
		repo = r.ClassRepository{Db: gormmockDB}
	})

	AfterEach(func() {
		sqlmockDB.Close()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := r.ClassRepository{Db: app.Db}
			Expect(r.NewClassRepository()).To(Equal(expected))
		})
	})

	Describe("FindAllByUserId", func() {
		var expectedQuery string

		BeforeEach(func() {
			expectedQuery = regexp.QuoteMeta(
				"SELECT * FROM `classes` WHERE `classes`.`user_id` = ?",
			)
		})

		It("should return classes by user id", func() {
			testClasses := []s.Class{
				{
					Base:        s.Base{ID: 1},
					UserID:      1,
					Name:        "test",
					Description: "testDescription",
				},
				{
					Base:        s.Base{ID: 2},
					UserID:      1,
					Name:        "test2",
					Description: "testDescription2",
				},
			}
			expectedRow := sqlmock.
				NewRows([]string{"id", "user_id", "name", "description"}).
				AddRow(testClasses[0].Base.ID, testClasses[0].UserID, testClasses[0].Name, testClasses[0].Description).
				AddRow(testClasses[1].Base.ID, testClasses[1].UserID, testClasses[1].Name, testClasses[1].Description)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.FindAllByUserId(1)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(testClasses))
		})

		It("should return nil if none is found", func() {
			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.FindAllByUserId(1)
			Expect(err).Should(BeNil())
		})
	})

	Describe("FindOne", func() {
		var expectedQuery string

		BeforeEach(func() {
			expectedQuery = regexp.QuoteMeta(
				"SELECT * FROM `classes` WHERE `classes`.`id` = ? ORDER BY `classes`.`id` LIMIT 1",
			)
		})

		It("should return one class", func() {
			testClass := s.Class{
				Base:        s.Base{ID: 1},
				UserID:      1,
				Name:        "test",
				Description: "testDescription",
			}
			expectedRow := sqlmock.
				NewRows([]string{"id", "user_id", "name", "description"}).
				AddRow(testClass.Base.ID, testClass.UserID, testClass.Name, testClass.Description)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.FindOne(1)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(testClass))
		})

		It("should return error if not found", func() {
			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.FindOne(1)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})

	Describe("Create", func() {
		It("should create one class", func() {
			expectedQuery := regexp.QuoteMeta(
				"INSERT INTO `classes` (`created_at`,`updated_at`,`user_id`,`name`,`description`) VALUES (?,?,?,?,?)",
			)
			date := time.Now().Unix()
			nInsertedID := int64(1)
			nAffectedRows := int64(1)
			testClass := s.Class{
				Base:        s.Base{CreatedAt: date, UpdatedAt: date},
				UserID:      1,
				Name:        "test",
				Description: "testd",
			}
			expectedClass := testClass
			expectedClass.Base.ID = 1

			mockManager.ExpectExec(expectedQuery).
				WithArgs(testClass.Base.CreatedAt, testClass.Base.UpdatedAt, testClass.UserID, testClass.Name, testClass.Description).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Create(testClass)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(expectedClass))
		})
	})

	Describe("Update", func() {
		const (
			nInsertedID   = int64(0)
			nAffectedRows = int64(1)
			classID       = 1
			userID        = 1
			className     = "testName"
			description   = "testDescription"
		)

		It("should update one class", func() {
			today := h.GetTodayUTCUnix()
			yesterday := h.GetYesterdayUTCUnix()
			expectedQuery := regexp.QuoteMeta(
				"UPDATE `classes` SET `id`=?,`updated_at`=?,`name`=? WHERE `id` = ?",
			)
			testClass := s.Class{
				Base: s.Base{ID: classID},
				Name: "updatedName",
			}
			expectedClass := s.Class{
				Base:        s.Base{ID: classID, CreatedAt: yesterday, UpdatedAt: today},
				UserID:      userID,
				Name:        "updatedName",
				Description: description,
			}

			storedClass := sqlmock.
				NewRows([]string{"id", "created_at", "updated_at", "user_id", "name", "description"}).
				AddRow(classID, yesterday, yesterday, userID, className, description)

			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(storedClass)
			mockManager.ExpectExec(expectedQuery).
				WithArgs(testClass.Base.ID, today, testClass.Name, classID).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Update(testClass)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(expectedClass))
		})

		It("should not update if not found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.Update(s.Class{})
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})

	Describe("Delete", func() {
		const (
			nInsertedID   = int64(0)
			nAffectedRows = int64(1)
			classID       = 1
			userID        = 1
			className     = "testName"
			description   = "testDescription"
		)

		It("should delete one class", func() {
			expectedQuery := regexp.QuoteMeta(
				"DELETE FROM `classes` WHERE `classes`.`id` = ?",
			)
			yesterday := h.GetYesterdayUTCUnix()
			expectedClass := s.Class{
				Base:        s.Base{ID: classID, CreatedAt: yesterday, UpdatedAt: yesterday},
				UserID:      userID,
				Name:        className,
				Description: description,
			}

			storedClass := sqlmock.
				NewRows([]string{"id", "created_at", "updated_at", "user_id", "name", "description"}).
				AddRow(classID, yesterday, yesterday, userID, className, description)

			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(storedClass)
			mockManager.ExpectExec(expectedQuery).
				WithArgs(classID).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Delete(classID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(expectedClass))
		})

		It("should not delete if not found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.Delete(1)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})
})
