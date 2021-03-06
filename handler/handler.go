package handler

import (
	"net/http"

	"github.com/Monrevil/fonder/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func listAll(db *gorm.DB, i interface{}) error {
	err := model.PingDB(db)
	if err != nil {
		return err
	}
	result := db.Find(i)
	if result.Error != nil {
		return err
	}
	return nil
}

func getByID(c echo.Context, db *gorm.DB, i interface{}) (int, error) {
	id := c.Param("id")
	err := model.PingDB(db)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	result := db.Take(i, id)
	if result.Error != nil {
		return http.StatusNotFound, result.Error
	}
	return 200, nil
}

type body interface {
	Sanitize(c echo.Context)
	Validate() error
}

func saveBody(c echo.Context, db *gorm.DB, b body) (int, error) {
	if err := c.Bind(b); err != nil {
		return http.StatusBadRequest, err
	}
	b.Sanitize(c)

	if err := b.Validate(); err != nil {
		return http.StatusBadRequest, err
	}
	model.PingDB(db)
	db.Create(b)
	return 200, nil
}

//Response format handling
func respond(c echo.Context, code int, i interface{}) error {
	accept := c.Request().Header.Get("accept")
	if accept == "text/xml" {
		return c.XML(code, i)
	}
	return c.JSON(code, i)
}
