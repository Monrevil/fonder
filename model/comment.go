package model

import (
	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/labstack/echo/v4"
)

// Comment struct
type Comment struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	PostID int    `json:"postID"`
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	Body   string `json:"body"`
}

//NewComment email is provided with jwt, id provided by db
type NewComment struct {
	PostID int    `json:"postID" example:"1"`
	Name   string `json:"name" validate:"required" example:"jon"`
	Body   string `json:"body" example:"I like tea"`
}

//Validate comment
func (c *Comment) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.Email, validation.Required, is.Email),
		validation.Field(&c.Body, validation.Required, validation.Length(1, 500)),
		validation.Field(&c.Name, validation.Required, validation.Length(1, 25)),
	)
}

//Sanitize sets id from JWT token and post_id to 0
func (c *Comment) Sanitize(ctx echo.Context) {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	JWTemail := claims["email"].(string)
	c.ID = 0
	c.Email = JWTemail

}
