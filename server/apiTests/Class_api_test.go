package apiTests_test

import (
	"bytes"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	s "kadvisor/server/repository/structs"
)

var _ = Describe("ClassApi", func() {
	const CLASS_ENDPOINT = "/api/kadvisor/:uid/class"
	const TEST_NAME = "testClass"
	const TEST_DESCRIPTION = "testDescription"
	buildTestClass := func(valid bool) s.Class {
		if valid {
			return s.Class{
				UserID:      testUserRegular.Login.UserID,
				Name:        TEST_NAME,
				Description: TEST_DESCRIPTION,
			}
		} else {
			return s.Class{
				UserID:      testUserRegular.Login.UserID,
				Description: TEST_DESCRIPTION,
			}
		}
	}

	Describe("PostClass", func() {
		Context("Valid payload", func() {
			It("should post a class with response ok", func() {
				reqBody := kReqBody(buildTestClass(true))
				req := kRequestWithUser("POST", CLASS_ENDPOINT, bytes.NewBuffer(reqBody), testUserAdmin)
				resp := kSendRequest(req, http.StatusOK)
				defer resp.Body.Close()

				var savedClass s.Class
				kReadBody(resp, &savedClass)
				Expect(savedClass.Base.ID).NotTo(BeNil())
				Expect(savedClass.Name).To(Equal(TEST_NAME))
				Expect(savedClass.Description).To(Equal(TEST_DESCRIPTION))
			})
		})

		Context("Invalid payload", func() {
			It("should return error if required fields are missing", func() {
				reqBody := kReqBody(buildTestClass(false))
				req := kRequestWithUser("POST", CLASS_ENDPOINT, bytes.NewBuffer(reqBody), testUserAdmin)
				resp := kSendRequest(req, http.StatusBadRequest)
				defer resp.Body.Close()

				errs := kGetResponseErrors(resp)
				Expect(len(errs)).To(Equal(1))
				Expect(errs[0].Error()).
					To(Equal("Key: 'Class.Name' Error:Field validation for 'Name' failed on the 'required' tag"))
			})
		})
	})
})
