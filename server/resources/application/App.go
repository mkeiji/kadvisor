package application

import (
	"bytes"
	"fmt"
	"kadvisor/server/repository/interfaces"
	"kadvisor/server/resources/constants"
	"log"
	"os"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var Router *gin.Engine
var Validator *validator.Validate
var once sync.Once

type App struct {
	EntityList  []interfaces.Entity
	Controllers []interfaces.Controller
}

func (a App) Initialize() {
	once.Do(func() {
		if os.Getenv("APP_ENV") == "PROD" {
			gin.SetMode(gin.ReleaseMode)
		}

		Db, _ = gorm.Open(mysql.Open(getDbConnection()), &gorm.Config{})
		Router = gin.Default()
		Validator = validator.New()
	})
}

func (a App) SetRouter() {
	if os.Getenv("APP_ENV") == "PROD" {
		Router.Use(static.Serve("/", static.LocalFile("./client/dist/apps/kadvisor-app", true)))
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{constants.DEV_ORIGIN, constants.PROD_ORIGIN}
	corsConfig.AllowMethods = []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"}
	corsConfig.AllowHeaders = []string{
		"Origin",
		"X-Requested-With",
		"Content-Type",
		"Accept",
		"Authorization",
	}
	Router.Use(cors.New(corsConfig))

	a.loadControllers()
}

func (a App) Run() {
	routerErr := Router.Run(":" + os.Getenv("PORT"))
	if routerErr != nil {
		log.Fatalln(routerErr)
	}
}

func (a App) DbMigrate() {
	for _, entity := range a.EntityList {
		entity.Migrate(Db)
	}
	for _, entity := range a.EntityList {
		// TODO: move this if to the loop above, no need for 2 loops
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
