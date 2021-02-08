package application

import (
	"bytes"
	"fmt"
)

func BuildMysqlDbConnection(
	dbHost string,
	dbName string,
	dbUser string,
	dbPass string,
) string {
	var dbConnect bytes.Buffer
	dbAddress := getDbAddress(dbHost, dbName)
	dbConnect.WriteString(dbUser + ":")
	dbConnect.WriteString(dbPass)
	dbConnect.WriteString(dbAddress)

	return dbConnect.String()
}

func getDbAddress(dbHost string, dbName string) string {
	return fmt.Sprintf(
		"@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbHost,
		dbName,
	)
}
