package application

import (
	"bytes"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"kadvisor/repository/interfaces"
	"os"
)
type Entity = interfaces.Entity
type Controller = interfaces.Controller

type App struct {
	Db 			*gorm.DB
	Router 		*gin.Engine
	EntityList 	[]Entity
	Controllers	[]Controller
}

func (a App) SetRouter() {
	a.Router.Use(static.Serve("/", static.LocalFile("./app/build", true)))
	a.loadControllers()
}

func (a App) Run() {
	routerErr := a.Router.Run(":" + os.Getenv("PORT"))
	if routerErr != nil { fmt.Println(routerErr) }

	dbErr := a.Db.Close()
	if dbErr != nil { fmt.Println(dbErr) }
}

func (a App) DbMigrate() {
	for _, entity := range a.EntityList {
		a.Db.DropTableIfExists(entity)
		a.Db.AutoMigrate(entity)
	}
	a.loadConstraints()
}

func (a App) GetDbConnection() string {
	var dbConnect bytes.Buffer
	dbConnect.WriteString(os.Getenv("DB_USER") + ":")
	dbConnect.WriteString(os.Getenv("DB_PASS"))
	dbConnect.WriteString(os.Getenv("DB_ADDRESS"))
	return dbConnect.String()
}

func (a App) loadConstraints() {
	for _, entity := range a.EntityList {
		entity.InitializeTable(a.Db)
	}
}

func (a App) loadControllers() {
	for _, controller := range a.Controllers {
		controller.LoadEndpoints(a.Router, a.Db)
	}
}
