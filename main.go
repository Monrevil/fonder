package main

import (
	"mnt/c/Users/DELL/nix/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title App Example API
// @version 1.0
// @description This is a sample server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @contact.name Login with google here
// @contact.url  https://accounts.google.com/o/oauth2/auth?client_id=732894594352-3pa74unjkjdsq6ql7nbmtaor2t2735jv.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost%3A1323%2FauthGood&response_type=code&scope=email+profile&state=1234

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
