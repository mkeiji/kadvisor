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

var _ = Describe("EntryRepository", func() {
	const (
		userID  = 1
		entryID = 1
		classID = 1
	)

	var (
		gormmockDB  *gorm.DB
		sqlmockDB   *sql.DB
		mockManager sqlmock.Sqlmock
		repo        r.EntryRepository
		today       time.Time
		yesterday   time.Time
	)

	BeforeEach(func() {
		today = h.GetTodayUTC()
		yesterday = h.GetYesterdayUTC()
		gormmockDB, sqlmockDB, mockManager = h.SetupMockDB()
		repo = r.EntryRepository{
			Db:     gormmockDB,
			Mapper: m.EntryMapper{},
		}
	})

	AfterEach(func() {
		sqlmockDB.Close()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := r.EntryRepository{
				Db:     app.Db,
				Mapper: m.EntryMapper{},
			}
			Expect(r.NewEntryRepository()).To(Equal(expected))
		})
	})

	Describe("FindAllByUserId / FindAllByClassId", func() {
		const (
			limit = 2
		)

		var (
			expectedUserIdQuery  string
			expectedClassIdQuery string
		)

		BeforeEach(func() {
			expectedUserIdQuery = regexp.QuoteMeta(
				getFindAllSelectQuery("user_id", limit),
			)
			expectedClassIdQuery = regexp.QuoteMeta(
				getFindAllSelectQuery("class_id", limit),
			)
		})

		Context("Found", func() {
			var (
				testEntries  []s.Entry
				expectedRows *sqlmock.Rows
				result       []s.Entry
				err          error
			)

			BeforeEach(func() {
				testEntries = getTestEntries(userID, classID, today)
				expectedRows = sqlmock.
					NewRows([]string{"id", "user_id", "class_id", "entry_type_code_id", "date", "amount", "description", "obs"}).
					AddRow(testEntries[0].Base.ID, testEntries[0].UserID, testEntries[0].ClassID, testEntries[0].EntryTypeCodeID, testEntries[0].Date, testEntries[0].Amount, testEntries[0].Description, testEntries[0].Obs).
					AddRow(testEntries[1].Base.ID, testEntries[1].UserID, testEntries[1].ClassID, testEntries[1].EntryTypeCodeID, testEntries[1].Date, testEntries[1].Amount, testEntries[1].Description, testEntries[1].Obs)
			})

			AfterEach(func() {
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).Should(Equal(testEntries))
			})

			It("UserID - should return entries", func() {
				mockManager.ExpectQuery(expectedUserIdQuery).
					WillReturnRows(expectedRows)

				result, err = repo.FindAllByUserId(userID, limit)
			})

			It("ClassID - should return entries", func() {
				mockManager.ExpectQuery(expectedClassIdQuery).
					WillReturnRows(expectedRows)

				result, err = repo.FindAllByClassId(classID, limit)
			})
		})

		Context("Not found", func() {
			var (
				err    error
				result []s.Entry
			)

			BeforeEach(func() {
				mockManager.ExpectQuery(h.AnySelectQuery()).
					WillReturnRows(sqlmock.NewRows(nil))
			})

			AfterEach(func() {
				Expect(err).Should(BeNil())
				Expect(result).Should(Equal([]s.Entry{}))
			})

			It("UserID - should return nil if none is found", func() {
				result, err = repo.FindAllByUserId(userID, limit)
			})

			It("ClassID - should return nil if none is found", func() {
				result, err = repo.FindAllByClassId(userID, limit)
			})
		})
	})

	Describe("FindOne", func() {
		const (
			entryID = 1
			limit   = 1
		)
		var (
			expectedQuery string
		)

		BeforeEach(func() {
			expectedQuery = regexp.QuoteMeta(
				"SELECT * FROM `entries` WHERE id=? ORDER BY `entries`.`id` LIMIT 1",
			)
		})

		It("should return one Entry", func() {
			testEntry := getTestEntries(userID, classID, today)[0]
			expectedRow := sqlmock.
				NewRows([]string{"id", "user_id", "class_id", "entry_type_code_id", "date", "amount", "description", "obs"}).
				AddRow(testEntry.Base.ID, testEntry.UserID, testEntry.ClassID, testEntry.EntryTypeCodeID, testEntry.Date, testEntry.Amount, testEntry.Description, testEntry.Obs)

			mockManager.ExpectQuery(expectedQuery).
				WillReturnRows(expectedRow)

			result, err := repo.FindOne(entryID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(testEntry))
		})

		It("should return error if not found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.FindOne(entryID)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})

	Describe("Create", func() {
		It("should create one entry", func() {
			expectedQuery := regexp.QuoteMeta(
				"INSERT INTO `entries` (`created_at`,`updated_at`,`user_id`,`class_id`,`entry_type_code_id`,`date`,`amount`,`description`,`obs`) VALUES (?,?,?,?,?,?,?,?,?)",
			)
			nInsertedID := int64(1)
			nAffectedRows := int64(1)
			testEntry := getTestEntries(userID, classID, today)[0]
			testEntry.Base = s.Base{CreatedAt: today, UpdatedAt: today}

			expectedEntry := testEntry
			expectedEntry.Base.ID = 1

			mockManager.ExpectExec(expectedQuery).
				WithArgs(
					testEntry.Base.CreatedAt,
					testEntry.Base.UpdatedAt,
					testEntry.UserID,
					testEntry.ClassID,
					testEntry.EntryTypeCodeID,
					testEntry.Date,
					testEntry.Amount,
					testEntry.Description,
					testEntry.Obs,
				).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Create(testEntry)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(expectedEntry))
		})
	})

	Describe("Update", func() {
		const (
			nInsertedID   = int64(0)
			nAffectedRows = int64(1)
			entryID       = 1
		)

		Context("entry found", func() {
			var (
				descriptionUpdate         string
				expectedUpdateQuery       string
				expectedAmountUpdateQuery string
				testEntry                 s.Entry
				expectedEntry             s.Entry
				storedEntry               *sqlmock.Rows
			)

			BeforeEach(func() {
				descriptionUpdate = "updatedDescription"
				expectedUpdateQuery = regexp.QuoteMeta(
					"UPDATE `entries` SET `id`=?,`updated_at`=?,`description`=? WHERE `id` = ?",
				)
				expectedAmountUpdateQuery = regexp.QuoteMeta(
					"UPDATE `entries` SET `amount`=? WHERE `id` = ?",
				)
				testEntry = s.Entry{
					Base:        s.Base{ID: entryID},
					Description: descriptionUpdate,
				}
				expectedEntry = getTestEntries(userID, classID, today)[0]
				expectedEntry.Base.CreatedAt = yesterday
				expectedEntry.Base.UpdatedAt = today
				expectedEntry.Description = descriptionUpdate
				storedEntry = sqlmock.
					NewRows([]string{"id", "created_at", "updated_at", "user_id", "class_id", "entry_type_code_id", "date", "amount", "description", "obs"}).
					AddRow(
						entryID,
						yesterday,
						yesterday,
						userID,
						classID,
						"testEntryTypeCodeID",
						today,
						10,
						"testDescription",
						"testObs",
					)
			})

			It("no amount - should update entry property and set amount to zero", func() {
				noAmount := 0
				testEntry.Amount = float64(noAmount)
				expectedEntry.Amount = float64(noAmount)

				mockManager.ExpectQuery(h.AnySelectQuery()).
					WillReturnRows(storedEntry)
				mockManager.ExpectExec(expectedUpdateQuery).
					WithArgs(entryID, today, testEntry.Description, entryID).
					WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))
				mockManager.ExpectExec(expectedAmountUpdateQuery).
					WithArgs(noAmount, entryID).
					WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))
				result, err := repo.Update(testEntry)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(result).Should(Equal(expectedEntry))
			})
		})

		It("should not update if not found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.Update(s.Entry{})
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})

	Describe("Delete", func() {
		const (
			nInsertedID   = int64(0)
			nAffectedRows = int64(1)
		)

		It("should delete one entry", func() {
			expectedQuery := regexp.QuoteMeta(
				"DELETE FROM `entries` WHERE `entries`.`id` = ?",
			)
			yesterday := h.GetYesterdayUTC()
			storedEntry := sqlmock.
				NewRows([]string{"id", "created_at", "updated_at", "user_id", "class_id", "entry_type_code_id", "date", "amount", "description", "obs"}).
				AddRow(
					entryID,
					yesterday,
					yesterday,
					userID,
					classID,
					"testEntryTypeCodeID",
					today,
					10,
					"testDescription",
					"testObs",
				)

			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(storedEntry)
			mockManager.ExpectExec(expectedQuery).
				WithArgs(entryID).
				WillReturnResult(sqlmock.NewResult(nInsertedID, nAffectedRows))

			result, err := repo.Delete(classID)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(Equal(entryID))
		})

		It("should not delete if not found", func() {
			mockManager.ExpectQuery(h.AnySelectQuery()).
				WillReturnRows(sqlmock.NewRows(nil))

			_, err := repo.Delete(1)
			Expect(err).Should(Equal(gorm.ErrRecordNotFound))
		})
	})
})

func getTestEntries(userID int, classID int, date time.Time) []s.Entry {
	userid := 1
	classid := 1

	if userID != 0 {
		userid = userID
	}
	if classID != 0 {
		classid = classID
	}

	return []s.Entry{
		{
			Base:            s.Base{ID: 1},
			UserID:          userid,
			ClassID:         classid,
			EntryTypeCodeID: "testEntryTypeCodeID",
			Date:            date,
			Amount:          float64(10),
			Description:     "testDescription",
			Obs:             "testObs",
		},
		{
			Base:            s.Base{ID: 2},
			UserID:          userid,
			ClassID:         classid,
			EntryTypeCodeID: "testEntryTypeCodeID2",
			Date:            date,
			Amount:          float64(11),
			Description:     "testDescription2",
			Obs:             "testObs2",
		},
	}
}

func getFindAllSelectQuery(property string, limit int) string {
	return fmt.Sprintf(
		"SELECT * FROM `entries` WHERE `entries`.`%v` = ? ORDER BY created_at desc LIMIT %v",
		property,
		limit,
	)
}
