package KeiGenUtil

import (
	"kadvisor/server/resources/enums"
	"strconv"
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
