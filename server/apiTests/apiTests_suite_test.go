package apiTests_test

import (
	"bytes"
	c "context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"kadvisor/server/apiTests/ApiTestUtil"
	"kadvisor/server/libs/KeiGenUtil"
	s "kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
	"kadvisor/server/resources/registration"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/docker/go-connections/nat"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	testDbName     = "testdb"
	testDbUserName = "root"
	testDbPassword = "root"
	testDbType     = "mysql"
	testAdminEmail = "admin@test.com"
	testAdminPwd   = "admin"
)

var (
	ctx             c.Context
	req             tc.ContainerRequest
	mysqlC          tc.Container
	sqlDb           *sql.DB
	gormDb          *gorm.DB
	testDbPort      nat.Port
	testDbHost      string
	testServer      *httptest.Server
	testUserAdmin   s.User
	testUserRegular s.User
	testAuth        s.Auth
	Khello          string
)

func TestApiTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApiTests Suite")
}

var _ = BeforeSuite(func() {
	createDbTestContainer()
	setupAppTestDb()

	runApp()
	setupTestUsers()
	getAdminToken()
})

var _ = AfterSuite(func() {
	testServer.Close()
	err := mysqlC.Terminate(ctx)
	if err != nil {
		panic(err)
	}
	sqlDb.Close()
})

func createDbTestContainer() {
	ctx = c.Background()
	req = tc.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": testDbPassword,
			"MYSQL_DATABASE":      testDbName,
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	mysqlC, _ = tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func setupAppTestDb() {
	host, _ := mysqlC.Host(ctx)
	p, _ := mysqlC.MappedPort(ctx, "3306/tcp")
	port := p.Int()
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		testDbUserName, testDbPassword, host, port, testDbName,
	)

	var sqlErr error
	sqlDb, sqlErr = sql.Open("mysql", connectionString)
	Expect(sqlErr).ShouldNot(HaveOccurred())
	sqlDb.SetMaxIdleConns(0)

	var gormErr error
	gormDb, gormErr = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	Expect(gormErr).ShouldNot(HaveOccurred())

	application.Db = gormDb
}

func runApp() {
	var app application.App
	app.InitializeInTestMode()
	app.EntityList = registration.EntityList
	app.Controllers = registration.ControllerList
	app.DbMigrate()
	app.SetRouter()

	testServer = httptest.NewServer(application.Router)
}

func setupTestUsers() {
	users := ApiTestUtil.CreateTestUsers()
	adminReqBody, adminJsonErr := json.Marshal(users[0])
	Expect(adminJsonErr).ShouldNot(HaveOccurred())
	regularReqBody, regularJsonErr := json.Marshal(users[1])
	Expect(regularJsonErr).ShouldNot(HaveOccurred())

	adminPostResp, adminPostErr := http.Post(
		testServer.URL+"/api/user",
		"application/json",
		bytes.NewBuffer(adminReqBody),
	)
	Expect(adminPostErr).ShouldNot(HaveOccurred())
	defer adminPostResp.Body.Close()
	adminBodyBytes, adminIoErr := ioutil.ReadAll(adminPostResp.Body)
	Expect(adminIoErr).ShouldNot(HaveOccurred())
	json.Unmarshal(adminBodyBytes, &testUserAdmin)

	regularPostResp, regularPostErr := http.Post(
		testServer.URL+"/api/user",
		"application/json",
		bytes.NewBuffer(regularReqBody),
	)
	Expect(regularPostErr).ShouldNot(HaveOccurred())
	defer regularPostResp.Body.Close()
	regularBodyBytes, regularIoErr := ioutil.ReadAll(regularPostResp.Body)
	Expect(regularIoErr).ShouldNot(HaveOccurred())
	json.Unmarshal(regularBodyBytes, &testUserRegular)
}

func getAdminToken() {
	reqBody, jsonErr := json.Marshal(map[string]string{
		"email":    testAdminEmail,
		"password": testAdminPwd,
	})
	Expect(jsonErr).ShouldNot(HaveOccurred())

	makeTokenRequest(reqBody)
}

