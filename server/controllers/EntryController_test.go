package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/controllers"
	"kadvisor/server/controllers/ControllerTestHelper"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	svc "kadvisor/server/services"
)

var _ = Describe("EntryController", func() {
	const (
		testUserID = 1
	)

	var (
		controller        controllers.EntryController
		mockCtrl          *g.Controller
		mockEntrySvc      *mocks.MockEntryService
		mockUsrSvc        *mocks.MockUserService
		mockAuthSvc       *mocks.MockKeiAuthService
		mockValidationSvc *mocks.MockValidationService
		w                 *httptest.ResponseRecorder
		c                 *gin.Context
		r                 *gin.Engine
		endpoint          string
		route             string
	)

	BeforeEach(func() {
		endpoint = "/entryTest"
		route = ControllerTestHelper.GetRoute(endpoint)
		w, c, r = ControllerTestHelper.SetupGinTest()
		mockCtrl = g.NewController(GinkgoT())
		mockEntrySvc = mocks.NewMockEntryService(mockCtrl)
		mockAuthSvc = mocks.NewMockKeiAuthService(mockCtrl)
		mockUsrSvc = mocks.NewMockUserService(mockCtrl)
		mockValidationSvc = mocks.NewMockValidationService(mockCtrl)
		controller = controllers.EntryController{
			Service:           mockEntrySvc,
			Auth:              mockAuthSvc,
			UsrService:        mockUsrSvc,
			ValidationService: mockValidationSvc,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := controllers.EntryController{
				Service:           svc.NewEntryService(),
				UsrService:        svc.NewUserService(),
				Auth:              svc.NewKeiAuthService(),
				ValidationService: svc.NewValidationService(),
			}

			Expect(controllers.NewEntryController()).To(Equal(expected))
		})
	})

	Describe("LoadEndpoints", func() {
		It("should set routes", func() {
			testJwt, _ := jwt.New(&jwt.GinJWTMiddleware{})
			expectedRoutes := []gin.RouteInfo{
				{
					Method: "GET",
					Path:   "/api/kadvisor/:uid/entry",
				},
				{
					Method: "POST",
					Path:   "/api/kadvisor/:uid/entry",
				},
				{
					Method: "PUT",
					Path:   "/api/kadvisor/:uid/entry",
				},
				{
					Method: "DELETE",
					Path:   "/api/kadvisor/:uid/entry",
				},
			}

			mockAuthSvc.EXPECT().
				GetAuthUtil(g.Any()).
				Return(testJwt, nil).
				Times(1)

			controller.LoadEndpoints(r)

			resultRoutes := r.Routes()
			ControllerTestHelper.VerifyRoutes(resultRoutes, expectedRoutes)
		})
	})

	Describe("GetEntry", func() {
		var (
			request     *http.Request
			queryParams map[string]string
		)

		Context("Invalid User", func() {
			It("should return error if UserService.GetOne fails", func() {
				queryParams := map[string]string{"id": "1"}
				url := ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				request, _ := http.NewRequest("GET", url, nil)
				expected := gin.H{
					"error": "test user error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusNotFound, expected)).
					Times(1)
				mockEntrySvc.EXPECT().
					GetOneById(g.Any()).
					Times(0)

				r.GET(route, controller.GetEntry)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(result).To(Equal(expected))
			})
		})

		Context("GetEntry by USER id", func() {
			BeforeEach(func() {
				queryParams = map[string]string{"limit": "5"}
				url := ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				request, _ = http.NewRequest("GET", url, nil)
			})

			It("should call entryService.GetManyByUserId and return an entries without error", func() {
				today := time.Now()
				limit := 5

				testEntries := []s.Entry{
					{
						Base:            s.Base{ID: 1},
						UserID:          testUserID,
						ClassID:         1,
						EntryTypeCodeID: "INCOME_ENTRY_TYPE",
						Date:            today,
						Amount:          float64(10),
					},
				}
				expected := []gin.H{
					{
						"classID":         float64(1),
						"entryTypeCodeID": "INCOME_ENTRY_TYPE",
						"date":            today.Format(time.RFC3339Nano),
						"amount":          float64(10),
						"id":              float64(1),
						"createdAt":       "0001-01-01T00:00:00Z",
						"updatedAt":       "0001-01-01T00:00:00Z",
						"userID":          float64(1),
					},
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockEntrySvc.EXPECT().
					GetManyByUserId(testUserID, limit).
					Return(dtos.NewKresponse(http.StatusOK, testEntries)).
					Times(1)

				r.GET(route, controller.GetEntry)
				r.ServeHTTP(c.Writer, request)

				var result []gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(len(result)).To(Equal(1))
				Expect(result).To(Equal(expected))
			})

			It("should return error if entryService.GetManyByUserId fails", func() {
				expected := gin.H{
					"error": "test entry error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockEntrySvc.EXPECT().
					GetManyByUserId(g.Any(), g.Any()).
					Return(dtos.NewKresponse(http.StatusNotFound, expected)).
					Times(1)

				r.GET(route, controller.GetEntry)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(result).To(Equal(expected))
			})
		})

		Context("GetEntry by id", func() {
			BeforeEach(func() {
				queryParams = map[string]string{"id": "1"}
				url := ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				request, _ = http.NewRequest("GET", url, nil)
			})

			It("should call entryService.GetOneById and return an entry without error", func() {
				testClassID := 1
				today := time.Now()

				testEntry := s.Entry{
					Base:            s.Base{ID: 1},
					UserID:          testUserID,
					ClassID:         testClassID,
					EntryTypeCodeID: "INCOME_ENTRY_TYPE",
					Date:            today,
					Amount:          float64(10),
				}
				expected := gin.H{
					"classID":         float64(1),
					"entryTypeCodeID": "INCOME_ENTRY_TYPE",
					"date":            today.Format(time.RFC3339Nano),
					"amount":          float64(10),
					"id":              float64(1),
					"createdAt":       "0001-01-01T00:00:00Z",
					"updatedAt":       "0001-01-01T00:00:00Z",
					"userID":          float64(1),
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockEntrySvc.EXPECT().
					GetOneById(testUserID).
					Return(dtos.NewKresponse(http.StatusOK, testEntry)).
					Times(1)

				r.GET(route, controller.GetEntry)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(result).To(Equal(expected))
			})

			It("should return error if EntryService fails", func() {
				expected := gin.H{
					"error": "test entry error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockEntrySvc.EXPECT().
					GetOneById(g.Any()).
					Return(dtos.NewKresponse(http.StatusNotFound, expected)).
					Times(1)

				r.GET(route, controller.GetEntry)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(result).To(Equal(expected))
			})
		})

		Context("GetEntries by class id", func() {
			const limit = 5

			BeforeEach(func() {
				queryParams = map[string]string{"classid": "1", "limit": u.ToString(limit)}
				url := ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				request, _ = http.NewRequest("GET", url, nil)
			})

			It("should call entryService.GetManyByClassId and return an entries without error", func() {
				today := time.Now()
				testClassID := 1

				testEntries := []s.Entry{
					{
						Base:            s.Base{ID: 1},
						UserID:          testUserID,
						ClassID:         testClassID,
						EntryTypeCodeID: "INCOME_ENTRY_TYPE",
						Date:            today,
						Amount:          float64(10),
					},
				}
				expected := []gin.H{
					{
						"classID":         float64(1),
						"entryTypeCodeID": "INCOME_ENTRY_TYPE",
						"date":            today.Format(time.RFC3339Nano),
						"amount":          float64(10),
						"id":              float64(1),
						"createdAt":       "0001-01-01T00:00:00Z",
						"updatedAt":       "0001-01-01T00:00:00Z",
						"userID":          float64(1),
					},
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockEntrySvc.EXPECT().
					GetManyByClassId(testClassID, limit).
					Return(dtos.NewKresponse(http.StatusOK, testEntries)).
					Times(1)

				r.GET(route, controller.GetEntry)
				r.ServeHTTP(c.Writer, request)

				var result []gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(len(result)).To(Equal(1))
				Expect(result).To(Equal(expected))
			})

			It("should return error if entryService.GetManyByClassId fails", func() {
				expected := gin.H{
					"error": "test entry error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockEntrySvc.EXPECT().
					GetManyByClassId(g.Any(), g.Any()).
					Return(dtos.NewKresponse(http.StatusNotFound, expected)).
					Times(1)

				r.GET(route, controller.GetEntry)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(result).To(Equal(expected))

			})
		})
	})

	Describe("PostEntry / PutEntry", func() {
		var (
			request             *http.Request
			reqBody             []byte
			today               string
			testClassID         int
			testSubClassID      int
			testEntryTypeCodeID string
			testAmount          int
			testDescription     string
			testObs             string
			url                 string
		)

		BeforeEach(func() {
			testClassID = 1
			testSubClassID = 1
			testEntryTypeCodeID = "EXPENSE_ENTRY_TYPE"
			testAmount = 5
			testDescription = "test"
			testObs = "test obs"

			today = time.Now().Format(time.RFC3339)
			reqBody = []byte(
				fmt.Sprintf(
					`{
						"userID":%v,
						"classID":"%v",
						"subClassID":"%v",
						"entryTypeCodeID":"%v",
						"date":"%v",
						"amount":"%v",
						"description":"%v",
						"obs":"%v"
					}`,
					testUserID,
					testClassID,
					testSubClassID,
					testEntryTypeCodeID,
					today,
					testAmount,
					testDescription,
					testObs,
				),
			)
			url = ControllerTestHelper.GetUrl(testUserID, route)
		})

		Context("No error", func() {
			var (
				expectedCode int
				expectedBody gin.H
				testEntry    s.Entry
			)
			BeforeEach(func() {
				testID := float64(1)
				testEntry = s.Entry{}
				expectedCode = http.StatusOK
				expectedBody = gin.H{
					"id":              testID,
					"createdAt":       "0001-01-01T00:00:00Z",
					"updatedAt":       "0001-01-01T00:00:00Z",
					"userID":          u.ToString(testUserID),
					"classID":         u.ToString(testClassID),
					"subClassID":      u.ToString(testSubClassID),
					"entryTypeCodeID": testEntryTypeCodeID,
					"date":            today,
					"amount":          u.ToString(testEntry.Amount),
					"description":     testDescription,
					"obs":             testObs,
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, g.Any())).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(expectedCode, testEntry)).
					Times(1)
			})

			AfterEach(func() {
				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})

			It("PostEntry - should call EntryService.Post and set context with ok response", func() {
				request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

				mockEntrySvc.EXPECT().
					Post(g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.POST(route, controller.PostEntry)
				r.ServeHTTP(c.Writer, request)
			})

			It("PutEntry - should call EntryService.Put and set context with ok response", func() {
				request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))

				mockEntrySvc.EXPECT().
					Put(g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.PUT(route, controller.PutEntry)
				r.ServeHTTP(c.Writer, request)
			})
		})

		Context("Error", func() {
			Context("User error", func() {
				var expectedError gin.H

				BeforeEach(func() {
					expectedError = gin.H{
						"error": "test user error",
					}

					mockUsrSvc.EXPECT().
						GetOne(testUserID, false).
						Return(dtos.NewKresponse(http.StatusNotFound, expectedError)).
						Times(1)
				})

				AfterEach(func() {
					var result gin.H
					json.Unmarshal(w.Body.Bytes(), &result)

					Expect(w.Code).To(Equal(http.StatusNotFound))
					Expect(result).To(Equal(expectedError))
				})

				It("PostEntry - should return error if UserService.GetOne fails", func() {
					request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

					mockEntrySvc.EXPECT().
						Post(g.Any()).
						Times(0)

					r.POST(route, controller.PostEntry)
					r.ServeHTTP(c.Writer, request)
				})

				It("PutEntry - should return error if UserService.GetOne fails", func() {
					request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))

					mockEntrySvc.EXPECT().
						Put(g.Any()).
						Times(0)

					r.PUT(route, controller.PutEntry)
					r.ServeHTTP(c.Writer, request)
				})
			})

			Context("ValidationService error", func() {
				var (
					expectedCode int
					expectedBody gin.H
				)

				BeforeEach(func() {
					expectedCode = http.StatusBadRequest
					expectedBody = gin.H{
						"error": "test validation error",
					}

					mockUsrSvc.EXPECT().
						GetOne(testUserID, false).
						Return(dtos.NewKresponse(http.StatusOK, g.Any())).
						Times(1)
					mockValidationSvc.EXPECT().
						GetResponse(g.Any(), g.Any()).
						Return(dtos.NewKresponse(expectedCode, expectedBody)).
						Times(1)
				})

				AfterEach(func() {
					var result gin.H
					json.Unmarshal(w.Body.Bytes(), &result)

					Expect(w.Code).To(Equal(expectedCode))
					Expect(result).To(Equal(expectedBody))
				})

				It("PostEntry - should return error if ValidationService fails", func() {
					request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

					mockEntrySvc.EXPECT().
						Post(g.Any()).
						Times(0)

					r.POST(route, controller.PostEntry)
					r.ServeHTTP(c.Writer, request)
				})

				It("PutEntry - should return error if ValidationService fails", func() {
					request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))

					mockEntrySvc.EXPECT().
						Put(g.Any()).
						Times(0)

					r.PUT(route, controller.PutEntry)
					r.ServeHTTP(c.Writer, request)
				})
			})

			Context("EntryService error", func() {
				var (
					testEntry    s.Entry
					expectedCode int
					expectedBody gin.H
				)

				BeforeEach(func() {
					testEntry = s.Entry{}
					expectedCode = http.StatusBadRequest
					expectedBody = gin.H{
						"error": "test post error",
					}

					mockUsrSvc.EXPECT().
						GetOne(testUserID, false).
						Return(dtos.NewKresponse(http.StatusOK, g.Any())).
						Times(1)
					mockValidationSvc.EXPECT().
						GetResponse(g.Any(), g.Any()).
						Return(dtos.NewKresponse(http.StatusOK, testEntry)).
						Times(1)
				})

				AfterEach(func() {
					var result gin.H
					json.Unmarshal(w.Body.Bytes(), &result)

					Expect(w.Code).To(Equal(expectedCode))
					Expect(result).To(Equal(expectedBody))
				})

				It("PostEntry - should return error if EntryService.Post fails", func() {
					request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

					mockEntrySvc.EXPECT().
						Post(g.Any()).
						Return(dtos.NewKresponse(expectedCode, expectedBody)).
						Times(1)

					r.POST(route, controller.PostEntry)
					r.ServeHTTP(c.Writer, request)
				})

				It("PutEntry - should return error if EntryService.Post fails", func() {
					request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))

					mockEntrySvc.EXPECT().
						Put(g.Any()).
						Return(dtos.NewKresponse(expectedCode, expectedBody)).
						Times(1)

					r.PUT(route, controller.PutEntry)
					r.ServeHTTP(c.Writer, request)
				})
			})
		})
	})

	Describe("DeleteEntry", func() {
		var (
			request     *http.Request
			queryParams map[string]string
			url         string
		)

		BeforeEach(func() {
			queryParams = map[string]string{"id": "1"}
			url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("DELETE", url, nil)
		})

		Context("No Error", func() {
			var (
				expectedCode int
			)

			It("should call EntryService.Delete and return deleted id", func() {
				expectedCode = http.StatusOK
				testDeletedID := 1
				expected := gin.H{"id": float64(testDeletedID)}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, g.Any())).
					Times(1)
				mockEntrySvc.EXPECT().
					Delete(testDeletedID).
					Return(dtos.NewKresponse(expectedCode, expected)).
					Times(1)

				r.DELETE(route, controller.DeleteEntry)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expected))
			})
		})

		Context("Error", func() {
			const testDeletedID = 1

			var (
				expectedError gin.H
				expectedCode  int
			)

			AfterEach(func() {
				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedError))
			})

			It("UserService - should return error if userService.GetOne fails", func() {
				expectedCode = http.StatusNotFound
				expectedError = gin.H{
					"error": "test user error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, expectedError)).
					Times(1)
				mockEntrySvc.EXPECT().
					Delete(g.Any()).
					Times(0)

				r.DELETE(route, controller.DeleteEntry)
				r.ServeHTTP(c.Writer, request)
			})

			It("EntryService - should return error if EntryService.Delete fails", func() {
				expectedCode = http.StatusNotFound
				expectedError = gin.H{
					"error": "test service error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockEntrySvc.EXPECT().
					Delete(testDeletedID).
					Return(dtos.NewKresponse(expectedCode, expectedError)).
					Times(1)

				r.DELETE(route, controller.DeleteEntry)
				r.ServeHTTP(c.Writer, request)
			})
		})
	})
})
