package apiTests_test

import (
	"net/http"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/KeiGenUtil"
	s "kadvisor/server/repository/structs"
)

var _ = Describe("UserApi", func() {
	const (
		USER_ENDPOINT                = "/api/user"
		GET_AND_DELETE_USER_ENDPOINT = "/api/user/:id"
		GET_MANY_USER_ENDPOINT       = "/api/users"
	)
	var (
		testPassword        string
		testUserName        string
		testEmail           string
		testUser            s.User
		testClassName       string
		testClassDescrption string
	)

	buildTestUser := func(class ...s.Class) s.User {
		return s.User{
			FirstName: "testLoginUser",
			IsPremium: false,
			Login: s.Login{
				RoleID:   2,
				Email:    testEmail,
				UserName: testUserName,
				Password: testPassword,
			},
			Classes: class,
		}
	}

	postUser := func(user ...s.User) s.User {
		var respErr []error
		var result s.User

		if len(user) == 0 {
			_, respErr = kMakeRequest("POST", USER_ENDPOINT, buildTestUser(), &result, nil, nil)
		} else {
			_, respErr = kMakeRequest("POST", USER_ENDPOINT, user[0], &result, nil, nil)
		}
		Expect(respErr).To(BeNil())

		return result
	}

	BeforeEach(func() {
		testPassword = "login"
		testUserName = "testLogin"
		testEmail = KeiGenUtil.RandomString(8)
		testClassName = "testClass"
		testClassDescrption = "testDescription"
	})

	Describe("PostUser", func() {
		Context("No error", func() {
			It("should post a user and return ok response", func() {
				userBody := buildTestUser()

				var result s.User
				respStatus, respErr := kMakeRequest("POST", USER_ENDPOINT, userBody, &result, nil, nil)

				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.Base.ID).NotTo(BeNil())
				Expect(result.FirstName).To(Equal(userBody.FirstName))
				Expect(result.Login.Email).To(Equal(userBody.Login.Email))
			})
		})

		Context("Error", func() {
			It("should return error with bad request response if login email already exists", func() {
				testUser = postUser()
				expectedErrMsg :=
					"Key: 'User.Login.Email' Error:Field validation for 'Email' email already exists"

				var result s.User
				respStatus, respErr := kMakeRequest("POST", USER_ENDPOINT, buildTestUser(), &result, nil, nil)
				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})

	Describe("GetOneUser", func() {
		var (
			testClass     s.Class
			userWithClass s.User
			getEndpoint   string
		)

		BeforeEach(func() {
			testClass = s.Class{
				Name:        testClassName,
				Description: testClassDescrption,
			}
			userWithClass = buildTestUser(testClass)
			testUser = postUser(userWithClass)

			getEndpoint = updateGetAndDeleteEndpoint(GET_AND_DELETE_USER_ENDPOINT, testUser)
		})

		Context("No error", func() {
			It("No preload - should return a user with response ok", func() {
				var result s.User
				params := map[string]string{"preloaded": "false"}
				respStatus, respErr := kMakeRequest("GET", getEndpoint, nil, &result, params, nil)
				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.Classes).To(BeNil())
				Expect(result.Base.ID).To(Equal(testUser.Base.ID))
			})

			It("Preloaded - should return a user with children and response ok", func() {
				var result s.User
				params := map[string]string{"preloaded": "true"}
				respStatus, respErr := kMakeRequest("GET", getEndpoint, nil, &result, params, nil)
				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result).To(Equal(testUser))
			})
		})

		Context("Error", func() {
			It("should return error with not found response if user does not exist", func() {
				expectedErrMsg := "record not found"
				invalidID := "99"
				getEndpoint = strings.Replace(
					GET_AND_DELETE_USER_ENDPOINT, ":id", invalidID, -1,
				)

				var result s.User
				respStatus, respErr := kMakeRequest("GET", getEndpoint, nil, &result, nil, nil)

				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})

	Describe("GetAllUsers", func() {
		Context("No error", func() {
			var (
				result     []s.User
				respStatus int
				respErr    []error
				testUser1  s.User
				testUser2  s.User
			)

			BeforeEach(func() {
				testClass := s.Class{
					Name:        testClassName,
					Description: testClassDescrption,
				}

				bodyUser1 := buildTestUser(testClass)
				bodyUser1.Login.Email = KeiGenUtil.RandomString(8)
				testUser1 = postUser(bodyUser1)

				bodyUser2 := buildTestUser(testClass)
				bodyUser2.Login.Email = KeiGenUtil.RandomString(8)
				testUser2 = postUser(bodyUser2)

			})

			AfterEach(func() {
				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(len(result) > 0).To(BeTrue())
			})

			It("No preload - should return multiple users with ok response", func() {
				params := map[string]string{"preloaded": "false"}
				respStatus, respErr = kMakeRequest("GET", GET_MANY_USER_ENDPOINT, nil, &result, params, nil)

				verifyUserIsInList(testUser1, result, false)
				verifyUserIsInList(testUser2, result, false)
			})

			It("Preloaded - should return multiple users with children and ok response", func() {
				params := map[string]string{"preloaded": "true"}
				respStatus, respErr = kMakeRequest("GET", GET_MANY_USER_ENDPOINT, nil, &result, params, nil)

				verifyUserIsInList(testUser1, result, true)
				verifyUserIsInList(testUser2, result, true)
			})
		})
	})

	Describe("UpdateUser", func() {
		const newName = "newName"

		BeforeEach(func() {
			testUser = postUser()
		})

		Context("No error", func() {
			It("should return updated user with ok response", func() {
				updateBody := s.User{
					Base:      s.Base{ID: testUser.Base.ID},
					FirstName: newName,
				}

				var result s.User
				respStatus, respErr := kMakeRequest("PUT", USER_ENDPOINT, updateBody, &result, nil, nil)
				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.FirstName).To(Equal(newName))
			})
		})

		Context("Error", func() {
			It("should return error with bad request response", func() {
				expectedErrMsg := "record not found"
				invalidID := 9999
				updateBody := s.User{
					Base:      s.Base{ID: invalidID},
					FirstName: newName,
				}

				var result s.User
				respStatus, respErr := kMakeRequest("PUT", USER_ENDPOINT, updateBody, &result, nil, nil)
				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})

	Describe("", func() {
		var (
			deleteEndpoint string
		)

		BeforeEach(func() {
			testUser = postUser()
		})

		Context("No error", func() {
			It("should return deleted user with ok response", func() {
				deleteEndpoint = updateGetAndDeleteEndpoint(GET_AND_DELETE_USER_ENDPOINT, testUser)

				var result s.User
				respStatus, respErr := kMakeRequest(
					"DELETE", deleteEndpoint, nil, &result, nil, nil,
				)

				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.Base.ID).To(Equal(testUser.Base.ID))
			})
		})

		Context("Error", func() {
			It("should return error with not found response if id is invalid", func() {
				expectedErrMsg := "record not found"
				invalidUser := s.User{Base: s.Base{ID: 9999}}
				deleteEndpoint = updateGetAndDeleteEndpoint(GET_AND_DELETE_USER_ENDPOINT, invalidUser)

				var result s.User
				respStatus, respErr := kMakeRequest(
					"DELETE", deleteEndpoint, nil, &result, nil, nil,
				)

				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})
})

func verifyUserIsInList(user s.User, uList []s.User, checkForClasses bool) {
	isInList := false
	for _, u := range uList {
		if u.Base.ID == user.Base.ID {
			isInList = true

			if checkForClasses {
				Expect(len(u.Classes) > 0).To(BeTrue())
			} else {
				Expect(len(u.Classes)).To(Equal(0))
			}
		}
	}
	Expect(isInList).To(BeTrue())
}

func updateGetAndDeleteEndpoint(endpoint string, user s.User) string {
	return strings.Replace(
		endpoint, ":id", strconv.Itoa(user.Base.ID), -1,
	)
}
