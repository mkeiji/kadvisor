package application

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kadvisor/server/repository/interfaces"
	"kadvisor/server/resources/constants"
	"os"
	"sync"
)

var Db *gorm.DB
var Router *gin.Engine
var once sync.Once

type App struct {
	EntityList  []interfaces.Entity
	Controllers []interfaces.Controller
}

func init() {
	once.Do(func() {
		Db, _ = gorm.Open(mysql.Open(getDbConnection()), &gorm.Config{})

		if os.Getenv("APP_ENV") == "PROD" {
			gin.SetMode(gin.ReleaseMode)
		}
		Router = gin.Default()
	})
}

func (a App) SetRouter() {
	if os.Getenv("APP_ENV") == "PROD" {
		Router.Use(static.Serve("/", static.LocalFile("./client/dist/apps/kadvisor-app", true)))
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{constants.DEV_ORIGIN, constants.PROD_ORIGIN}
	corsConfig.AllowMethods = []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"}
	Router.Use(cors.New(corsConfig))

	a.loadControllers()
}

func (a App) Run() {
	routerErr := Router.Run(":" + os.Getenv("PORT"))
	if routerErr != nil {
		fmt.Println(routerErr)
	}
}

func (a App) DbMigrate() {
	for _, entity := range a.EntityList {
		entity.Migrate(Db)
	}
	for _, entity := range a.EntityList {
		if entity.IsInitializable() {
			entity.Initialize(Db)
		}
	}
}

func (a App) loadControllers() {
	for _, controller := range a.Controllers {
		controller.LoadEndpoints(Router)
	}
}

func getDbConnection() string {
	var dbConnect bytes.Buffer
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbAddress := fmt.Sprintf(
		"@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbHost,
		dbName)

	if os.Getenv("DB_TYPE") == "mysql" {
		dbConnect.WriteString(dbUser + ":")
		dbConnect.WriteString(dbPass)
		dbConnect.WriteString(dbAddress)
	} else if os.Getenv("DB_TYPE") == "sqlite3" {
		dbConnect.WriteString(os.Getenv("DB_SQLITE_PATH"))
	}

	return dbConnect.String()
}
