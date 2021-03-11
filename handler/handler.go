package handler

import (
	"github.com/labstack/echo/v4"
)

//Response format handling
func respond(c echo.Context, code int, i interface{}) error {
	accept := c.Request().Header.Get("accept")
	if accept == "text/xml" {
		return c.XML(code, i)
	}
	return c.JSON(code, i)
}
