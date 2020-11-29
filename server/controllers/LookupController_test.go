package controllers_test

import (
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

var _ = Describe("LookupController", func() {
	const (
		testUserID = 1
	)

	var (
		controller    controllers.LookupController
		mockCtrl      *g.Controller
		mockLookupSvc *mocks.MockLookupService
		mockUsrSvc    *mocks.MockUserService
		mockAuthSvc   *mocks.MockKeiAuthService
		w             *httptest.ResponseRecorder
		c             *gin.Context
		r             *gin.Engine
		endpoint      string
		route         string
	)

	BeforeEach(func() {
		endpoint = "/lookupTest"
		route = ControllerTestHelper.GetRoute(endpoint)
		w, c, r = ControllerTestHelper.SetupGinTest()
		mockCtrl = g.NewController(GinkgoT())
		mockLookupSvc = mocks.NewMockLookupService(mockCtrl)
		mockUsrSvc = mocks.NewMockUserService(mockCtrl)
		mockAuthSvc = mocks.NewMockKeiAuthService(mockCtrl)
		controller = controllers.LookupController{
			Service:    mockLookupSvc,
			UsrService: mockUsrSvc,
			Auth:       mockAuthSvc,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := controllers.LookupController{
				Service:    svc.NewLookupService(),
				UsrService: svc.NewUserService(),
				Auth:       svc.NewKeiAuthService(),
			}

			Expect(controllers.NewLookupController()).To(Equal(expected))
		})
	})

	Describe("LoadEndpoints", func() {
		It("should set routes", func() {
			testJwt, _ := jwt.New(&jwt.GinJWTMiddleware{})
			expectedRoutes := []gin.RouteInfo{
				{
					Method: "GET",
					Path:   "/api/kadvisor/:uid/lookup",
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

	Describe("GetLookup", func() {
		var (
			url           string
			queryParams   map[string]string
			request       *http.Request
			testCodeGroup string
		)

		BeforeEach(func() {
			testCodeGroup = "TestCodeGroup"
			queryParams = map[string]string{"codeGroup": testCodeGroup}
			url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("GET", url, nil)
		})

		It("should call service.GetAllByCodeGroup", func() {
			expectedCode := http.StatusOK
			expectedBody := []s.Code{}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, g.Any())).
				Times(1)
			mockLookupSvc.EXPECT().
				GetAllByCodeGroup(testCodeGroup).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)

			r.GET(route, controller.GetLookup)
			r.ServeHTTP(c.Writer, request)

			var result []s.Code
			json.Unmarshal(w.Body.Bytes(), &result)

			Expect(w.Code).To(Equal(expectedCode))
			Expect(result).To(Equal(expectedBody))
		})

		It("should return error if UserService.GetOne fails", func() {
			expectedCode := http.StatusNotFound
			expectedBody := gin.H{
				"error": "test user error",
			}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
			mockLookupSvc.EXPECT().
				GetAllByCodeGroup(g.Any()).
				Times(0)

			r.GET(route, controller.GetLookup)
			r.ServeHTTP(c.Writer, request)

			var result gin.H
			json.Unmarshal(w.Body.Bytes(), &result)

			Expect(w.Code).To(Equal(expectedCode))
			Expect(result).To(Equal(expectedBody))
		})
	})
})
