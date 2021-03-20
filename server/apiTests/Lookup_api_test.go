package apiTests_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/dtos"
)

var _ = Describe("LookupApi", func() {
	const (
		LOOKUP_ENDPOINT = "/api/kadvisor/:uid/lookup"
		CODE_GROUP_TYPE = "EntryTypeCodeID"
		INCOME_TYPE     = "INCOME_ENTRY_TYPE"
		EXPENSE_TYPE    = "EXPENSE_ENTRY_TYPE"
	)

	Describe("GetLookup", func() {
		Context("No error", func() {
			It("should return lookups by codeGroup with ok response", func() {
				expectedCode := http.StatusOK

				params := map[string]string{"codeGroup": CODE_GROUP_TYPE}
				req := kRequestWithParamAndUser("GET", LOOKUP_ENDPOINT, nil, testUserAdmin, params)
				resp := kSendAndAssert(req, expectedCode)

				var result []dtos.LookupEntry
				kReadBody(resp, &result)
				Expect(len(result)).To(Equal(2))
				Expect(result[0].Code).To(Equal(INCOME_TYPE))
				Expect(result[1].Code).To(Equal(EXPENSE_TYPE))
			})
		})

		Context("Error", func() {
			It("should return error with bad request response", func() {
				expectedCode := http.StatusBadRequest
				expectedErrMsg := "code_group not found"
				invalidCodeGroup := "invalid"

				var result []dtos.LookupEntry
				params := map[string]string{"codeGroup": invalidCodeGroup}
				respStatus, respErr := kMakeRequest("GET", LOOKUP_ENDPOINT, nil, &result, params, nil)

				Expect(respStatus).To(Equal(expectedCode))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})
})
