package repository_test

import (
	"database/sql"
	"errors"
	r "kadvisor/server/repository"
	h "kadvisor/server/repository/RepositoryTestHelper"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
)

var _ = Describe("CodeCodeTextRepository", func() {
	var (
		gormmockDB  *gorm.DB
		sqlmockDB   *sql.DB
		mockManager sqlmock.Sqlmock
		repo        r.CodeCodeTextRepository
	)

	BeforeEach(func() {
		gormmockDB, sqlmockDB, mockManager = h.SetupMockDB()
		repo = r.CodeCodeTextRepository{Db: gormmockDB}
	})

	AfterEach(func() {
		sqlmockDB.Close()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := r.CodeCodeTextRepository{Db: app.Db}
			Expect(r.NewCodeCodeTextRepository()).To(Equal(expected))
		})
	})

	Describe("FindAllByCodeGroup", func() {
		var (
			expectedQuery string
			testCodeGroup string
		)

		BeforeEach(func() {
			testCodeGroup = "testCodeGroup"
			expectedQuery = regexp.QuoteMeta(
				"SELECT * FROM `codes` WHERE code_group=?",
			)
		})

		It("should return Code by codeGroup", func() {
			testCodes := []s.Code{
				{
					Base:       s.Base{ID: 1},
					CodeTypeID: "testType",
					CodeGroup:  testCodeGroup,
					Name:       "testName1",
				},
				{
					Base:       s.Base{ID: 2},
					CodeTypeID: "testType",
					CodeGroup:  testCodeGroup,
					Name:       "testName1",
				},
			}
			expectedRow := sqlmock.
				NewRows([]string{"id", "code_type_id", "code_group", "name"}).
				AddRow(testCodes[0].Base.ID, testCodes[0].CodeTypeID, testCodes[0].CodeGroup, testCodes[0].Name).
				AddRow(testCodes[1].Base.ID, testCodes[1].CodeTypeID, testCodes[1].CodeGroup, testCodes[1].Name)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.FindAllByCodeGroup(testCodeGroup)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(testCodes))
		})

		It("should return error if none is found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.FindAllByCodeGroup(testCodeGroup)
			Expect(err).Should(Equal(errors.New("code_group not found")))
		})
	})

	Describe("FindOne", func() {
		var (
			expectedQuery  string
			testCodeTypeID string
		)

		BeforeEach(func() {
			testCodeTypeID = "testCodeTypeID"
			expectedQuery = regexp.QuoteMeta(
				"SELECT * FROM `codes` WHERE code_type_id=? ORDER BY `codes`.`id` LIMIT 1",
			)
		})

		It("should return Code by codeTypeID", func() {
			testCode := s.Code{
				Base:       s.Base{ID: 1},
				CodeTypeID: "testType",
				CodeGroup:  "testCodeGroup",
				Name:       "testName1",
			}

			expectedRow := sqlmock.
				NewRows([]string{"id", "code_type_id", "code_group", "name"}).
				AddRow(testCode.Base.ID, testCode.CodeTypeID, testCode.CodeGroup, testCode.Name)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.FindOne(testCodeTypeID)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(testCode))
		})

		It("should return error if none is found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.FindOne(testCodeTypeID)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})
})
