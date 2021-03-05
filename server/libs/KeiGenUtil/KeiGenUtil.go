package KeiGenUtil

import (
	"errors"
	"fmt"
	"kadvisor/server/resources/enums"
	"math/rand"
	"strconv"
	"time"
)

func IntToUint(n int) uint {
	var result uint
	base := 10
	bitSize := 64
	strID := strconv.Itoa(n)
	if u64, err := strconv.ParseUint(strID, base, bitSize); err == nil {
		result = uint(u64)
	}
	return result
}

func UintToInt(n uint) int {
	var result int
	base := 16
	bitSize := 64
	strN := strconv.FormatUint(uint64(n), base)
	if i64, err := strconv.ParseInt(strN, base, bitSize); err == nil {
		result = int(i64)
	}
	return result
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func HasPermission(roleID int, permissionLevel enums.RoleEnum) bool {
	if enums.RoleEnum(roleID) <= permissionLevel {
		return true
	} else {
		return false
	}
}

func MapErrList(errList []error) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, err := range errList {
		result = append(result, MapErrorMsg(err))
	}
	return result
}

func MapErrorMsg(err error) map[string]interface{} {
	return map[string]interface{}{"error": err.Error()}
}

func GetErrorList(errListMap []map[string]interface{}) []error {
	var errList []error
	for _, m := range errListMap {
		for _, v := range m {
			errList = append(errList, errors.New(fmt.Sprint(v)))
		}
	}
	return errList
}

func GetError(errMap map[string]interface{}) error {
	return errors.New(fmt.Sprint(errMap["error"]))
}

func GetNow() time.Time {
	return DateToUTCISO8601(time.Now())
}

func DateToUTCISO8601(date time.Time) time.Time {
	result, err := time.Parse(time.RFC3339, date.UTC().Format(time.RFC3339))
	if err != nil {
		panic(err.Error)
	}
	return result
}

func IsError(interfc interface{}) bool {
	_, isError := interfc.(error)
	return isError
}

func IsOKresponse(code int) bool {
	return code == 200
}

func RandomString(n int) string {
	var letters = []rune(
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	)

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func ToString(obj interface{}) string {
	return fmt.Sprintf("%v", obj)
}

func ConsoleLog(obj interface{}) {
	fmt.Printf("\nLog: %v\n\n", obj)
}