func makeTokenRequest(reqBody []byte) {
	resp, respErr := http.Post(
		testServer.URL+"/api/auth",
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	Expect(respErr).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(Equal(http.StatusOK))
	defer resp.Body.Close()

	authBodyBytes, authIoErr := ioutil.ReadAll(resp.Body)
	Expect(authIoErr).ShouldNot(HaveOccurred())
	json.Unmarshal(authBodyBytes, &testAuth)
}

func kRequestWithParamAndUser(
	reqType string, endpoint string, body io.Reader, user s.User, params map[string]string,
) *http.Request {
	url := strings.Replace(endpoint, ":uid", strconv.Itoa(user.Login.UserID), -1)
	return kRequestWithParam(reqType, url, params, body)
}

func kRequestWithUser(reqType string, endpoint string, body io.Reader, user s.User) *http.Request {
	url := strings.Replace(endpoint, ":uid", strconv.Itoa(user.Login.UserID), -1)
	return kRequest(reqType, url, body)
}

func kRequestWithParam(
	reqType string, endpoint string, params map[string]string, body io.Reader,
) *http.Request {
	if len(params) > 0 {
		endpoint = endpoint + kRequestParams(params)
	}
	return kRequest(reqType, endpoint, body)
}

func kRequest(reqType string, endpoint string, body io.Reader) *http.Request {
	bearer := fmt.Sprintf("Bearer %v", testAuth.Token)

	req, reqErr := http.NewRequest(reqType, testServer.URL+endpoint, body)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	Expect(reqErr).ShouldNot(HaveOccurred())

	return req
}

func kRequestParams(params map[string]string) string {
	result := "?"
	i := 0
	for k, v := range params {
		var ampersend string
		if i == 0 {
			ampersend = ""
		} else {
			ampersend = "&"
		}

		result = fmt.Sprintf("%s%s%s=%s", result, ampersend, k, v)
		i++
	}
	return result
}

func kReqBody(v interface{}) []byte {
	reqBody, jsonErr := json.Marshal(v)
	Expect(jsonErr).ShouldNot(HaveOccurred())
	return reqBody
}

func kReadBody(resp *http.Response, c interface{}) {
	bodyBytes, ioErr := ioutil.ReadAll(resp.Body)
	Expect(ioErr).ShouldNot(HaveOccurred())
	json.Unmarshal(bodyBytes, c)
}

func kGetResponseErrors(resp *http.Response) []error {
	var errMap []map[string]interface{}
	kReadBody(resp, &errMap)

	var errs []error
	for _, mapErr := range errMap {
		errs = append(errs, kReadError(mapErr))
	}

	return errs
}

func kReadError(errMap map[string]interface{}) error {
	return errors.New(errMap["error"].(string))
}

func kSendAndAssert(req *http.Request, expectedCode int) *http.Response {
	client := &http.Client{}
	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Printf("\nError: %v", respErr)
	}
	Expect(respErr).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).Should(Equal(expectedCode))
	return resp
}

func kSend(req *http.Request) (*http.Response, []error) {
	var errList []error

	client := &http.Client{}
	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Printf("\nError: %v", respErr)
	}
	Expect(respErr).ShouldNot(HaveOccurred())

	if resp.StatusCode != http.StatusOK {
		var errL []map[string]interface{}
		kReadBody(resp, &errL)
		errList = KeiGenUtil.GetErrorList(errL)
	}

	return resp, errList
}

func kMakeRequest(
	reqType string,
	endpoint string,
	sendBody interface{},
	returnBody interface{},
	params map[string]string,
	user interface{},
) (int, []error) {
	var req *http.Request
	var reqUser s.User
	if user == nil {
		reqUser = testUserRegular
	} else {
		reqUser = user.(s.User)
	}

	reqBody := kReqBody(sendBody)

	if params == nil {
		req = kRequestWithUser(reqType, endpoint, bytes.NewBuffer(reqBody), reqUser)
	} else {
		req = kRequestWithParamAndUser(reqType, endpoint, bytes.NewBuffer(reqBody), reqUser, params)
	}

	resp, err := kSend(req)
	if err != nil {
		return resp.StatusCode, err
	}
	defer resp.Body.Close()

	kReadBody(resp, &returnBody)
	return resp.StatusCode, nil
}
