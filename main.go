package main

import (
	_ "database/sql"
	_ "github.com/gin-contrib/static"  //for serving static files
	_ "github.com/go-sql-driver/mysql" //for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/joho/godotenv/autoload"
	"kadvisor/server/controllers"
	"kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
	_ "time"
)

var dbErr error

type Entity = interfaces.Entity
type User = structs.User
type Login = structs.Login
type Role = structs.Role
type Controller = interfaces.Controller
type UserController = controllers.UserController

var app application.App

func main() {
	app.EntityList = []Entity {
		&Login{},
		&User{},
		&Role{},
	}
	app.Controllers = []Controller {
		&UserController{},
	}

	app.DbMigrate()
	app.SetRouter()
	app.Run()
}