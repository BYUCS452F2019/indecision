package main

import (
	"github.com/jpw547/indecision/handlers"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// "os"
)

func main() {
	// var migrationDir = flag.String("migration.files", "./migrations", "Directory where the migration files are located ?")
	// var username = flag.String("username", "user", "Username for mysql database")
	// var password = flag.String("password", "user", "Password for mysql database")
	port := ":5500"
	// flag.Parse()

	// //Connect to database
	// dbURL := fmt.Sprint(*username, ":", *password, "@tcp(localhost:3306)/indecision")
	// db, err := sql.Open("mysql", dbURL)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Successfully connected")
	// }
	// handlers.InitStore(db)

	// //database migrations
	// driver, err := mysql.WithInstance(db, &mysql.Config{})
	// if err != nil {
	// 	fmt.Printf("could not start sql migration... %v", err)
	// }

	// m, err := migrate.NewWithDatabaseInstance(
	// 	fmt.Sprintf("file://%s", *migrationDir), // file://path/to/directory
	// 	"mysql", driver)

	// if err != nil {
	// 	fmt.Printf("migration failed... %v", err)
	// }

	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	fmt.Printf("An error occurred while syncing the database.. %v", err)
	// }

	// fmt.Println("Database migrated")
	// defer db.Close()

	server := echo.New()
	server.Pre(middleware.RemoveTrailingSlash())
	server.Use(middleware.CORS())

	// define endpoints
	server.GET("/choices/:type", handlers.GetChoicesByType)
	server.POST("/choices/:type", handlers.CreateChoice)
	server.DELETE("/choices/:type", handlers.DeleteChoiceByID)

	server.GET("/list", handlers.GetNewList)

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
