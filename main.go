package main

import (
	_ "database/sql"
	_ "github.com/gin-contrib/static"  //for serving static files
	_ "github.com/go-sql-driver/mysql" //for mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/joho/godotenv/autoload"
	"kadvisor/controllers"
	"kadvisor/repository/interfaces"
	"kadvisor/repository/structs"
	"kadvisor/resources/application"
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

	//TODO:
	// docker run --rm -d --name test -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:5.7

	app.DbMigrate()
	app.SetRouter()
	app.Run()
}

