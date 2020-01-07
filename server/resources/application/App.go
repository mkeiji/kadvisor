package application

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"kadvisor/server/repository/interfaces"
	"os"
	"sync"
)

var Db 			*gorm.DB
var Router 		*gin.Engine
var once sync.Once

type App struct {
	EntityList  []interfaces.Entity
	Controllers []interfaces.Controller
}

func init() {
	once.Do(func() {
		Db , _ = gorm.Open(os.Getenv("DB_TYPE"), getDbConnection())

		if os.Getenv("APP_ENV") == os.Getenv("PROD_ENV") {
			gin.SetMode(gin.ReleaseMode)
		}
		Router = gin.Default()
	})
}

func (a App) SetRouter() {
	Router.Use(static.Serve("/", static.LocalFile("./client/build", true)))
	a.loadControllers()
}

func (a App) Run() {
	routerErr := Router.Run(":" + os.Getenv("PORT"))
	if routerErr != nil { fmt.Println(routerErr) }

	dbErr := Db.Close()
	if dbErr != nil { fmt.Println(dbErr) }
}

func (a App) DbMigrate() {
	for _, entity := range a.EntityList {
		if os.Getenv("APP_ENV") == os.Getenv("DEV_ENV") {
			Db.DropTableIfExists(entity)
		}
		Db.AutoMigrate(entity)
	}
	a.loadConstraints()
}

func (a App) loadConstraints() {
	for _, entity := range a.EntityList {
		entity.InitializeTable(Db)
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
