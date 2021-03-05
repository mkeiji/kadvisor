package KeiGenUtil_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/resources/enums"
)

var _ = Describe("KeiGenUtil", func() {
	Describe("IntToUint", func() {
		It("should convert int to Uint", func() {
			intNum := 5
			result := KeiGenUtil.IntToUint(intNum)
			resultType := []interface{}{result}

			_, isUint := resultType[0].(uint)
			Expect(isUint).To(BeTrue())
		})
	})

	Describe("UintToInt", func() {
		It("should convert Uint to int", func() {
			uIntNum := uint(5)
			result := KeiGenUtil.UintToInt(uIntNum)
			resultType := []interface{}{result}

			_, isUint := resultType[0].(int)
			Expect(isUint).To(BeTrue())
		})
	})

	Describe("Contains", func() {
		It("should return true if string is present", func() {
			container := []string{"container", "should", "include", "string"}
			strWord := "include"

			isPresent := KeiGenUtil.Contains(container, strWord)
			Expect(isPresent).To(BeTrue())
		})
	})

	Describe("Find", func() {
		It("should return the index of the string in the array", func() {
			container := []string{"container", "should", "include", "string"}
			strWord := "include"
			expectedIndex := 2

			result := KeiGenUtil.Find(container, strWord)
			Expect(result).To(Equal(expectedIndex))
		})
	})

	Describe("HasPermission", func() {
		It("should return true if role is within permitted level", func() {
			role1 := int(enums.REGULAR)
			role2 := int(enums.ADMIN)
			permissionLevel := enums.ADMIN

			role1HasPermission := KeiGenUtil.HasPermission(role1, permissionLevel)
			role2HasPermission := KeiGenUtil.HasPermission(role2, permissionLevel)

			Expect(role1HasPermission).To(BeFalse())
			Expect(role2HasPermission).To(BeTrue())
		})
	})

	Describe("MapErrList", func() {
		It("should convert a list of error to a string map", func() {
			errList := []error{
				errors.New("error1"),
				errors.New("error2"),
			}
			expected := []map[string]interface{}{
				{"error": "error1"},
				{"error": "error2"},
			}

			result := KeiGenUtil.MapErrList(errList)
			Expect(result).To(Equal(expected))
		})
	})

	Describe("MapErrorMsg", func() {
		It("should convert an error to a string map", func() {
			err := errors.New("error")
			expected := map[string]interface{}{
				"error": "error",
			}

			result := KeiGenUtil.MapErrorMsg(err)
			Expect(result).To(Equal(expected))
		})
	})

	Describe("GetErrorList", func() {
		It("should convert a string map to a list of errors", func() {
			sMap := []map[string]interface{}{
				{"error": "error1"},
				{"error": "error2"},
			}
			expected := []error{
				errors.New("error1"),
				errors.New("error2"),
			}

			result := KeiGenUtil.GetErrorList(sMap)
			Expect(result).To(Equal(expected))
		})
	})

	Describe("GetError", func() {
		It("should convert a string map to an error", func() {
			sMap := map[string]interface{}{
				"error": "error",
			}
			expected := errors.New("error")

			result := KeiGenUtil.GetError(sMap)
			Expect(result).To(Equal(expected))
		})
	})

	Describe("GetNow", func() {
		It("should return NOW in RFC3339 format", func() {
			expectedDate, _ := time.Parse(time.RFC3339, time.Now().UTC().Format(time.RFC3339))
			result := KeiGenUtil.GetNow()
			Expect(result).To(Equal(expectedDate))
		})
	})

	Describe("DateToUTCISO8601", func() {
		It("should return date in UTC with RFC3339 format", func() {
			date := time.Now()
			expectedDate, _ := time.Parse(time.RFC3339, date.UTC().Format(time.RFC3339))
			result := KeiGenUtil.DateToUTCISO8601(date)
			Expect(result).To(Equal(expectedDate))
		})
	})

	Describe("IsError", func() {
		It("should return true if interface has type error", func() {
			err := errors.New("error")
			result := KeiGenUtil.IsError(err)
			Expect(result).To(BeTrue())
		})
	})

	Describe("IsOKresponse", func() {
		It("should return true if code is 200", func() {
			code := 200
			result := KeiGenUtil.IsOKresponse(code)
			Expect(result).To(BeTrue())
		})
	})

	Describe("RandomString", func() {
		It("should generate a random string with given size", func() {
			expectedSize := 20
			result := KeiGenUtil.RandomString(expectedSize)
			Expect(len(result)).To(Equal(expectedSize))
		})
	})

	Describe("ToString", func() {
		It("should return obj as string", func() {
			intObj := 5
			expected := "5"
			result := KeiGenUtil.ToString(intObj)
			Expect(result).To(Equal(expected))
		})
	})
})
