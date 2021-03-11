package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/Monrevil/fonder/model"
)

// ListComments godoc
// @Summary List comments
// @Description get list of comments comments
// @Tags comments
// @Accept  json
// @Produce  json
// @Produce xml
// @Success 200 {array} model.Comment
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /comment/ [get]
func ListComments(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var comments []model.Comment
		if err := listAll(db, &comments); err != nil {
			respond(c, http.StatusInternalServerError, err.Error())
		}
		return respond(c, http.StatusOK, comments)
	}

}

// SaveComment godoc
// @Summary Add an comment
// @Description add by json comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Produce xml
// @Param comment body model.NewComment true "Add comment"
// @Security ApiKeyAuth
// @Success 201 {object} model.Comment
// @Failure 400 {object} HTTPError
// @Failure 401 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /comment/ [post]
func SaveComment(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		comment := new(model.Comment)
		if code, err := saveBody(c, db, comment); err != nil {
			return respond(c, code, err.Error())
		}
		return respond(c, http.StatusCreated, comment)
	}
}

// GetComment godoc
// @Summary Show an comment
// @Description get string by ID
// @Tags comments
// @Accept  json
// @Produce  json
// @Produce xml
// @Param id path int true "Comment ID"
// @Success 200 {object} model.Comment
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /comment/{id} [get]
func GetComment(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		comment := &model.Comment{}
		if code, err := getByID(c, db, comment); err != nil {
			return respond(c, code, err.Error())
		}
		return respond(c, http.StatusOK, comment)
	}
}

// UpdateComment godoc
// @Summary Update an comment
// @Description Update by json comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Produce xml
// @Param  comment body model.Comment true "Update comment"
// @Success 200 {object} model.Comment
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /comment/ [put]
func UpdateComment(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		comment := &model.Comment{}
		if err := c.Bind(comment); err != nil {
			return c.JSON(http.StatusBadRequest, "Could not bind comment")
		}
		if err := db.Take(&model.Comment{}, comment.ID); err != nil {
			return c.JSON(http.StatusNotFound, "record not found")
		}
		db.Save(comment)

		accept := c.Request().Header.Get("accept")
		if accept == "text/xml" {
			return c.XML(http.StatusOK, comment)
		}
		return c.JSON(http.StatusOK, comment)
	}
}

// DeleteComment godoc
// @Summary Delete an comment
// @Description Delete by comment ID
// @Tags comments
// @Accept  json
// @Produce  json
// @Produce  xml
// @Param  id path int true "comment ID" Format(int64)
// @Success 204 {object} model.Comment
// @Router /comment/{id} [delete]
func DeleteComment(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		db.Delete(&model.Comment{}, id)
		return c.JSON(http.StatusNoContent, map[string]string{id: "Was deleted"})
	}
}
