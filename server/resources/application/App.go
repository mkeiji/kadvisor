package application

import (
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

func (this App) InitializeInTestMode() {
	once.Do(func() {
		gin.SetMode(gin.ReleaseMode)
		Router = gin.New()
		Validator = validator.New()
	})
}

func (this App) Initialize() {
	once.Do(func() {
		if os.Getenv("APP_ENV") == "PROD" {
			gin.SetMode(gin.ReleaseMode)
		}

		if Db == nil {
			Db, _ = gorm.Open(mysql.Open(getDbConnection()), &gorm.Config{})
		}

		Router = gin.Default()
		Validator = validator.New()
	})
}

func (this App) SetRouter() {
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

	this.loadControllers()
}

func (this App) Run() {
	routerErr := Router.Run(":" + os.Getenv("PORT"))
	if routerErr != nil {
		log.Fatalln(routerErr)
	}
}

func (this App) DbMigrate() {
	// 2 separated loops needed to migrate all entities before initializing
	for _, entity := range this.EntityList {
		entity.Migrate(Db)
	}
	for _, entity := range this.EntityList {
		if entity.IsInitializable() {
			entity.Initialize(Db)
		}
	}
}

func (this App) loadControllers() {
	for _, controller := range this.Controllers {
		controller.LoadEndpoints(Router)
	}
}

func getDbConnection() string {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	return BuildMysqlDbConnection(dbHost, dbName, dbUser, dbPass)
}
