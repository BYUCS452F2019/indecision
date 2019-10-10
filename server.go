package main

import (
	"github.com/jpw547/indecision/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := ":5500"

	server := echo.New()
	server.Pre(middleware.RemoveTrailingSlash())
	server.Use(middleware.CORS())

	// define endpoints
	server.GET("/choices/:type", handlers.GetChoicesByType)

	server.Group("", middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "web-dist",
		Index:  "index.html",
		HTML5:  true,
		Browse: true,
	}))

	server.Start(port)
}
