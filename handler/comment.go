package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"mnt/c/Users/DELL/nix/model"
)

// ListComments godoc
// @Summary List comments
// @Description get list of comments comments
// @Tags comments
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Comment
// @Header 200 {string} Token "qwerty"
// @Failure 400,404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /comment/ [get]
func ListComments(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := model.PingDB(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error()+" Could not establish db connection")
		}
		var comments []model.Comment
		result := db.Find(&comments)

		if result.Error != nil {
			return c.String(http.StatusInternalServerError, result.Error.Error()+" Could not get data from db")
		}
		return c.JSON(http.StatusOK, comments)
	}
}

// SaveComment godoc
// @Summary Add an comment
// @Description add by json comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param comment body model.NewComment true "Add comment"
// @Security ApiKeyAuth
// @Success 201 {object} model.Comment
// @Failure 400 {object} HTTPError
// @Failure 401 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /comment/ [post]
func SaveComment(db *gorm.DB) echo.HandlerFunc {
	//TODO:
	// email from jwt
	// validate email
	return func(c echo.Context) error {
		comment := new(model.Comment)
		if err := c.Bind(comment); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		JWTemail := claims["email"].(string)
		comment.Email = JWTemail

		if err := comment.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		//get info from jwt

		//sanitize
		sanitized := &model.Comment{
			PostID: comment.PostID,
			Name:   comment.Name,
			Email:  JWTemail,
			Body:   comment.Body,
		}
		model.PingDB(db)
		db.Create(sanitized)

		return c.JSON(http.StatusCreated, sanitized)
	}
}

// GetComment godoc
// @Summary Show an comment
// @Description get string by ID
// @Tags comments
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Success 200 {object} model.Comment
// @Failure 400 {object} HTTPError
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /comment/{id} [get]
func GetComment(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		err := model.PingDB(db)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		comment := &model.Comment{}
		result := db.Take(comment, id)

		if result.Error != nil {
			return c.String(http.StatusNotFound, result.Error.Error())
		}
		return c.JSON(http.StatusOK, comment)
	}
}

// UpdateComment godoc
// @Summary Update an comment
// @Description Update by json comment
// @Tags comments
// @Accept  json
// @Produce  json
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
		return c.JSON(http.StatusOK, comment)
	}
}

// DeleteComment godoc
// @Summary Delete an comment
// @Description Delete by comment ID
// @Tags comments
// @Accept  json
// @Produce  json
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
