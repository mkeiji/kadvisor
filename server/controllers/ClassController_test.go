package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/controllers"
	"kadvisor/server/controllers/ControllerTestHelper"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	svc "kadvisor/server/services"
)

var _ = Describe("ClassController", func() {
	const (
		testUserID = 1
	)

	var (
		controller        controllers.ClassController
		mockCtrl          *g.Controller
		mockClassSvc      *mocks.MockClassService
		mockAuthSvc       *mocks.MockKeiAuthService
		mockUsrSvc        *mocks.MockUserService
		mockValidationSvc *mocks.MockValidationService
		w                 *httptest.ResponseRecorder
		c                 *gin.Context
		r                 *gin.Engine
		endpoint          string
		route             string
	)

	BeforeEach(func() {
		endpoint = "/classTest"
		route = ControllerTestHelper.GetRoute(endpoint)
		w, c, r = ControllerTestHelper.SetupGinTest()
		mockCtrl = g.NewController(GinkgoT())
		mockClassSvc = mocks.NewMockClassService(mockCtrl)
		mockAuthSvc = mocks.NewMockKeiAuthService(mockCtrl)
		mockUsrSvc = mocks.NewMockUserService(mockCtrl)
		mockValidationSvc = mocks.NewMockValidationService(mockCtrl)
		controller = controllers.ClassController{
			Service:           mockClassSvc,
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
			expected := controllers.ClassController{
				Service:           svc.NewClassService(),
				Auth:              svc.NewKeiAuthService(),
				UsrService:        svc.NewUserService(),
				ValidationService: svc.NewValidationService(),
			}

			Expect(controllers.NewClassController()).To(Equal(expected))
		})
	})

	Describe("LoadEndpoints", func() {
		It("should set routes", func() {
			testJwt, _ := jwt.New(&jwt.GinJWTMiddleware{})
			expectedRoutes := []gin.RouteInfo{
				{
					Method: "GET",
					Path:   "/api/kadvisor/:uid/class",
				},
				{
					Method: "POST",
					Path:   "/api/kadvisor/:uid/class",
				},
				{
					Method: "PUT",
					Path:   "/api/kadvisor/:uid/class",
				},
				{
					Method: "DELETE",
					Path:   "/api/kadvisor/:uid/class",
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

	Describe("GetOneById", func() {
		var request *http.Request

		BeforeEach(func() {
			queryParams := map[string]string{"id": "1"}
			url := ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("GET", url, nil)
		})

		Context("No Error", func() {
			It("should call ClassService and set context with ok response", func() {
				testID := s.Base{ID: 1}

				testClass := s.Class{
					Base:        testID,
					UserID:      testUserID,
					Name:        "test",
					Description: "test",
				}
				expected := gin.H{
					"description": "test",
					"id":          float64(testID.ID),
					"createdAt":   "0001-01-01T00:00:00Z",
					"updatedAt":   "0001-01-01T00:00:00Z",
					"userID":      float64(testUserID),
					"name":        "test",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockClassSvc.EXPECT().
					GetClass(testUserID, testID.ID).
					Return(dtos.NewKresponse(http.StatusOK, testClass)).
					Times(1)

				r.GET(route, controller.GetOneById)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(result).To(Equal(expected))
			})
		})

		Context("Error", func() {
			It("should return error if UserService.GetOne fails", func() {
				expected := gin.H{
					"error": "test user error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusNotFound, expected)).
					Times(1)
				mockClassSvc.EXPECT().
					GetClass(g.Any(), g.Any()).
					Times(0)

				r.GET(route, controller.GetOneById)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(result).To(Equal(expected))
			})

			It("should return error if Service.GetClass fails", func() {
				testID := 1
				expected := gin.H{
					"error": "test class error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockClassSvc.EXPECT().
					GetClass(testUserID, testID).
					Return(dtos.NewKresponse(http.StatusBadRequest, expected)).
					Times(1)

				r.GET(route, controller.GetOneById)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
				Expect(result).To(Equal(expected))
			})
		})
	})

	Describe("PostClass", func() {
		var (
			request         *http.Request
			testID          int
			testName        string
			testDescription string
		)

		BeforeEach(func() {
			testID = 1
			testName = "test name"
			testDescription = "test description"
			reqBody := []byte(
				fmt.Sprintf(
					`{
						"userID":%v,
						"name":"%v",
						"description":"%v"
					}`,
					testUserID,
					testName,
					testDescription,
				),
			)

			url := ControllerTestHelper.GetUrl(testUserID, route)
			request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
		})

		Context("No Error", func() {
			It("should call ClassService.Post and set context with ok response", func() {
				testClass := s.Class{}
				expectedCode := http.StatusOK
				expectedBody := gin.H{
					"id":          float64(testID),
					"createdAt":   "0001-01-01T00:00:00Z",
					"updatedAt":   "0001-01-01T00:00:00Z",
					"userID":      float64(testUserID),
					"name":        testName,
					"description": testDescription,
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, g.Any())).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(expectedCode, testClass)).
					Times(1)
				mockClassSvc.EXPECT().
					Post(g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.POST(route, controller.PostClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})
		})

		Context("Error", func() {
			It("should return error if UserService.GetOne fails", func() {
				expectedCode := http.StatusNotFound
				expectedBody := gin.H{
					"error": "test user error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)
				mockClassSvc.EXPECT().
					Post(g.Any()).
					Times(0)

				r.POST(route, controller.PostClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})

			It("should return error if body validation fails", func() {
				expectedCode := http.StatusBadRequest
				expectedBody := gin.H{
					"error": "test validation error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(http.StatusBadRequest, expectedBody)).
					Times(1)
				mockClassSvc.EXPECT().
					Post(g.Any()).
					Times(0)

				r.POST(route, controller.PostClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})

			It("should return error if post fails", func() {
				expectedCode := http.StatusBadRequest
				expectedBody := gin.H{
					"error": "test post error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockClassSvc.EXPECT().
					Post(g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.POST(route, controller.PostClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})
		})
	})

	Describe("PutClass", func() {
		var (
			request                *http.Request
			testClassID            int
			testUpdatedName        string
			testUpdatedDescription string
		)

		BeforeEach(func() {
			testClassID = 1
			testUpdatedName = "updated name"
			testUpdatedDescription = "updated description"
			reqBody := []byte(
				fmt.Sprintf(
					`{
						"id":%v,
						"name":"%v",
						"description":"%v"
					}`,
					testClassID,
					testUpdatedName,
					testUpdatedDescription,
				),
			)

			url := ControllerTestHelper.GetUrl(testUserID, route)
			request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
		})

		Context("No Error", func() {
			It("should call ClassService.Put and set context with ok response", func() {
				testClass := s.Class{}
				expectedCode := http.StatusOK
				expectedBody := gin.H{
					"id":          float64(testClassID),
					"createdAt":   "0001-01-01T00:00:00Z",
					"updatedAt":   "0001-01-01T00:00:00Z",
					"userID":      float64(testUserID),
					"name":        testUpdatedName,
					"description": testUpdatedDescription,
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, g.Any())).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(expectedCode, testClass)).
					Times(1)
				mockClassSvc.EXPECT().
					Put(g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.PUT(route, controller.PutClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})
		})

		Context("Error", func() {
			It("should return error if UserService.GetOne fails", func() {
				expectedCode := http.StatusNotFound
				expectedBody := gin.H{
					"error": "test user error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)
				mockClassSvc.EXPECT().
					Put(g.Any()).
					Times(0)

				r.PUT(route, controller.PutClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})

			It("should return error if body validation fails", func() {
				expectedCode := http.StatusBadRequest
				expectedBody := gin.H{
					"error": "test validation error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(http.StatusBadRequest, expectedBody)).
					Times(1)
				mockClassSvc.EXPECT().
					Put(g.Any()).
					Times(0)

				r.PUT(route, controller.PutClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})

			It("should return error if put fails", func() {
				expectedCode := http.StatusBadRequest
				expectedBody := gin.H{
					"error": "test put error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockClassSvc.EXPECT().
					Put(g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.PUT(route, controller.PutClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})
		})
	})

	Describe("DeleteClass", func() {
		var (
			request     *http.Request
			testClassID int
		)

		BeforeEach(func() {
			testClassID = 1
			queryParams := map[string]string{"id": "1"}
			url := ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)

			request, _ = http.NewRequest("DELETE", url, nil)
		})

		Context("No Error", func() {
			It("should call ClassService.DeleteClass and set context with ok response", func() {
				expectedCode := http.StatusOK
				expectedBody := gin.H{
					"id":        float64(testClassID),
					"createdAt": "0001-01-01T00:00:00Z",
					"updatedAt": "0001-01-01T00:00:00Z",
					"userID":    float64(testUserID),
					"name":      "test name",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, g.Any())).
					Times(1)
				mockClassSvc.EXPECT().
					Delete(testClassID).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.DELETE(route, controller.DeleteClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})
		})

		Context("Error", func() {
			It("should return error if UserService.GetOne fails", func() {
				expectedCode := http.StatusNotFound
				expectedBody := gin.H{
					"error": "test user error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)
				mockClassSvc.EXPECT().
					Delete(g.Any()).
					Times(0)

				r.DELETE(route, controller.DeleteClass)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})
		})
	})
})
