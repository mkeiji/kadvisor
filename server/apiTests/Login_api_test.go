package apiTests_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/KeiGenUtil"
	s "kadvisor/server/repository/structs"
)

var _ = Describe("LoginApi", func() {
	const (
		USER_ENDPOINT   = "/api/user"
		LOGIN_ENDPOINT  = "/api/login"
		LOGOUT_ENDPOINT = "/api/logout"
	)
	var (
		testPassword  string
		testUserName  string
		testEmail     string
		invalidPwd    string
		invalidEmail  string
		testLoginUser s.User
	)

	BeforeEach(func() {
		testPassword = "login"
		testUserName = "testLogin"
		testEmail = KeiGenUtil.RandomString(8)
		invalidPwd = "invalid"
		invalidEmail = "invalid@email.com"
		newUserReqBody := s.User{
			FirstName: "testLoginUser",
			IsPremium: false,
			Login: s.Login{
				RoleID:   2,
				Email:    testEmail,
				UserName: testUserName,
				Password: testPassword,
			},
		}

		respStatus, respErr := kMakeRequest("POST", USER_ENDPOINT, newUserReqBody, &testLoginUser, nil, nil)
		Expect(respErr).To(BeNil())
		Expect(respStatus).To(Equal(http.StatusOK))
	})

	Describe("PostLogin", func() {
		Context("No error", func() {
			It("Should update login status to true", func() {
				loginBody := s.Login{
					Email:    testEmail,
					Password: testPassword,
				}

				var result s.Login
				respStatus, respErr := kMakeRequest("POST", LOGIN_ENDPOINT, loginBody, &result, nil, nil)
				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.Email).To(Equal(testLoginUser.Login.Email))
				Expect(result.IsLoggedIn).To(BeTrue())
			})
		})

		Context("Error", func() {
			It("Should return wrong password error", func() {
				expectedErrMsg := "Key: 'Login.Password' Error:Field validation for 'Password' wrong password"

				loginBody := s.Login{
					Email:    testEmail,
					Password: invalidPwd,
				}

				var result s.Login
				respStatus, respErr := kMakeRequest("POST", LOGIN_ENDPOINT, loginBody, &result, nil, nil)
				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})

			It("should return invalid email error if email does not exist", func() {
				expectedErrMsg := "Key: 'Login.Email' Error:Field validation for 'Email' invalid email"

				loginBody := s.Login{
					Email:    invalidEmail,
					Password: testPassword,
				}

				var result s.Login
				respStatus, respErr := kMakeRequest("POST", LOGIN_ENDPOINT, loginBody, &result, nil, nil)
				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})

	Describe("PutLogin", func() {
		const newName = "newUserName"
		Context("No error", func() {
			It("Should update login property", func() {
				loginBody := s.Login{
					Base:     s.Base{ID: testLoginUser.Login.Base.ID},
					UserName: newName,
				}

				var result s.Login
				respStatus, respErr := kMakeRequest("PUT", LOGIN_ENDPOINT, loginBody, &result, nil, nil)
				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.UserName).To(Equal(newName))
			})
		})

		Context("Error", func() {
			It("Should return not found error if email is invalid", func() {
				expectedErrMsg := "record not found"
				invalidID := 9999

				loginBody := s.Login{
					Base:     s.Base{ID: invalidID},
					UserName: newName,
				}

				var result s.Login
				respStatus, respErr := kMakeRequest("PUT", LOGIN_ENDPOINT, loginBody, &result, nil, nil)
				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})

	Describe("PostLogout", func() {
		Context("No error", func() {
			It("Should update login status to false", func() {
				loginBody := s.Login{
					Email: testEmail,
				}

				var result s.Login
				respStatus, respErr := kMakeRequest("POST", LOGOUT_ENDPOINT, loginBody, &result, nil, nil)
				Expect(respErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.Email).To(Equal(testLoginUser.Login.Email))
				Expect(result.IsLoggedIn).To(BeFalse())
			})
		})

		Context("Error", func() {
			It("Should return not found error if email is invalid", func() {
				expectedErrMsg := "record not found"

				loginBody := s.Login{
					Email: invalidEmail,
				}

				var result s.Login
				respStatus, respErr := kMakeRequest("POST", LOGOUT_ENDPOINT, loginBody, &result, nil, nil)
				Expect(respStatus).To(Equal(http.StatusNotFound))
				Expect(len(respErr)).To(Equal(1))
				Expect(respErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})
})
