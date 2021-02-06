package RepositoryTestHelper

import (
	"database/sql"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupMockDB() (*gorm.DB, *sql.DB, sqlmock.Sqlmock) {
	sqlmockDB, mockManager, err := sqlmock.New()
	Expect(err).ShouldNot(HaveOccurred())

	mockDB, _ := gorm.Open(
		mysql.Dialector{
			Config: &mysql.Config{
				DriverName:                "mysql",
				Conn:                      sqlmockDB,
				SkipInitializeWithVersion: true,
			},
		},
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
			NowFunc: func() time.Time {
				d := (60 * time.Second)
				return time.Now().UTC().Truncate(d)
			},
		},
	)

	return mockDB, sqlmockDB, mockManager
}

func AnySelectQuery() string {
	return `^SELECT+`
}

func GetTodayUTC() time.Time {
	d := (60 * time.Second)
	return time.Now().UTC().Truncate(d)
}

func GetYesterdayUTC() time.Time {
	minusADay := -24 * time.Hour
	d := (60 * time.Second)
	return time.Now().UTC().Add(minusADay).Truncate(d)
}
