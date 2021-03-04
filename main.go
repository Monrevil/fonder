package main

import (
	"github.com/Monrevil/fonder/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title App Example API
// @version 1.0
// @description This is a sample api server.

// @contact.name Dmitrii Kozii
// @contact.email monrevil@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
//@in header
//@name Authorization

// @host localhost:1323
// @BasePath /
func main() {

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := model.InitDB()
	setRoutes(e, db)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}
