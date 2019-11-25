package main

import (
	"flag"
	"fmt"
	"github.com/jpw547/indecision/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	var migrationDir = flag.String("migration.files", "./migrations", "Directory where the migration files are located ?")
	var username = flag.String("username", "user", "Username for mysql database")
	var password = flag.String("password", "user", "Password for mysql database")
	var dbType = flag.String("database", "mysql", "Database type: mongo or mysql")
	var mockData = flag.Bool("mockData",false, "Insert Mock Data")

	port := ":5500"
	flag.Parse()

	var err error
	if *dbType == "mongo" {
		err = handlers.InitNoSQLStore(*mockData)
	}else{
		err = handlers.InitStore(*username,*password,*migrationDir)
	}
	if err!=nil {
		fmt.Println("Unable to connect ",err)
	}

	server := echo.New()
	server.Pre(middleware.RemoveTrailingSlash())
	server.Use(middleware.CORS())

	// define endpoints
	server.GET("/choices/:type", handlers.GetChoicesByType)
	server.POST("/choices/:type", handlers.CreateChoice)
	server.DELETE("/choices/:type", handlers.DeleteChoiceByID)

	server.POST("/users", handlers.CreateUser)
	server.GET("/users/:username", handlers.GetUser)
	server.DELETE("/users/:username", handlers.DeleteUserByUsername)
	server.Group("", middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "web-dist",
		Index:  "index.html",
		HTML5:  true,
		Browse: true,
	}))

	server.Start(port)
}
