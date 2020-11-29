package controllers_test

import (
	"bytes"
	"encoding/json"
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

var _ = Describe("ForecastEntryController", func() {
	const (
		testUserID = 1
	)

	var (
		controller     controllers.ForecastEntryController
		mockCtrl       *g.Controller
		mockFcEntrySvc *mocks.MockForecastEntryService
		mockUsrSvc     *mocks.MockUserService
		mockAuthSvc    *mocks.MockKeiAuthService
		w              *httptest.ResponseRecorder
		c              *gin.Context
		r              *gin.Engine
		endpoint       string
		route          string
	)

	BeforeEach(func() {
		endpoint = "/forecastEntryTest"
		route = ControllerTestHelper.GetRoute(endpoint)
		w, c, r = ControllerTestHelper.SetupGinTest()
		mockCtrl = g.NewController(GinkgoT())
		mockFcEntrySvc = mocks.NewMockForecastEntryService(mockCtrl)
		mockAuthSvc = mocks.NewMockKeiAuthService(mockCtrl)
		mockUsrSvc = mocks.NewMockUserService(mockCtrl)
		controller = controllers.ForecastEntryController{
			Service:    mockFcEntrySvc,
			Auth:       mockAuthSvc,
			UsrService: mockUsrSvc,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := controllers.ForecastEntryController{
				Service:    svc.NewForecastEntryService(),
				Auth:       svc.NewKeiAuthService(),
				UsrService: svc.NewUserService(),
			}

			Expect(controllers.NewForecastEntryController()).To(Equal(expected))
		})
	})

	Describe("LoadEndpoints", func() {
		It("should set routes", func() {
			testJwt, _ := jwt.New(&jwt.GinJWTMiddleware{})
			expectedRoutes := []gin.RouteInfo{
				{
					Method: "PUT",
					Path:   "/api/kadvisor/:uid/forecastentry",
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

	Describe("PutForecastEntry", func() {
		var (
			request      *http.Request
			url          string
			reqBody      []byte
			expectedCode int
			expectedBody gin.H
		)

		BeforeEach(func() {
			reqBody = []byte("{}")
			url = ControllerTestHelper.GetUrl(testUserID, route)
		})

		AfterEach(func() {
			r.PUT(route, controller.PutForecastEntry)
			r.ServeHTTP(c.Writer, request)

			var result gin.H
			json.Unmarshal(w.Body.Bytes(), &result)

			Expect(w.Code).To(Equal(expectedCode))
			Expect(result).To(Equal(expectedBody))
		})

		It("should call service.Put and return an entry", func() {
			testID := float64(1)
			testForecastEntry := s.ForecastEntry{}

			expectedCode = http.StatusOK
			request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
			expectedBody = gin.H{
				"id":        testID,
				"createdAt": "0001-01-01T00:00:00Z",
				"updatedAt": "0001-01-01T00:00:00Z",
			}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, g.Any())).
				Times(1)
			mockFcEntrySvc.EXPECT().
				Put(testForecastEntry).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
		})

		It("Invalid User - should return error if UserService.Post fails", func() {
			expectedCode = http.StatusNotFound
			request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
			expectedBody = gin.H{
				"error": "test user error",
			}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
			mockFcEntrySvc.EXPECT().
				Put(g.Any()).
				Times(0)
		})
	})
})
