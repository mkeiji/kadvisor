package main

import (
	_ "database/sql"
	"fmt"
	_ "github.com/gin-contrib/static" //for serving static files
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //for mysql
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/joho/godotenv/autoload"
	"kadvisor/repository/controllers"
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

	//TODO: for mysql
	// docker run --rm -d --name test -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 mysql:5.7
	app.Db, dbErr = gorm.Open("mysql", app.GetDbConnection())
	if dbErr != nil { fmt.Println(dbErr) }
	app.DbMigrate()

	app.Router = gin.Default()
	app.SetRouter()
	app.Run()
}
