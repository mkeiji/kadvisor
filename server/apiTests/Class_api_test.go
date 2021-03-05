package apiTests_test

import (
	"bytes"
	"net/http"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	s "kadvisor/server/repository/structs"
)

var _ = Describe("ClassApi", func() {
	const (
		CLASS_ENDPOINT   = "/api/kadvisor/:uid/class"
		TEST_NAME        = "testClass"
		TEST_DESCRIPTION = "testDescription"
	)
	var (
		testClass s.Class
	)

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

	postTestClass := func() {
		reqBody := kReqBody(buildTestClass(true))
		req := kRequestWithUser("POST", CLASS_ENDPOINT, bytes.NewBuffer(reqBody), testUserAdmin)
		resp := kSendAndAssert(req, http.StatusOK)
		kReadBody(resp, &testClass)
		defer resp.Body.Close()
	}

	/* TESTS */
	Describe("GetClass", func() {
		BeforeEach(func() {
			postTestClass()
		})

		Context("Valid id", func() {
			It("should return class with response 200", func() {
				expectedCode := http.StatusOK

				params := map[string]string{"id": strconv.Itoa(testClass.Base.ID)}
				req := kRequestWithParamAndUser("GET", CLASS_ENDPOINT, nil, testUserAdmin, params)
				resp := kSendAndAssert(req, expectedCode)

				var result s.Class
				kReadBody(resp, &result)
				Expect(result).To(Equal(testClass))
			})
		})

		Context("Invalid id", func() {
			It("should return error with response 404", func() {
				invalidID := "9999"
				expectedCode := http.StatusNotFound

				params := map[string]string{"id": invalidID}
				req := kRequestWithParamAndUser("GET", CLASS_ENDPOINT, nil, testUserAdmin, params)
				resp := kSendAndAssert(req, expectedCode)

				var result s.Class
				kReadBody(resp, &result)
				Expect(result).To(Equal(s.Class{}))
			})
		})
	})

	Describe("PostClass", func() {
		Context("No error", func() {
			It("should post a class with response ok", func() {
				reqBody := kReqBody(buildTestClass(true))
				req := kRequestWithUser("POST", CLASS_ENDPOINT, bytes.NewBuffer(reqBody), testUserAdmin)
				resp := kSendAndAssert(req, http.StatusOK)
				defer resp.Body.Close()

				var savedClass s.Class
				kReadBody(resp, &savedClass)
				Expect(savedClass.Base.ID).NotTo(BeNil())
				Expect(savedClass.Name).To(Equal(TEST_NAME))
				Expect(savedClass.Description).To(Equal(TEST_DESCRIPTION))
			})
		})

		Context("Error", func() {
			It("should return error if required fields are missing", func() {
				reqBody := kReqBody(buildTestClass(false))
				req := kRequestWithUser("POST", CLASS_ENDPOINT, bytes.NewBuffer(reqBody), testUserAdmin)
				resp := kSendAndAssert(req, http.StatusBadRequest)
				defer resp.Body.Close()

				errs := kGetResponseErrors(resp)
				Expect(len(errs)).To(Equal(1))
				Expect(errs[0].Error()).
					To(Equal("Key: 'Class.Name' Error:Field validation for 'Name' failed on the 'required' tag"))
			})
		})
	})

	Describe("PutClass", func() {
		BeforeEach(func() {
			postTestClass()
		})

		Context("No error", func() {
			It("should return updated class with ok response", func() {
				newName := "newName"
				update := buildTestClass(true)
				update.Base.ID = testClass.Base.ID
				update.Name = newName
				expectedCode := http.StatusOK

				reqBody := kReqBody(update)
				req := kRequestWithUser("PUT", CLASS_ENDPOINT, bytes.NewBuffer(reqBody), testUserAdmin)
				resp := kSendAndAssert(req, expectedCode)
				defer resp.Body.Close()

				var result s.Class
				kReadBody(resp, &result)
				Expect(result.Base.ID).To(Equal(testClass.Base.ID))
				Expect(result.Name).To(Equal(newName))
			})
		})

		Context("Errors", func() {
			It("invalid payload: should return empty object with bad request response", func() {
				expectedCode := http.StatusBadRequest

				reqBody := kReqBody(s.User{})
				req := kRequestWithUser("PUT", CLASS_ENDPOINT, bytes.NewBuffer(reqBody), testUserAdmin)
				resp := kSendAndAssert(req, expectedCode)
				defer resp.Body.Close()

				var result s.Class
				kReadBody(resp, &result)
				Expect(result).To(Equal(s.Class{}))
			})

			It("invalid id: should return empty object with not found response", func() {
				expectedCode := http.StatusNotFound
				invalidID := 222
				update := buildTestClass(true)
				update.Base.ID = invalidID

				reqBody := kReqBody(update)
				req := kRequestWithUser("PUT", CLASS_ENDPOINT, bytes.NewBuffer(reqBody), testUserAdmin)
				resp := kSendAndAssert(req, expectedCode)
				defer resp.Body.Close()

				var result s.Class
				kReadBody(resp, &result)
				Expect(result).To(Equal(s.Class{}))
			})
		})
	})

	Describe("Delete class", func() {
		Context("No error", func() {
			It("should return the deleted class with response ok", func() {
				postTestClass()
				params := map[string]string{"id": strconv.Itoa(testClass.Base.ID)}
				expectedCode := http.StatusOK

				req := kRequestWithParamAndUser("DELETE", CLASS_ENDPOINT, nil, testUserAdmin, params)
				resp := kSendAndAssert(req, expectedCode)
				defer resp.Body.Close()

				var result s.Class
				kReadBody(resp, &result)
				Expect(result).To(Equal(testClass))
			})
		})

		Context("Error", func() {
			It("should return empty object with response not found", func() {
				invalidID := "111"
				params := map[string]string{"id": invalidID}
				expectedCode := http.StatusNotFound

				req := kRequestWithParamAndUser("DELETE", CLASS_ENDPOINT, nil, testUserAdmin, params)
				resp := kSendAndAssert(req, expectedCode)
				defer resp.Body.Close()

				var result s.Class
				kReadBody(resp, &result)
				Expect(result).To(Equal(s.Class{}))
			})
		})
	})
})
