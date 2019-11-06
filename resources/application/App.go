package application

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"kadvisor/repository/interfaces"
	"os"
	"sync"
)
type Entity = interfaces.Entity
type Controller = interfaces.Controller

var Db 			*gorm.DB
var Router 		*gin.Engine
var once sync.Once

type App struct {
	EntityList  []Entity
	Controllers []Controller
}

func init() {
	once.Do(func() {
		Db , _ = gorm.Open("mysql", getDbConnection())
		Router = gin.Default()
	})
}

func (a App) SetRouter() {
	Router.Use(static.Serve("/", static.LocalFile("./app/build", true)))
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
		Db.DropTableIfExists(entity)
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
	dbConnect.WriteString(os.Getenv("DB_USER") + ":")
	dbConnect.WriteString(os.Getenv("DB_PASS"))
	dbConnect.WriteString(os.Getenv("DB_ADDRESS"))
	return dbConnect.String()
}
