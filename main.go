package main

import (
	_ "database/sql"
	_ "github.com/gin-contrib/static"  //for serving static files
	_ "github.com/go-sql-driver/mysql" //for mysql
	_ "gorm.io/driver/mysql"
	_ "gorm.io/gorm"
	_ "github.com/joho/godotenv/autoload"
	_ "time"
	"kadvisor/server/resources/application"
	"kadvisor/server/resources/registration"
)

var app application.App

func main() {
	app.EntityList = registration.EntityList
	app.Controllers = registration.ControllerList

	app.DbMigrate()
	app.SetRouter()
	app.Run()
}
