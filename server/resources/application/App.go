package application

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
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
		Db, _ = gorm.Open(os.Getenv("DB_TYPE"), getDbConnection())

		if os.Getenv("APP_ENV") == os.Getenv("PROD_ENV") {
			gin.SetMode(gin.ReleaseMode)
		}
		Router = gin.Default()
	})
}

func (a App) SetRouter() {
	if os.Getenv("APP_ENV") == os.Getenv("PROD_ENV") {
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

	dbErr := Db.Close()
	if dbErr != nil {
		fmt.Println(dbErr)
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

	if os.Getenv("DB_TYPE") == os.Getenv("DB_MYSQL") {
		dbConnect.WriteString(os.Getenv("DB_USER") + ":")
		dbConnect.WriteString(os.Getenv("DB_PASS"))
		dbConnect.WriteString(os.Getenv("DB_ADDRESS"))
	} else if os.Getenv("DB_TYPE") == os.Getenv("DB_SQLITE") {
		dbConnect.WriteString(os.Getenv("DB_SQLITE_PATH"))
	}

	return dbConnect.String()
}
